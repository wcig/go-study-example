package template

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// html/template
// 实现了数据驱动的模板，用于生成对代码注入安全的 HTML 输出。它提供与包 text/template 相同的接口，并且只要输出是 HTML，就应该使用它代替 text/template。

// func HTMLEscape(w io.Writer, b []byte)：将b的html转义后数据写入到w
// func HTMLEscapeString(s string) string：返回s的html转义结果
// func HTMLEscaper(args ...interface{}) string：返回多个参数html转义后拼接结果
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

// func IsTrue(val interface{}) (truth, ok bool)：返回val是否是其类型非零值和其值是否有意义的真值
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

// func JSEscape(w io.Writer, b []byte)：纯文本b的js转义结果写入w
// func JSEscapeString(s string) string：返回s的js转义结果
// func JSEscaper(args ...interface{}) string：返回多个参数的js转义拼接结果
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

// func URLQueryEscaper(args ...interface{}) string：url query的转义
func TestURLQueryEscaper(t *testing.T) {
	result := template.URLQueryEscaper("a&b", "hello>world.")
	fmt.Println(result) // a%26bhello%3Eworld.
}

// 类型：template.CSS
func TestTypeCSS(t *testing.T) {
	c := template.CSS("a")
	fmt.Println(c) // a
}

// 类型：template.Error
func TestTypeError(t *testing.T) {
	err := template.Error{
		ErrorCode:   template.ErrBadHTML,
		Node:        nil,
		Name:        "",
		Line:        0,
		Description: "",
	}
	fmt.Println(err) // {2 <nil>  0 }
}

// 类型：template.FuncMap 名称函数的映射
func TestTypeFuncMap(t *testing.T) {
	var m = template.FuncMap{"join": strings.Join}
	fmt.Println(len(m))
}

// 类型：template.Html html文档片段
func TestTypeHtml(t *testing.T) {
	var h = template.HTML("ok")
	fmt.Println(h) // ok
}

// 类型：template.HtmlAttr html属性
func TestTypeHtmlAttr(t *testing.T) {
	var ha = template.HTMLAttr(`type="text"`)
	fmt.Println(ha) // type="text"
}

// 类型：template.JS 封装一已知安全的ES5表达式
func TestTypeJS(t *testing.T) {
	var j = template.JS(`a=1`)
	fmt.Println(j) // a=1
}

// 类型：template.JSStr 封装了一系列字符，这些字符旨在嵌入 JavaScript 表达式中的引号之间
func TestTypeJSStr(t *testing.T) {
	var js = template.JSStr(`ok`)
	fmt.Println(js) // ok
}

// 类型：template.Srcset 已知安全的srcset属性
func TestTypeSrcset(t *testing.T) {
	var s = template.Srcset("srcset")
	fmt.Println(s) // srcset
}

// 类型：template.Template 用于生成安全的html文档片段
func TestTypeTemplate(t *testing.T) {
	_ = template.Template{}
}

// func Must(t *Template, err error) *Template：html模板包装函数，err非空则导致panic
func TestMust(t *testing.T) {
	t1 := template.Must(template.New("master").Parse("html"))
	assert.NotNil(t, t1)
}

// func New(name string) *Template：分配以参数名的html模板
func TestNew(t *testing.T) {
	t1 := template.New("master")
	assert.NotNil(t, t1)
}

// func ParseFS(fs fs.FS, patterns ...string) (*Template, error)：类似ParseFile和ParseGlob，从文件系统fs和指定模式glob patterns加载
func TestParseFS(t *testing.T) {
	fs := os.DirFS("testdata")
	t1, err := template.ParseFS(fs, "DOES NOT EXIST")
	fmt.Println(t1, err)
}

// func ParseFiles(filenames ...string) (*Template, error)：从多个文件加载为模板，返回模板的名称具有第一个文件的名称和内容
func TestParseFiles(t *testing.T) {
	t1, err := template.ParseFiles("testdata/file1.tmpl")
	assert.Nil(t, err)
	assert.NotNil(t, t1)
}

// func ParseGlob(pattern string) (*Template, error)：从模式识别中解析模板
func TestParseGlob(t *testing.T) {
	t1, err := template.ParseGlob("testdata/*.tmpl")
	assert.Nil(t, err)
	assert.NotNil(t, t1)
}

// template.Template方法：
// 1.使用名称和解析树创建一个新的模板，并使其与t关联。
// func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)
// 2.返回模板的副本，包括所有关联的模板。
// func (t *Template) Clone() (*Template, error)
// 3.返回一字符串，列出定义的模板。
// func (t *Template) DefinedTemplates() string
// 4.将操作定界符设置为指定的字符串，以便在对 Parse、ParseFiles 或 ParseGlob 的后续调用中使用。
// func (t *Template) Delims(left, right string) *Template
// 5.将数据对象data应用于模板t，并将结果写入w
// func (t *Template) Execute(wr io.Writer, data interface{}) error
// 6.将数据对象data应用于具有指定名称且与t关联的模板，并将结果写入wr
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
// 7.将参数函数映射添加到模板t的函数映射中
// func (t *Template) Funcs(funcMap FuncMap) *Template
// 8.返回与t关联的指定名称的模板，没有则返回nil
// func (t *Template) Lookup(name string) *Template
// 9.返回模板t的名称
// func (t *Template) Name() string
// 10.分配一指定名称的新模板，且该模板与t关联，有相同的分隔符
// func (t *Template) New(name string) *Template
// 11.为模板t设定选项，可以是字符串也可以是"key=value"
// func (t *Template) Option(opt ...string) *Template
// 12.将文本text解析为t的模板主体
// func (t *Template) Parse(text string) (*Template, error)
// 13.类似于ParseFiles或ParseGlob，但是从文件系统fs读取
// func (t *Template) ParseFS(fs fs.FS, patterns ...string) (*Template, error)
// 14.解析传入文件生成模板，并将其与t关联
// func (t *Template) ParseFiles(filenames ...string) (*Template, error)
// 15.解析模式标识的文件生成模板，并将其与t关联
// func (t *Template) ParseGlob(pattern string) (*Template, error)
// 16.返回与t关联的所有模板
// func (t *Template) Templates() []*Template

// 示例
func TestExample(t *testing.T) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	tt, err := template.New("webpage").Parse(tpl)
	check(err)

	fmt.Println("template tt name:", tt.Name())
	fmt.Println("defined templates:", tt.DefinedTemplates())

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = tt.Execute(os.Stdout, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = tt.Execute(os.Stdout, noItems)
	check(err)
	// output:
	// template tt name: webpage
	// defined templates: ; defined templates are: "webpage"
	//
	// <!DOCTYPE html>
	// <html>
	//	<head>
	//		<meta charset="UTF-8">
	//		<title>My page</title>
	//	</head>
	//	<body>
	//		<div>My photos</div><div>My blog</div>
	//	</body>
	// </html>
	// <!DOCTYPE html>
	// <html>
	//	<head>
	//		<meta charset="UTF-8">
	//		<title>My another page</title>
	//	</head>
	//	<body>
	//		<div><strong>no rows</strong></div>
	//	</body>
	// </html>--- PASS: TestExample
}
