package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/RoaringBitmap/roaring"
)

func main() {
	// 创建一个空的Roaring Bitmap
	rb := roaring.New()

	// 向Roaring Bitmap中添加元素
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)

	// 检查元素是否存在
	fmt.Println(rb.Contains(2)) // 输出: true

	// 迭代Roaring Bitmap中的元素
	fmt.Println("Roaring Bitmap的元素：")
	iterator := rb.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

	// 从Roaring Bitmap中删除元素
	rb.Remove(2)

	// 获取Roaring Bitmap的基本信息
	fmt.Println("Roaring Bitmap的基本信息：")
	fmt.Println("元素数量：", rb.GetCardinality())
	fmt.Println("最小值：", rb.Minimum())
	fmt.Println("最大值：", rb.Maximum())

	// 从数组创建Roaring Bitmap
	arr := []uint32{4, 5, 6}
	rb2 := roaring.BitmapOf(arr...)
	fmt.Println(rb2.String())

	// 执行Roaring Bitmap的集合操作（例如并集、交集、差集）
	rb.Or(rb2)
	fmt.Println("并集：", rb.ToArray())

	rb.And(rb2)
	fmt.Println("交集：", rb.ToArray())

	rb.AndNot(rb2)
	fmt.Println("差集：", rb.ToArray())

	// 序列化和反序列化
	buf := &bytes.Buffer{}
	n, err := rb2.WriteTo(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("rb2成功写入:", n)

	rb3 := roaring.New()
	p, err := rb3.ReadFrom(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("rb3成功读取:", p)
	fmt.Println(rb3.ToArray())

	// Output:
	// true
	// Roaring Bitmap的元素：
	// 1
	// 2
	// 3
	// Roaring Bitmap的基本信息：
	// 元素数量： 2
	// 最小值： 1
	// 最大值： 3
	// {4,5,6}
	// 并集： [1 3 4 5 6]
	// 交集： [4 5 6]
	// 差集： []
	// rb2成功写入: 22
	// rb3成功读取: 22
	// [4 5 6]
}
