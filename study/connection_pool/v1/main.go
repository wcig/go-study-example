package main

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

// v1：连接池
// 参考：https://studygolang.com/articles/12333

func main() {
	fmt.Println("----------------------")
	pool, err := NewConnPool(3, 5, func() (io.Closer, error) {
		return &conn{}, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("new, num of open:", pool.numOpen)

	fmt.Println("----------------------")
	var connList []io.Closer
	for i := 0; i < 5; i++ {
		conn, err := pool.Acquire()
		if err != nil {
			panic(err)
		}
		connList = append(connList, conn)
		fmt.Printf("acquire %d after, num of open: %d\n", i, pool.numOpen)
	}

	fmt.Println("----------------------")
	for i := 0; i < len(connList); i++ {
		err := pool.Release(connList[i])
		if err != nil {
			panic(err)
		}
		fmt.Printf("release %d after, num of open: %d\n", i, pool.numOpen)
	}

	fmt.Println("----------------------")
	for i := 0; i < 2; i++ {
		conn, err := pool.Acquire()
		if err != nil {
			panic(err)
		}
		err = pool.Close(conn)
		if err != nil {
			panic(err)
		}
		fmt.Printf("close %d after, num of open: %d\n", i, pool.numOpen)
	}

	fmt.Println("----------------------")
	err = pool.Shutdown()
	if err != nil {
		panic(err)
	}
	fmt.Println("shutdown, num of open:", pool.numOpen)
}

var (
	ErrPoolConfig = errors.New("invalid config")
	ErrPoolClosed = errors.New("conn pool closed")
)

type conn struct{}

func (c *conn) Close() error {
	return nil
}

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

	for {
		select {
		case conn := <-cp.pool:
			return conn, nil
		default:
		}

		cp.mu.Lock()

		if cp.numOpen >= cp.maxOpen {
			conn := <-cp.pool
			cp.mu.Unlock()
			return conn, nil
		}

		conn, err := cp.factory()
		if err != nil {
			cp.mu.Unlock()
			return nil, err
		}
		cp.numOpen++
		cp.mu.Unlock()
		return conn, nil
	}
}

func (cp *ConnPool) Release(conn io.Closer) error {
	if cp.closed {
		return ErrPoolClosed
	}

	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.pool <- conn
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
