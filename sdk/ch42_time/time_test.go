package ch42_time

// time：提供测量和显示时间功能（历法计算始终采用公历，没有闰秒）

// 常量
// 1.时间格式化：layout
// const (
// 	ANSIC       = "Mon Jan _2 15:04:05 2006"
// 	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// 	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// 	RFC822      = "02 Jan 06 15:04 MST"
// 	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// 	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// 	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// 	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// 	RFC3339     = "2006-01-02T15:04:05Z07:00"
// 	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// 	Kitchen     = "3:04PM"
// 	// Handy time stamps.
// 	Stamp      = "Jan _2 15:04:05"
// 	StampMilli = "Jan _2 15:04:05.000"
// 	StampMicro = "Jan _2 15:04:05.000000"
// 	StampNano  = "Jan _2 15:04:05.000000000"
// )

// 2.持续时间单位：time.Duration
// const (
// 	Nanosecond  Duration = 1
// 	Microsecond          = 1000 * Nanosecond
// 	Millisecond          = 1000 * Microsecond
// 	Second               = 1000 * Millisecond
// 	Minute               = 60 * Second
// 	Hour                 = 60 * Minute
// )

// 3.月份：time.Month
// const (
// 	January Month = 1 + iota
// 	February
// 	March
// 	April
// 	May
// 	June
// 	July
// 	August
// 	September
// 	October
// 	November
// 	December
// )

// 4.周几：time.Weekday
// const (
// 	Sunday Weekday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )

// 函数
// func After(d Duration) <-chan Time // 等过去时间d，然后在返回的通道发送当前时间。相当于NewTimer(d).C。
// func Sleep(d Duration)             // 当前goroutine暂停时间d，如果d为负数或0则立即返回。
// func Tick(d Duration) <-chan Time  // NewTicker 的便捷包装器，仅提供对滴答通道的访问。

// 类型
// 1.时间间隔：将两个瞬间之间经过的时间表示为 int64 纳秒计数。该表示将最大可表示持续时间限制为大约 290 年。
// type Duration int64
// func ParseDuration(s string) (Duration, error) // 字符串解析为duration
// func Since(t Time) Duration                    // 等价于time.Now().Sub(t)，返回当前时间减去时间t的间隔
// func Until(t Time) Duration                    // 等价于t.Sub(time.Now))，返回t减去当前时间的间隔
// func (d Duration) Hours() float64
// func (d Duration) Microseconds() int64
// func (d Duration) Milliseconds() int64
// func (d Duration) Minutes() float64
// func (d Duration) Nanoseconds() int64
// func (d Duration) Round(m Duration) Duration // 时间间隔d以m为单位四舍五入
// func (d Duration) Seconds() float64
// func (d Duration) String() string               // 以“72h3m0.5s”的形式返回表示持续时间的字符串
// func (d Duration) Truncate(m Duration) Duration // 时间间隔d以m为单位截断

// 2.位置：位置将时间映射为当前使用区域
// type Location struct {
// 	// contains filtered or unexported fields
// }
// var Local *Location = &localLoc // 本地位置
// var UTC *Location = &utcLoc // utc位置
// func FixedZone(name string, offset int) *Location // 根据区域名称name和偏移量（UTC以东的秒数），返回位置
// func LoadLocation(name string) (*Location, error) // 根据给定名称name返回其位置。
// 1.如果name为空或"UTC"则返回UTC位置，如果name是"Local"则返回本地位置，否则将被与IANA时区数据库对应的位置（例如"America/New_York"）
// 2.所有系统都不存在LoadLocation所需的时区数据库，尤其是非UNIX系统。LoadLocation在ZoneInfo环境变量中指定的目录或未压缩的zip文件中，如果有的话，然后在UNIX系统上查看已知的安装位置，最后查找$GOROOT/lib/time/dioundfo.zip。
// func LoadLocationFromTZData(name string, data []byte) (*Location, error) // 从名称name和IANA数据库数据获取位置
// func (l *Location) String() string // 返回位置的时区信息

// 3.Month：月份
// type Month int
// func (m Month) String() string

// 4.ParseError：解析时间字符串的问题
// type ParseError struct {
// 	Layout     string
// 	Value      string
// 	LayoutElem string
// 	ValueElem  string
// 	Message    string
// }
// func (e *ParseError) Error() string

// 5.Ticker：持有一个channel，时钟刻度
// type Ticker struct {
// 	C <-chan Time // The channel on which the ticks are delivered.
// 	// contains filtered or unexported fields
// }
// func NewTicker(d Duration) *Ticker // 根据时间间隔d返回一Ticker，该Ticker每过指定时间发送时间channel
// func (t *Ticker) Reset(d Duration) // 暂停t并重置其周期为指定时间间隔d，下一次触发在新的周期过去
// func (t *Ticker) Stop()            // 关闭t，tick将不再发送，但不会关闭channel，以防止并发的goroutine读取错误

