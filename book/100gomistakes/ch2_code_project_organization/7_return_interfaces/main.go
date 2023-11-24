package main

import (
	"log"
)

func main() {
	n := NewUserNotifier("tom")
	n.Notify()
	// Output:
	// 2023/11/23 17:52:06 >> user [tom] notify
}

type Notifier interface {
	Notify()
}

type User struct {
	Name string
}

func (u *User) Notify() {
	log.Printf(">> user [%s] notify", u.Name)
}

func NewUserNotifier(name string) *User {
	return &User{Name: name}
}
