package gosimple

import (
	"fmt"
	"testing"
)

// type assert switch
func TestTypeAssertSwitch(t *testing.T) {
	typeSwitch1(123)
	typeSwitch1("tom")
	typeSwitch2(321)
	typeSwitch2("jerry")
	// Output:
	// >> int: 123
	// >> string: tom
	// >> int: 321
	// >> string: jerry
}

// nolint
func typeSwitch1(input interface{}) {
	switch input.(type) {
	case int:
		v := input.(int)
		printInt(v)
	case string:
		v := input.(string)
		printString(v)
	}
}

func typeSwitch2(input interface{}) {
	switch v := input.(type) {
	case int:
		printInt(v)
	case string:
		printString(v)
	}
}

func printInt(v int) {
	fmt.Printf(">> int: %d\n", v)
}

func printString(v string) {
	fmt.Printf(">> string: %s\n", v)
}

// ➜ golangci-lint run
// type_assert_switch_test.go:23:9: S1034: assigning the result of this type assertion to a variable (switch input := input.(type)) could eliminate type assertions in switch cases (gosimple)
//        switch input.(type) {
//               ^
// type_assert_switch_test.go:25:8: S1034(related information): could eliminate this type assertion (gosimple)
//                v := input.(int)
//                     ^
// type_assert_switch_test.go:28:8: S1034(related information): could eliminate this type assertion (gosimple)
//                v := input.(string)
//                     ^                  ^
