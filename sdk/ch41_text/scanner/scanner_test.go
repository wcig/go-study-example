package scanner

// text/scanner：UTF-8文本扫描

// 常量
// const (
//    ScanIdents     = 1 << -Ident
//    ScanInts       = 1 << -Int
//    ScanFloats     = 1 << -Float // includes Ints and hexadecimal floats
//    ScanChars      = 1 << -Char
//    ScanStrings    = 1 << -String
//    ScanRawStrings = 1 << -RawString
//    ScanComments   = 1 << -Comment
//    SkipComments   = 1 << -skipComment // if set with ScanComments, comments become white space
//    GoTokens       = ScanIdents | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
// )

// const (
//    EOF = -(iota + 1)
//    Ident
//    Int
//    Float
//    Char
//    String
//    RawString
//    Comment
// )

// const GoWhitespace = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '

// 函数
// func TokenString(tok rune) string // 返回一可打印字符串，用于token或unicode字符

// 类型
// 1.Position：表示位置值，如果Line>0则位置有效
// type Position struct {
// 	Filename string // filename, if any
// 	Offset   int    // byte offset, starting at 0
// 	Line     int    // line number, starting at 1
// 	Column   int    // column number, starting at 1 (character count per line)
// }
// func (pos *Position) IsValid() bool // 报告位置是否有效
// func (pos Position) String() string // 位置字符串格式

// 2.Scanner：从io.Reader实现读取unicode字符和token
// type Scanner struct {
// 	Error       func(s *Scanner, msg string)
// 	ErrorCount  int
// 	Mode        uint
// 	Whitespace  uint64
// 	IsIdentRune func(ch rune, i int) bool // Go 1.4
// 	Position
// 	// contains filtered or unexported fields
// }
// func (s *Scanner) Init(src io.Reader) *Scanner // 从src初始化一扫描器，其Error为nil，ErrorCount为0，Model为GoTokens，Whitespace为GoWhitespace
// func (s *Scanner) Next() rune                  // 读取并返回下一个unicode字符，到源末尾返回EOF
// func (s *Scanner) Peek() rune                  // 返回源的下一个unicode字符，但不推进扫描器，如果扫描器的位置在源的最后一个字符，则返回EOF
// func (s *Scanner) Pos() (pos Position)         // 返回最后一次调用Next或Scan后立即返回的字符或token的位置，使用扫描器的Position字段以获取最近赛欧美token的其实位置
// func (s *Scanner) Scan() rune                  // 读取并返回源中的下一个token或unicode字符
// func (s *Scanner) TokenText() string           // 返回最近扫描的token对应的字符串，在调用Scan和Scanner.Error调用中有效。
