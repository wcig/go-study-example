package ch28_json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// interface -> json string
func TestInterfaceToJsonString(t *testing.T) {
	u := &user{Id: 1, Name: "tom"}
	b, err := json.Marshal(u)
	assert.Nil(t, err)

	jsonStr := string(b)
	fmt.Println("json str:", jsonStr)
}

// output:
// json str: {"id":1,"name":"tom"}

// json string -> interface
func TestJsonStringToInterface(t *testing.T) {
	jsonStr := `{"id":1,"name":"tom"}`
	var u user
	err := json.Unmarshal([]byte(jsonStr), &u)
	assert.Nil(t, err)

	fmt.Println("user:", u)
}

// output:
// user: {1 tom}

// 格式化输出json
func TestPrettyJsonString(t *testing.T) {
	u := &user{Id: 1, Name: "tom"}
	b, err := json.MarshalIndent(u, "", "\t")
	assert.Nil(t, err)

	jsonStr := string(b)
	fmt.Printf("pretty json str:\n%s\n", jsonStr)
}

// output:
// pretty json str:
// {
//	"id": 1,
//	"name": "tom"
// }

// json转换: 带html标签
func TestJsonWithHtmlTag(t *testing.T) {
	u1 := &user{Id: 1, Name: "<div>tom</div>"}
	b, err := json.Marshal(u1)
	assert.Nil(t, err)

	jsonStr := string(b)
	fmt.Println("json str:", jsonStr)

	var u2 user
	err = json.Unmarshal([]byte(jsonStr), &u2)
	assert.Nil(t, err)
	fmt.Println("u2:", u2)
}

// output:
// json str: {"id":1,"name":"\u003cdiv\u003etom\u003c/div\u003e"}
// u2: {1 <div>tom</div>}

// json.HTMLEscape: html特殊字符处理
func TestJsonHTMLEscape(t *testing.T) {
	var b, want bytes.Buffer
	jsonStr := `{"id":1,"name":"<div>tom</div>"}`
	want.WriteString(`{"id":1,"name":"\u003cdiv\u003etom\u003c/div\u003e"}`)

	json.HTMLEscape(&b, []byte(jsonStr))
	assert.True(t, bytes.Equal(b.Bytes(), want.Bytes()))
	fmt.Println("html escape result:", string(b.Bytes()))
}

// interface -> json io.Writer
func TestInterfaceToIoWriter(t *testing.T) {
	var (
		b   bytes.Buffer
		err error
	)

	u := &user{Id: 1, Name: "<div>tom</div>"}
	encoder := json.NewEncoder(&b)
	err = encoder.Encode(u)
	assert.Nil(t, err)
	fmt.Printf("default json str:\n%s\n", string(b.Bytes()))

	b.Reset()
	encoder.SetIndent("", "\t")
	err = encoder.Encode(u)
	assert.Nil(t, err)
	fmt.Printf("after set indent, json str:\n%s\n", string(b.Bytes()))

	b.Reset()
	encoder.SetEscapeHTML(false) // 默认true
	err = encoder.Encode(u)
	assert.Nil(t, err)
	fmt.Printf("after set escape html false, json str:\n%s\n", string(b.Bytes()))
}

// output:
// default json str:
// {"id":1,"name":"\u003cdiv\u003etom\u003c/div\u003e"}
//
// after set indent, json str:
// {
//	"id": 1,
//	"name": "\u003cdiv\u003etom\u003c/div\u003e"
// }
//
// after set escape html false, json str:
// {
//	"id": 1,
//	"name": "<div>tom</div>"
// }
//

// json io.Reader -> interface
func TestIoWriterToInterface(t *testing.T) {
	var (
		b   bytes.Buffer
		err error
	)

	b.WriteString(`{"id":1,"name":"\u003cdiv\u003etom\u003c/div\u003e"}`)
	decoder := json.NewDecoder(&b)
	var u user
	err = decoder.Decode(&u)
	assert.Nil(t, err)
	fmt.Println("user:", u)
}

// output:
// user: {1 <div>tom</div>}

// 校验是否json字符串
func TestValidJsonString(t *testing.T) {
	var isJson bool
	isJson = json.Valid([]byte(`{"id":1,"name":"tom"}`))
	assert.Equal(t, true, isJson)

	isJson = json.Valid([]byte(`{"id":1,"name":`))
	assert.Equal(t, false, isJson)
}

// json字符串压缩
func TestJsonStringCompact(t *testing.T) {
	var (
		b   bytes.Buffer
		err error
	)

	srcStr := `{
	"id": 1,
	"name": "tom"
}`
	expectedStr := `{"id":1,"name":"tom"}`
	err = json.Compact(&b, []byte(srcStr))
	assert.Nil(t, err)
	assert.Equal(t, expectedStr, string(b.Bytes()))
}

// json字符串美化
func Test(t *testing.T) {
	var (
		b   bytes.Buffer
		err error
	)

	srcStr := `{"id":1,"name":"tom"}`
	expectedStr := `{
	"id": 1,
	"name": "tom"
}`
	err = json.Indent(&b, []byte(srcStr), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedStr, string(b.Bytes()))
}
