package ch21_generic

import (
	"fmt"
	"testing"
)

// 泛型函数
func TestFunc(t *testing.T) {
	s1 := Add(int(1), int(2))
	s2 := Add(int32(11), int32(22))
	s3 := Add(int64(111), int64(222))
	fmt.Println(s1, s2, s3)

	// error
	// Add(int(1), int64(2))
	// Add("a", "b")
}

func Add[T int | int32 | int64](a T, b T) T {
	return a + b
}

// 定义多个类型形参的泛型函数
func TestMultiTypeParamFunc(t *testing.T) {
	v := Foo("tom", 123)
	fmt.Println(v)
}

func Foo[T string, K int | int32 | int64](a T, b K) string {
	return fmt.Sprintf("Foo: %s-%d", a, b)
}

// 匿名函数不支持泛型
func TestAnonymousFunc(t *testing.T) {
	// error
	// fn := func[T int|int32|int64](v T) string {
	// 	return strconv.FormatInt(int64(v), 10)
	// }
}
