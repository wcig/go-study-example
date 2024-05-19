package test

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

// TestBytesStringCompare bytes/string比较
func TestBytesStringCompare(t *testing.T) {

	var s1 []byte = []byte("aaa")
	var s2 []byte = []byte("aaa")
	fmt.Println(bytes.Compare(s1, s2) == 0)
	fmt.Println(bytes.Compare(s1, s2) != 0)

	var s3 []byte = []byte("aaa")
	var s4 []byte = []byte("aaa")
	fmt.Println(bytes.Equal(s3, s4))

	s5 := "aaa"
	s6 := "aaa"
	fmt.Println(s5 == s6)
}

// TestContainSubStr 字符串是否包含子串
func TestContainSubStr(t *testing.T) {
	s1 := "aabbccaa"
	s2 := "aa"
	fmt.Println(strings.Index(s1, s2))      //0
	fmt.Println(strings.IndexAny(s1, s2))   //0
	fmt.Println(strings.IndexRune(s1, 'c')) //4
	fmt.Println(strings.IndexByte(s1, 'd')) //-1

	fmt.Println(strings.Contains(s1, s2))      //true
	fmt.Println(strings.ContainsAny(s1, s2))   //true
	fmt.Println(strings.ContainsRune(s1, 'c')) //true

	s := []int{1, 2, 3}
	for i := range s {
		fmt.Println(i, ":", s[i])
	}

	for _, v := range s {
		fmt.Println(v)
	}

	for range s {
		fmt.Println(s)
	}
}

// TestRawStr 使用raw字符串避免字符串转义
func TestRawStr(t *testing.T) {
	//不推荐
	regexp.MustCompile("\\.")

	//推荐
	regexp.MustCompile(`\.`)
}