// 6.Time：以纳秒为精度的时间
// type Time struct {
// 	// contains filtered or unexported fields
// }
// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time // 指定参数初始化时间
// func Now() Time                                                                     // 返回当前时间
// func Parse(layout, value string) (Time, error)                                      // 基于指定格式layout和时间字符串value解析，返回时间和错误
// func ParseInLocation(layout, value string, loc *Location) (Time, error)             // 以指定时区解析时间
// func Unix(sec int64, nsec int64) Time                                               // 自1970年1月1日起返回与给定的UNIX时间，SEC秒和NSEC纳秒相对应的本地时间。（sec和nsec不能重复）
// func (t Time) Add(d Duration) Time                                                  // 返回t增加间隔d后的时间
// func (t Time) AddDate(years int, months int, days int) Time                         // 返回t增加指定年月日后的时间
// func (t Time) After(u Time) bool                                                    // 报告t是否在d之后
// func (t Time) AppendFormat(b []byte, layout string) []byte                          // 在b基础上，返回增加以layout格式的时间字符串后的字节切片
// func (t Time) Before(u Time) bool                                                   // 报告t是否在u之后
// func (t Time) Clock() (hour, min, sec int)                                          // 返回t的时分秒
// func (t Time) Date() (year int, month Month, day int)                               // 返回t的年月日
// func (t Time) Day() int                                                             // 返回t的月份的天
// func (t Time) Equal(u Time) bool                                                    // 报告t和u时间是否相等，不同时区也可以相等
// func (t Time) Format(layout string) string                                          // 以指定格式layout格式化时间t
// func (t *Time) GobDecode(data []byte) error                                         // 实现了GobDecoder接口
// func (t Time) GobEncode() ([]byte, error)                                           // 实现了GobEncoder接口
// func (t Time) Hour() int                                                            // 返回t的时（范围0到23）
// func (t Time) ISOWeek() (year, week int)                                            // 返回t的ISO8601年月和周数
// func (t Time) In(loc *Location) Time                                                // 返回t的副本示例，该示例的时区为loc
// func (t Time) IsZero() bool                                                         // 报告t是否为零时间：January 1, year 1, 00:00:00 UTC
// func (t Time) Local() Time                                                          // 返回t以本地时区的时间
// func (t Time) Location() *Location                                                  // 返回t的时区信息
// func (t Time) MarshalBinary() ([]byte, error)                                       // 实现了Binary.Marshaler接口
// func (t Time) MarshalJSON() ([]byte, error)                                         // 实现了json.Marshaler接口，格式为RFC 3339
// func (t Time) MarshalText() ([]byte, error)                                         // 实现了encoding.TextMarshaler接口
// func (t Time) Minute() int                                                          // 返回t的分钟数（范围0到59）
// func (t Time) Month() Month                                                         // 返回t的月份
// func (t Time) Nanosecond() int                                                      // 返回t的秒单位的纳秒数（范围0到999999999）
// func (t Time) Round(d Duration) Time                                                // 返回时间t以间隔单位d四舍五入后的时间
// func (t Time) Second() int                                                          // 返回t的秒单位数（范围0到59）
// func (t Time) String() string                                                       // 返回t以指定格式的字符串："2006-01-02 15:04:05.999999999 -0700 MST"
// func (t Time) Sub(u Time) Duration                                                  // 返回时间t减去时间u的时间间隔
// func (t Time) Truncate(d Duration) Time                                             // 返回t以间隔单位d截断后的时间
// func (t Time) UTC() Time                                                            // 返回t以UTC时区的时间
// func (t Time) Unix() int64                                                          // 返回t的秒时间戳
// func (t Time) UnixNano() int64                                                      // 返回t的纳秒时间戳
// func (t *Time) UnmarshalBinary(data []byte) error                                   // 实现接口：encoding.BinaryUnmarshaler
// func (t *Time) UnmarshalJSON(data []byte) error                                     // 实现接口：json.Unmarshaler
// func (t *Time) UnmarshalText(data []byte) error                                     // 实现接口：encoding.TextUnmarshaler
// func (t Time) Weekday() Weekday                                                     // 返回t的Weekday
// func (t Time) Year() int                                                            // 返回t的年份数
// func (t Time) YearDay() int                                                         // 返回t的一年中第几天（范围1到365或1到366）
// func (t Time) Zone() (name string, offset int)                                      // 返回t的时区名称和偏移量

// 7.Timer：表示一个事件，当定时器到期，将会在C上发送当前时间（除非计时器是由AfterFunc创建的。必须使用newtimer或afterfunc创建计时器。）
// type Timer struct {
// 	C <-chan Time
// 	// contains filtered or unexported fields
// }
// func AfterFunc(d Duration, f func()) *Timer // 等待过去的持续时间，然后在自己的Goroutine中调用f。它返回一个可用于使用其停止方法取消呼叫的计时器。
// func NewTimer(d Duration) *Timer            // 创建一定时时间为d的定时器
// func (t *Timer) Reset(d Duration) bool      // 改变定时器t，使得在时间间隔d过后失效，定时器已激活返回true，定时器已过期或被暂停返回false
// func (t *Timer) Stop() bool                 // 暂停定时器t，如果调用stops返回true，如果定时器已过期或暂停则返回false

// 8.Weekday：指定一周的一天
// type Weekday int
// func (d Weekday) String() string
