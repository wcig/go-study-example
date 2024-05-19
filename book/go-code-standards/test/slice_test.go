package test

import (
	"fmt"
	"testing"
)

// TestSliceInit slice初始化
// 推荐使用：(1)
// 原因：(1)只是声明了变量，在栈区分配了变量名的地址，(2)/(3)声明了变量，同时在堆区开辟了一块内存地址存存放了一个长度为0的切片。
// 当你在使用这个切片的时候，必然会重新开辟新内存地址存档数据，之前的地址就是垃圾数据。所以声明的话用第一种。如果你能比较确切的
// 知道切片大小，第二种里写上你的切片大小
func TestSliceInit(t *testing.T) {
	var s1 []string         //(1)
	s2 := make([]string, 0) //(2)
	s3 := []string{}        //(3)

	fmt.Println("s1:", s1, ", s1==nil:", s1 == nil, ", len(s1)=", len(s1))
	fmt.Println("s2:", s2, ", s2==nil:", s2 == nil, ", len(s2)=", len(s2))
	fmt.Println("s3:", s3, ", s3==nil:", s3 == nil, ", len(s3)=", len(s3))

	//s1: [] , s1==nil: true , len(s1)= 0
	//s2: [] , s2==nil: false , len(s2)= 0
	//s3: [] , s3==nil: false , len(s3)= 0
}

// TestSliceCopy slice copy
func TestSliceCopy(t *testing.T) {
	s1 := []int{1, 3, 2}

	var s2 []int
	for _, val := range s1 {
		s2 = append(s2, val)
	}

	var s3 []int
	copy(s3, s1)

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
}

// TestSliceAppend slice append
func TestSliceAppend(t *testing.T) {
	s1 := []int{1, 3, 2}
	s2 := []int{5, 6}

	for _, val := range s2 {
		s1 = append(s1, val)
	}

	s1 = append(s1, s2...)

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}
