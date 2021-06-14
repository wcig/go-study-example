package maphash

import (
	"fmt"
	"hash/maphash"
	"testing"
)

// "hash/maphash"
// 提供字节序列的哈希函数。这些散列函数旨在用于实现散列表或其他需要将任意字符串或字节序列映射到无符号 64 位整数上的统一分布的数据结构。哈希表或数据结构的每个不同实例都应该使用自己的种子。
// maphash.Hash继承了hash.Hash64接口

// func MakeSeed() Seed：创建一新的随机种子
func TestMakeSeed(t *testing.T) {
	seed1 := maphash.MakeSeed()
	fmt.Println(seed1)

	seed2 := maphash.MakeSeed()
	fmt.Println(seed2)
	// output:
	// {12473823243245795109}
	// {4043434146092475436}
}

// 示例
func TestExample(t *testing.T) {
	// 创建maphash.Hash
	var h1, h2 maphash.Hash

	// 写入数据，并计算h的64位置
	n, err := h1.WriteString("hello ")
	if err != nil {
		panic(err)
	}
	fmt.Println(n, h1.Sum64())

	// 附加新的数据，并计算h的64位置
	n, err = h1.Write([]byte("world."))
	if err != nil {
		panic(err)
	}
	fmt.Println(n, h1.Sum64())

	// h2使用与h1一样的种子（h1与h2将有一样的行为）
	h2.SetSeed(h1.Seed())
	n, err = h2.WriteString("hello world.")
	if err != nil {
		panic(err)
	}
	fmt.Println(n, h2.Sum64())

	// h1丢失之前写入到数据，但种子保持不变
	h1.Reset()

	// h1、h2重新计算发现此时两者结果将不一样
	h1.WriteString("same")
	h2.WriteString("same")
	fmt.Println(h1.Sum64(), h2.Sum64())
	// output:
	// 6 10529433727409431631
	// 6 1572624585198278501
	// 12 1572624585198278501
	// 12110176567285930650 2656311455508479395
}
