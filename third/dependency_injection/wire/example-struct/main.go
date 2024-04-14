package main

import "sync"

// Struct Providers

type Foo int

func ProviderFoo() Foo { return 1 }

type Bar int

func ProviderBar() Bar { return 2 }

type FooBar struct {
	mu    sync.Mutex `wire:"-"` // wire忽略此字段
	MyFoo Foo
	MyBar Bar
}
