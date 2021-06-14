package ch20_html

import (
	"fmt"
	"html"
	"testing"
)

// html：提供html的转义和取消转义功能

// func EscapeString(s string) string：转义5种特殊字符：>, <, &, ', "
func TestEscapeString(t *testing.T) {
	src := `<div>'a&b"</div>`
	result := html.EscapeString(src)
	fmt.Println(result) // &lt;div&gt;&#39;a&amp;b&#34;&lt;/div&gt;
}

// func UnescapeString(s string) string：反转义指定字符
func TestUnescapeString(t *testing.T) {
	src := `<div>'a&b"</div>`
	val := html.EscapeString(src)
	result := html.UnescapeString(val)
	fmt.Println(result) // <div>'a&b"</div>
}
