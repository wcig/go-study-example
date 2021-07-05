package template

// text/template：实现数据驱动模板来生成文本输出。

// 函数
// func HTMLEscape(w io.Writer, b []byte)           // 将数据b的HTML转义写入w
// func HTMLEscapeString(s string) string           // 返回纯文本数据s的HTML转义
// func HTMLEscaper(args ...interface{}) string     // 返回多个参数HTML转义后拼接结果
// func IsTrue(val interface{}) (truth, ok bool)    // 返回val是否是其类型非零值和其值是否有意义的真值
// func JSEscape(w io.Writer, b []byte)             // 纯文本数据b的JS转义写入w
// func JSEscapeString(s string) string             // 返回纯文本数据s的JS转义
// func JSEscaper(args ...interface{}) string       // 返回多个参数JS转义后拼接结果
// func URLQueryEscaper(args ...interface{}) string // 返回多个参数url编码后拼接结果

// 1.ExecError：评估模板时的错误
// type ExecError struct {
//    Name string // Name of template.
//    Err  error  // Pre-formatted error.
// }
// func (e ExecError) Error() string
// func (e ExecError) Unwrap() error

// 2.FuncMap：名称到函数的映射
// type FuncMap map[string]interface{}

// 3.Template：表示以解析的模板，*parse.Tree字段仅供html/template使用，应被其他客户端视为不可导出的。
// type Template struct {
//    *parse.Tree
//    // contains filtered or unexported fields
// }
// func Must(t *Template, err error) *Template // html模板包装函数，err非空则导致panic
// func New(name string) *Template // 分配以参数名的html模板
// func ParseFS(fsys fs.FS, patterns ...string) (*Template, error) // 类似ParseFile和ParseGlob，从文件系统fs和指定模式glob patterns加载
// func ParseFiles(filenames ...string) (*Template, error) // 从多个文件加载为模板，返回模板的名称具有第一个文件的名称和内容
// func ParseGlob(pattern string) (*Template, error) // 从模式识别中解析模板
// func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error) // 使用名称和解析树创建一个新的模板，并使其与t关联。
// func (t *Template) Clone() (*Template, error) // 返回模板的副本，包括所有关联的模板。
// func (t *Template) DefinedTemplates() string // 返回一字符串，列出定义的模板。
// func (t *Template) Delims(left, right string) *Template // 将操作定界符设置为指定的字符串，以便在对 Parse、ParseFiles 或 ParseGlob 的后续调用中使用。
// func (t *Template) Execute(wr io.Writer, data interface{}) error // 将数据对象data应用于模板t，并将结果写入w
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error // 将数据对象data应用于具有指定名称且与t关联的模板，并将结果写入wr
// func (t *Template) Funcs(funcMap FuncMap) *Template // 将参数函数映射添加到模板t的函数映射中
// func (t *Template) Lookup(name string) *Template // 返回与t关联的指定名称的模板，没有则返回nil
// func (t *Template) Name() string // 返回模板t的名称
// func (t *Template) New(name string) *Template // 分配一指定名称的新模板，且该模板与t关联，有相同的分隔符
// func (t *Template) Option(opt ...string) *Template // 为模板t设定选项，可以是字符串也可以是"key=value"
// func (t *Template) Parse(text string) (*Template, error) // 将文本text解析为t的模板主体
// func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error) // 类似于ParseFiles或ParseGlob，但是从文件系统fs读取
// func (t *Template) ParseFiles(filenames ...string) (*Template, error) // 解析传入文件生成模板，并将其与t关联
// func (t *Template) ParseGlob(pattern string) (*Template, error) // 解析模式标识的文件生成模板，并将其与t关联
// func (t *Template) Templates() []*Template // 返回与t关联的所有模板
