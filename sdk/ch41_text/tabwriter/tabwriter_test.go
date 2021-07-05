package tabwriter

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

// text/tabwriter：实现了写过滤器（tabwriter.Writer），可将输入的选项列转换为正确对齐文本。

// 常量
// const (
// 	FilterHTML uint = 1 << iota
// 	StripEscape
// 	AlignRight
// 	DiscardEmptyColumns
// 	TabIndent
// 	Debug
// )
// const Escape = '\xff'

// 类型
// 1.Writer：写过滤器，它在其输入中以制表符分隔的列周围插入填充以在输出中对齐它们。
// type Writer struct {
// 	// contains filtered or unexported fields
// }
// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer        // 创建一Writer，参数与Init一样
// func (b *Writer) Flush() error                                                                             // 将缓存区数据写入输出中，赢在最后一次调用Write后调用Flush
// func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer // 初始化一Writer，output诶指定的过滤器输出，其余参数为控制格式
// func (b *Writer) Write(buf []byte) (n int, err error)                                                      // 写入buf数据到b中，返回唯一错误是底层写入时遇到错误

func TestElastic(t *testing.T) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()
	// output:
	// ....a|..b|c
	// ...aa|.bb|cc
	// ..aaa|
	// .aaaa|.dddd|eeee
}

func TestTrailingTab(t *testing.T) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // no trailing tab
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	w.Flush()
	// output:
	// ------a|------b|---aligned|
	// -----aa|-----bb|---aligned|
	// ----aaa|----bbb|unaligned
	// ---aaaa|---bbbb|---aligned|
}

func TestTypeWriter(t *testing.T) {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()
	// output:
	// a	b	c	d		.
	// 123	12345	1234567	123456789	.
	//
	//    a     b       c         d.
	//  123 12345 1234567 123456789.
}
