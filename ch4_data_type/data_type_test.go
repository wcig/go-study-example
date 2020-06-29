package ch4_data_type

import (
	"fmt"
	"reflect"
	"testing"
)

// 布尔类型
func TestBool(t *testing.T) {
	var f1 bool = true
	f2 := false

	fmt.Println(f1, f2)
}

// 整型
func TestInt(t *testing.T) {
	var i1 int = 0
	var i2 uint = 0

	var i3 int8 = 0
	var i4 uint8 = 0

	var i5 int16 = 0
	var i6 uint16 = 0

	var i7 int32 = 0
	var i8 uint32 = 0

	var i9 int64 = 0
	var i10 uint64 = 0

	fmt.Println(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10)

	var i11 = 0
	fmt.Println(reflect.TypeOf(i11)) // int
}
