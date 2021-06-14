package suffixarray

import (
	"bytes"
	"fmt"
	"index/suffixarray"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// index/suffixarray：使用内存的后缀数组来实现对数时间的子字符串搜索

// 类型：suffixarray.Index 实现了一个用于快速子字符串搜索的后缀数组。
func TestTypeIndex(t *testing.T) {
	_ = suffixarray.Index{}
}

// func New(data []byte) *Index：创建一data的Index，创建时间为O(N)，N=len(data)
func TestNew(t *testing.T) {
	index := suffixarray.New([]byte("hello world."))
	assert.NotNil(t, index)
}

// func (x *Index) Bytes() []byte：返回创建Index的data，不可被修改
func TestIndexBytes(t *testing.T) {
	index := suffixarray.New([]byte("hello world."))
	b := index.Bytes()
	fmt.Println(string(b)) // hello world.
}

// func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)：返回正则表达式r的非重复匹配排序队列
// n<0返回所有，n>0返回指定个数
func TestFindAllIndex(t *testing.T) {
	index := suffixarray.New([]byte("hello world."))
	rex, err := regexp.Compile("lo")
	if err != nil {
		panic(err)
	}
	result := index.FindAllIndex(rex, -1)
	fmt.Println(result) // [[3 5]]
}

// func (x *Index) Lookup(s []byte, n int) (result []int)：返回最多n个为排序列表，列表值为字节字符串s出现在所有data的位置
// n<0返回所有，n>0返回指定个数，如果未找到返回nil
func TestLookUp(t *testing.T) {
	index := suffixarray.New([]byte("banana"))
	offsets := index.Lookup([]byte("ana"), -1)
	for _, v := range offsets {
		fmt.Println(v)
	}
	// output:
	// 3
	// 1
}

// func (x *Index) Read(r io.Reader) error：从r读取诶index x，r不为nil
// func (x *Index) Write(w io.Writer) error：将index x写入w
func TestReadWrite(t *testing.T) {
	index := suffixarray.New([]byte("banana"))

	var buf bytes.Buffer
	err := index.Write(&buf)
	if err != nil {
		panic(err)
	}

	var i suffixarray.Index
	err = i.Read(&buf)
	if err != nil {
		panic(err)
	}
}
