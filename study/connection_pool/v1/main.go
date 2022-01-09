package main

import (
	"errors"
	"io"
	"sync"
)

// v1：连接池
// 参考：https://studygolang.com/articles/12333
func main() {
	//
}

var (
	ErrPoolConfig = errors.New("invalid config")
	ErrPoolClosed = errors.New("conn pool closed")
)

type Pool interface {
	Acquire() (io.Closer, error) // 获取资源
	Release(io.Closer) error     // 释放资源
	Close(io.Closer) error       // 关闭资源
	Shutdown() error             // 关闭池
}

type factory func() (io.Closer, error)

type ConnPool struct {
	mu      sync.Mutex
	factory factory
	pool    chan io.Closer
	minOpen int
	maxOpen int
	numOpen int
	closed  bool
}

func NewConnPool(minOpen, maxOpen int, factory factory) (*ConnPool, error) {
	if minOpen <= 0 || maxOpen <= 0 || minOpen > maxOpen {
		return nil, ErrPoolConfig
	}

	cp := &ConnPool{
		mu:      sync.Mutex{},
		factory: factory,
		pool:    make(chan io.Closer, maxOpen),
		minOpen: minOpen,
		maxOpen: maxOpen,
		numOpen: 0,
		closed:  false,
	}

	for i := 0; i < minOpen; i++ {
		conn, err := factory()
		if err != nil {
			continue
		}
		cp.pool <- conn
		cp.numOpen++
	}
	return cp, nil
}

func (cp *ConnPool) Acquire() (io.Closer, error) {
	if cp.closed {
		return nil, ErrPoolClosed
	}

	cp.mu.Lock()
	defer cp.mu.Unlock()

	if cp.numOpen >= cp.maxOpen {
		for {
			select {
			case conn := <-cp.pool:
				return conn, nil
			}
		}
	} else {
		conn, err := cp.factory()
		if err != nil {
			return nil, err
		}
		cp.numOpen++
		return conn, err
	}
}

func (cp *ConnPool) Release(conn io.Closer) error {
	if cp.closed {
		return ErrPoolClosed
	}

	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.pool <- conn
	cp.numOpen--
	return nil
}

func (cp *ConnPool) Close(conn io.Closer) error {
	if cp.closed {
		return ErrPoolClosed
	}

	cp.mu.Lock()
	defer cp.mu.Unlock()

	err := conn.Close()
	cp.numOpen--
	return err
}

func (cp *ConnPool) Shutdown() error {
	if cp.closed {
		return ErrPoolClosed
	}

	cp.mu.Lock()
	defer cp.mu.Unlock()

	close(cp.pool)
	for conn := range cp.pool {
		_ = conn.Close()
		cp.numOpen--
	}
	cp.closed = true
	return nil
}
