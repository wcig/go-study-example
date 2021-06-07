package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// encoding/csc: 构造和解析csv文件（CSV文件有很多种；这个包支持 RFC 4180 中描述的格式）

// csv定义错误常量
func TestErr(t *testing.T) {
	_ = csv.ErrQuote      // 引号字段中多余或丢失的\"
	_ = csv.ErrBareQuote  // 非引号字段出现裸的\"
	_ = csv.ErrFieldCount // 错误字段数,调用Read()方法可能出现
}

// func NewWriter(w io.Writer) *Writer
// 返回一个新的Writer
func TestNewWriter(t *testing.T) {
	writer := csv.NewWriter(os.Stdout)
	assert.NotNil(t, writer)
}

// func (w *Writer) Write(record []string) error
// Write将单个CSV记录连同任何必要的引号一起写入w。 记录是一片字符串，每个字符串是一个字段。 写操作被缓冲，因此必须最终调用Flush以确保将记录写到基础io.Writer。
func TestWrite(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	file, err := os.Create("writer.csv")
	assert.Nil(t, err)

	writer := csv.NewWriter(file)
	for _, record := range records {
		err := writer.Write(record)
		assert.Nil(t, err)
	}

	writer.Flush()
	err = writer.Error()
	assert.Nil(t, err)
}

// func (w *Writer) WriteAll(records [][]string) error
// WriteAll使用Write将多个CSV记录写入w，然后调用Flush，并从Flush返回任何错误。
func TestWriteAll(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	file, err := os.Create("writer.csv")
	assert.Nil(t, err)

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	assert.Nil(t, err)

	err = writer.Error()
	assert.Nil(t, err)
}

// func NewReader(r io.Reader) *Reader
// 返回一个新的Writer
func TestNewReader(t *testing.T) {
	reader := csv.NewReader(os.Stdin)
	assert.NotNil(t, reader)
}

// func (r *Reader) Read() (record []string, err error)
// 读取从r读取一条记录（一片字段）。 如果记录中的字段数量超出预期，则Read会返回记录以及错误ErrFieldCount。
// 除这种情况外，Read总是返回一个非空记录或一个非空错误，但不会同时返回两者。 如果没有要读取的数据，则Read返回nil，即io.EOF。
// 如果ReuseRecord为true，则可以在多个Read调用之间共享返回的slice。
func TestRead(t *testing.T) {
	initCsvFile(t)

	file, err := os.Open("writer.csv")
	assert.Nil(t, err)

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		assert.Nil(t, err)
		fmt.Println(record, reader.ReuseRecord)
	}
	// output:
	// [first_name last_name username] false
	// [Rob Pike rob] false
	// [Ken Thompson ken] false
	// [Robert Griesemer gri] false
}

// func (r *Reader) ReadAll() (records [][]string, err error)
// ReadAll从r读取所有剩余的记录。 每个记录都是一片字段。 成功的调用将返回err == nil，而不是err == io.EOF。
// 因为ReadAll被定义为在EOF之前读取，所以它不会将文件结尾视为要报告的错误。
func TestReadAll(t *testing.T) {
	initCsvFile(t)

	file, err := os.Open("writer.csv")
	assert.Nil(t, err)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	assert.Nil(t, err)
	for _, record := range records {
		fmt.Println(record)
	}
	// output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}

// 读取其他配置类型的csv文件
func TestReaderWithOptions(t *testing.T) {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	err := ioutil.WriteFile("tmp.csv", []byte(in), os.ModePerm)
	assert.Nil(t, err)
	file, err := os.Open("tmp.csv")
	assert.Nil(t, err)

	reader := csv.NewReader(file)
	reader.Comma = ';'   // 字段分隔符
	reader.Comment = '#' // 注释符
	records, err := reader.ReadAll()
	assert.Nil(t, err)
	for _, record := range records {
		fmt.Println(record)
	}
	// output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}

func initCsvFile(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	file, err := os.Create("writer.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err = writer.WriteAll(records); err != nil {
		panic(err)
	}

	if err = writer.Error(); err != nil {
		panic(err)
	}
}
