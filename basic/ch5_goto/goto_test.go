package ch5_goto

import (
	"fmt"
	"testing"
	"time"
)

// if
func TestIf(t *testing.T) {
	n := 10
	if n > 0 {
		fmt.Println("n为正数")
	}
}

// if else
func TestIfElse(t *testing.T) {
	n1 := -10
	if n1 > 0 {
		fmt.Println("n为正数")
	} else {
		fmt.Println("n为负数")
	}

	if n2 := 9; n2 < 0 {
		fmt.Println(n2, "is negative")
	} else if n2 < 10 {
		fmt.Println(n2, "has 1 digit")
	} else {
		fmt.Println(n2, "has multiple digits")
	}
}

// for
func TestFor(t *testing.T) {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
	}
	fmt.Println(a) // 4

	b := 1
	for b <= 3 {
		b++
	}
	fmt.Println(b) // 4

	c := 1
	for i := 0; i < 3; i++ {
		c++
	}
	fmt.Println(c) // 4
}

// switch
func TestSwitch(t *testing.T) {
	b := 1
	switch {
	case b >= 1:
		fmt.Println("b>=1") // b>=1
		fallthrough
	case b >= 2:
		fmt.Println("b>=2") // b>=2
	case b >= 0:
		fmt.Println("b>=0")
	}

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	tt := time.Now()
	switch {
	case tt.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func TestBreak(t *testing.T) {
LABEL:
	for {
		for i := 0; i < 5; i++ {
			if i > 3 {
				break LABEL
			} else {
				fmt.Println(i)
			}
		}
	}
}

// 0
// 1
// 2
// 3

func TestContinue(t *testing.T) {
LABEL:
	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}
}

// 0
// 1
// 2
// 3
// 4

func TestGoto(t *testing.T) {
	a := 10
LABEL:
	for a < 15 {
		if a == 12 {
			a++
			goto LABEL
		}
		fmt.Println("a=", a)
		a++
	}
}

// a= 10
// a= 11
// a= 13
// a= 14
