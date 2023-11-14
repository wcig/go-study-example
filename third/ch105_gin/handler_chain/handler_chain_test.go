package handler_chain

import (
	"log"
	"testing"
)

type HandlerFunc func(*Context)

type Context struct {
	handlers []HandlerFunc
	index    int
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}

func TestHandlerChain(t *testing.T) {
	handler1 := func(c *Context) {
		log.Println(">> 1 start, index:", c.index)
		c.Next()
		log.Println(">> 1 end,   index:", c.index)
	}
	handler2 := func(c *Context) {
		log.Println(">> 2 start, index:", c.index)
		c.Next()
		log.Println(">> 2 end,   index:", c.index)
	}
	handler3 := func(c *Context) {
		log.Println(">> 3 start, index:", c.index)
		c.Next()
		log.Println(">> 3 end,   index:", c.index)
	}
	handlers := []HandlerFunc{handler1, handler2, handler3}
	c := Context{
		handlers: handlers,
		index:    -1,
	}
	c.Next()
	// Output:
	// 2023/11/14 22:19:20 >> 1 start, index: 0
	// 2023/11/14 22:19:20 >> 2 start, index: 1
	// 2023/11/14 22:19:20 >> 2 end,   index: 3
	// 2023/11/14 22:19:20 >> 1 end,   index: 4
}

func TestHandlerChainAbort(t *testing.T) {
	handler1 := func(c *Context) {
		log.Println(">> 1 start, index:", c.index)
		c.Next()
		log.Println(">> 1 end,   index:", c.index)
	}
	handler2 := func(c *Context) {
		log.Println(">> 2 start, index:", c.index)
		c.Abort()
		log.Println(">> 2 end,   index:", c.index)
	}
	handler3 := func(c *Context) {
		log.Println(">> 3 start, index:", c.index)
		c.Next()
		log.Println(">> 3 end,   index:", c.index)
	}
	handlers := []HandlerFunc{handler1, handler2, handler3}
	c := Context{
		handlers: handlers,
		index:    -1,
	}
	c.Next()
	// Output:
	// 2023/11/14 22:15:55 >> 1 start, index: 0
	// 2023/11/14 22:15:55 >> 2 start, index: 1
	// 2023/11/14 22:15:55 >> 3 start, index: 2
	// 2023/11/14 22:15:55 >> 3 end,   index: 3
	// 2023/11/14 22:15:55 >> 2 end,   index: 4
	// 2023/11/14 22:15:55 >> 1 end,   index: 5
}
