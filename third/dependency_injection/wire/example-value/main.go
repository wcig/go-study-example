package main

import (
	"io"

	"github.com/google/wire"
)

// Binding values

var Provider = wire.NewSet(NewFoo)

type Foo struct {
	X int
}

func NewFoo(x int) Foo {
	return Foo{X: x}
}

type App struct {
	foo Foo
	r   io.Reader
}

func NewApp(foo Foo, r io.Reader) App {
	return App{foo: foo, r: r}
}
