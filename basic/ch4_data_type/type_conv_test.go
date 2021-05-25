package ch4_data_type

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// int -> float
func TestIntToFloat(t *testing.T) {
	i := 10
	f1 := float32(i)
	f2 := float64(i)
	fmt.Println(f1, f2)
}

// float -> int
func TestFloatToInt1(t *testing.T) {
	// 忽略小数点后数字
	var f1 float32 = 3.14
	var f2 float64 = 3.14

	i1 := int(f1)
	i2 := int64(f2)
	fmt.Println(i1, i2) // 3 3

	// 向上取整
	i3 := math.Ceil(f2)
	fmt.Println(i3) // 4

	// 向下取整
	i4 := math.Floor(f2)
	fmt.Println(i4) // 3
}

// float -> int (四舍五入)
func TestFloatToInt2(t *testing.T) {
	f1 := 3.14
	f2 := 3.54
	fmt.Println(round(f1), round(f2)) // 3 4
}

func round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// string <-> int
func TestStringInt(t *testing.T) {
	// int -> string
	var a int = 1
	s1 := strconv.Itoa(a)
	fmt.Println("s1=", s1)

	// int32 -> string
	var b int32 = 2
	s3 := strconv.FormatInt(int64(b), 10)
	fmt.Println("s3=", s3)

	// int64 -> string
	var c int64 = 3
	s2 := strconv.FormatInt(c, 10)
	fmt.Println("s2=", s2)

	// string -> int
	d := "4"
	n1, err := strconv.Atoi(d)
	if err == nil {
		fmt.Println("n1=", n1)
	}

	// string -> int32
	e := "5"
	n3, err := strconv.ParseInt(e, 10, 32)
	if err == nil {
		fmt.Println("n3=", n3)
	}

	// string -> int64
	f := "6"
	n2, err := strconv.ParseInt(f, 10, 64)
	if err == nil {
		fmt.Println("n2=", n2)
	}
}

// string <-> float
func TestStringFloat(t *testing.T) {
	// float32 -> string
	var f1 float32 = 3.14
	s1 := strconv.FormatFloat(float64(f1), 'f', 2, 32)
	fmt.Println(s1) // 3.14

	// float64 -> string
	var f2 float64 = 3.14
	s2 := strconv.FormatFloat(f2, 'f', -1, 64)
	fmt.Println(s2) // 3.14

	// string -> float32
	s3 := "3.14"
	f3, err := strconv.ParseFloat(s3, 32)
	fmt.Println(f3, err) // 3.140000104904175 <nil>

	// string -> float64
	s4 := "3.14"
	f4, err := strconv.ParseFloat(s4, 64)
	fmt.Println(f4, err) // 3.14 <nil>

	// float64 -> string
	f := 123.456

	// 科学计数法
	fmt.Println(fmt.Sprintf("%e", f)) // 1.234560e+02

	// 小数格式
	fmt.Println(fmt.Sprintf("%f", f)) // 123.456000

	// 指定小数点显示后几位 (四舍五入)
	fmt.Println(fmt.Sprintf("%.2f", f)) // 123.46

	// 显示必要小数点后几位
	fmt.Println(fmt.Sprintf("%g", f)) // 123.456
}

// string <-> byte[]
func TestStringBytes(t *testing.T) {
	str := "hello哦"

	// string -> byte[]
	bytes := []byte(str)
	fmt.Println("bytes:", bytes) // [104 101 108 108 111 229 147 166]

	// byte[] -> string
	str2 := string(bytes)
	fmt.Println("str2:", str2)
}

// string <-> rune[]
func TestStringRunes(t *testing.T) {
	str := "hello哦"

	// string -> rune[]
	runes := []rune(str)
	fmt.Println("runes:", runes) // [104 101 108 108 111 21734]

	// rune[] -> string
	str2 := string(runes)
	fmt.Println("str2:", str2)
}
