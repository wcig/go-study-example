package main

import "fmt"

func main() {
	example1()

	example2()

	// Output:
	// {1} &{12}
	// {11} &{12}
}

type user struct {
	id int
}

func example1() {
	u1 := user{id: 1}
	u2 := &user{id: 2}

	defer fmt.Println(u1, u2)

	u1.id += 10
	u2.id += 10
}

func example2() {
	u1 := user{id: 1}
	u2 := &user{id: 2}

	defer func() {
		fmt.Println(u1, u2)
	}()

	u1.id += 10
	u2.id += 10
}
