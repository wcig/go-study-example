package template

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func TestHTMLEscape(t *testing.T) {
	src := `<div>'a&b"</div>`

	var buf bytes.Buffer
	template.HTMLEscape(&buf, []byte(src))
	fmt.Println(buf.String())

	result := template.HTMLEscapeString(src)
	fmt.Println(result)

	result = template.HTMLEscaper(src, "hello>world.")
	fmt.Println(result)
	// output:
	// &lt;div&gt;&#39;a&amp;b&#34;&lt;/div&gt;
	// &lt;div&gt;&#39;a&amp;b&#34;&lt;/div&gt;
	// &lt;div&gt;&#39;a&amp;b&#34;&lt;/div&gt;hello&gt;world.
}

func TestIsTrue(t *testing.T) {
	truth, ok := template.IsTrue(false)
	fmt.Println(truth, ok)

	truth, ok = template.IsTrue("ok")
	fmt.Println(truth, ok)

	truth, ok = template.IsTrue("")
	fmt.Println(truth, ok)
	// output:
	// false true
	// true true
	// false true
}

func TestJSEscape(t *testing.T) {
	src := `'a&b"`

	var buf bytes.Buffer
	template.JSEscape(&buf, []byte(src))
	fmt.Println(buf.String())

	result := template.JSEscapeString(src)
	fmt.Println(result)

	result = template.JSEscaper(src, "hello>world.")
	fmt.Println(result)
	// output:
	// \'a\u0026b\"
	// \'a\u0026b\"
	// \'a\u0026b\"hello\u003Eworld.
}

func TestURLQueryEscaper(t *testing.T) {
	result := template.URLQueryEscaper("a&b", "hello>world.")
	fmt.Println(result) // a%26bhello%3Eworld.
}
