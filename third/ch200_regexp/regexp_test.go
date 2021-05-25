package ch200_regexp

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 基本匹配 (区分大小写)
func TestMatch1(t *testing.T) {
	// pattern1 := `the`
	// pattern2 := `The`
	// str := `the fat cat sat on the mat.`
	// isMatch, err := regexp.MatchString(pattern1, str)
	// assert.True(t, err == nil && isMatch)
	//
	// isMatch, err = regexp.MatchString(pattern2, str)
	// assert.True(t, err == nil && !isMatch)

	pattern1 := `the`
	pattern2 := `The`
	str := `The fat cat sat on the mat.`
	printMatch(pattern1, str)
	printMatch(pattern2, str)
}

func TestMatch2(t *testing.T) {
	pattern := `.ar`
	str := `The car parked in the garage.`

	r := regexp.MustCompile(pattern)
	isMatch := r.MatchString(str)
	assert.True(t, isMatch)
}

// 点运算符 (.)
func TestMatch3(t *testing.T) {
	pattern := `.ar`
	str := `The car parked in the garage.`
	printMatch(pattern, str) // [car par gar]
}

// 字符集
func TestMatch4(t *testing.T) {
	// 字符集 ([])
	pattern := `[Tt]he`
	str := `The car parked in the garage.`
	printMatch(pattern, str) // [The the]

	// 否定字符集 ([^])
	pattern = `ar[.]`
	str = `A garage is a good place to park a car.`
	printMatch(pattern, str) // [ar.]
}

// 重复次数
func TestMatch5(t *testing.T) {
	// *: 匹配0个或多个
	pattern := `[a-z]*`
	str := `The car parked in the garage #21.`
	printMatch(pattern, str) // [ he car parked in the garage     ]

	// +: 匹配1个或多个
	pattern = `c.+t`
	str = `The fat cat sat on the mat.`
	printMatch(pattern, str) // [cat sat on the mat]

	// ?: 标记在符号前面的字符为可选(可出现0次或1次)
	pattern = `[T]?he`
	str = `The car is parked in the garage.`
	printMatch(pattern, str) // [The he]
}

// {}: 大括号
func TestMatch6(t *testing.T) {
	str := `The number was 9.9997 but we rounded it off to 10.0.`

	// 固定次数
	pattern := `[0-9]{3}`
	printMatch(pattern, str) // [999]

	// num >= n1
	pattern = `[0-9]{2,}`
	printMatch(pattern, str) // [9997 10]

	// n1 <= num <= n2
	pattern = `[0-9]{2,3}`
	printMatch(pattern, str) // [999 10]
}

// (...) 特征标群: 括号内的内容被看成一个整体
// (ab)*: 匹配0个或多个ab, ab*: 匹配0个或多个b
func TestMatch7(t *testing.T) {
	pattern := `(c|g|p)ar`
	str := `The car is parked in the garage.`
	printMatch(pattern, str) // [car par gar]
}

// |: 或运算符
func TestMatch8(t *testing.T) {
	pattern := `(T|t)he|car`
	str := `The car is parked in the garage.`
	printMatch(pattern, str) // [The car the]
}

// \: 转义字符
func TestMatch9(t *testing.T) {
	pattern := `(f|c|m)at\.`
	str := `The fat cat sat on the mat.`
	printMatch(pattern, str) // [mat.]
}

// 锚点
func TestMatch10(t *testing.T) {
	// ^: 匹配以对应字符串开头
	pattern := `^(T|t)he`
	str := `The car is parked in the garage.`
	printMatch(pattern, str) // [The]

	// $: 匹配以对应字符串结尾
	pattern = `(at\.)$`
	str = `The fat cat. sat. on the mat.`
	printMatch(pattern, str) // [at.]
}

// 零宽度断言 (前后预查) (golang不支持)
func TestMatch11(t *testing.T) {}

// 标志 (golang不支持)
func TestMatch12(t *testing.T) {
	// 忽略大小写
	pattern := `/The/gi`
	str := `The fat cat sat on the mat.`
	printMatch(pattern, str)

	// 全局搜索
	pattern = `/.(at)/`
	str = `The fat cat sat on the mat.`
	printMatch(pattern, str)

	pattern = `/.(at)/g`
	str = `The fat cat sat on the mat.`
	printMatch(pattern, str)

	// 多行修饰符
	pattern = `/.at(.)?$/`
	str = `The fat
           cat sat
           on the mat.`
	printMatch(pattern, str)

	pattern = `/.at(.)?$/gm`
	str = `The fat
           cat sat
           on the mat.`
	printMatch(pattern, str)

	// 贪婪匹配与惰性匹配 (Greedy vs lazy matching)
	pattern = `/(.*at)/`
	str = `The fat cat sat on the mat. `
	printMatch(pattern, str)

	pattern = `/(.*?at)/`
	str = `The fat cat sat on the mat. `
	printMatch(pattern, str)
}

// 打印字符串中匹配的所有子串
func printMatch(pattern, str string) {
	r := regexp.MustCompile(pattern)
	result := r.FindAllString(str, -1)
	fmt.Println(result)
}
