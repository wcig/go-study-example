package rand

// math/rand: 实现了伪随机数生成器。

// 函数
// func ExpFloat64() float64                // 返回一个在 (0, +math.MaxFloat64] 范围内的指数分布的 float64 值
// func Float32() float32                   // 返回[0.0,1.0)的伪随机数float32值
// func Float64() float64                   // 返回[0.0,1.0)的伪随机数float64值
// func Int() int                           // 返回非负的伪随机int值
// func Int31() int32                       // 返回非负的伪随机的31位整数的int32值
// func Int31n(n int32) int32               // 返回[0,n)的伪随机int32值
// func Int63() int64                       // 返回非负的伪随机的63位整数的int64值
// func Int63n(n int64) int64               // 返回[0,n)的伪随机int64值
// func Intn(n int) int                     // 返回[0,n)的伪随机int值
// func NormFloat64() float64               // 返回范围为 [-math.MaxFloat64, +math.MaxFloat64] 的正态分布的 float64 值
// func Perm(n int) []int                   // 返回[0,n)的伪随机排列
// func Read(p []byte) (n int, err error)   // 从默认 Source 生成 len(p) 个随机字节并将它们写入 p。它总是返回 len(p) 和一个 nil 错误。与 Rand.Read 方法不同，Read 对于并发使用是安全的。
// func Seed(seed int64)                    // 使用提供的种子值将默认 Source 初始化为确定性状态。如果未调用 Seed，则生成器的行为就像由 Seed(1) 播种一样。与 Rand.Seed 方法不同，Seed 可以安全地并发使用。
// func Shuffle(n int, swap func(i, j int)) // 使用默认 Source 伪随机化元素的顺序。
// func Uint32() uint32                     // 以默认source返回一个伪随机uint32值
// func Uint64() uint64                     // 以默认source返回一个伪随机uint64值

// 类型
// 1.Rand: 随机数source
// type Rand struct {
//    // contains filtered or unexported fields
// }
// func New(src Source) *Rand                         // 指定src创建Rand
// func (r *Rand) ExpFloat64() float64                // 返回一个在 (0, +math.MaxFloat64] 范围内的指数分布的 float64，其指数分布的速率参数 (lambda) 为 1，平均值为 1/lambda (1)。
// func (r *Rand) Float32() float32                   // 返回[0.0,1.0)的伪随机数float32值
// func (r *Rand) Float64() float64                   // 返回[0.0,1.0)的伪随机数float64值
// func (r *Rand) Int() int                           // 返回非负的伪随机int值
// func (r *Rand) Int31() int32                       // 返回非负的伪随机的31位整数的int32值
// func (r *Rand) Int31n(n int32) int32               // 返回[0,n)的伪随机int32值
// func (r *Rand) Int63() int64                       // 返回非负的伪随机的63位整数的int64值
// func (r *Rand) Int63n(n int64) int64               // 返回[0,n)的伪随机int64值
// func (r *Rand) Intn(n int) int                     // 返回[0,n)的伪随机int值
// func (r *Rand) NormFloat64() float64               // 返回一个正态分布的 float64，范围为 -math.MaxFloat64 到 +math.MaxFloat64（含），具有标准正态分布（均值 = 0，stddev = 1）。
// func (r *Rand) Perm(n int) []int                   // 返回[0,n)的伪随机排列
// func (r *Rand) Read(p []byte) (n int, err error)   // 生成 len(p) 个随机字节并将它们写入 p。它总是返回 len(p) 和一个 nil 错误。 Read 不应与任何其他 Rand 方法同时调用。
// func (r *Rand) Seed(seed int64)                    // 使用提供的种子值将生成器初始化为确定性状态。不应与任何其他 Rand 方法同时调用 Seed。
// func (r *Rand) Shuffle(n int, swap func(i, j int)) // 伪随机化元素的顺序。
// func (r *Rand) Uint32() uint32                     // 返回一个伪随机uint32值
// func (r *Rand) Uint64() uint64                     // 返回一个伪随机uint64值

// 2.Source: 表示在 [0, 1<<63) 范围内均匀分布的伪随机 int64 值的源
// type Source interface {
//    Int63() int64
//    Seed(seed int64)
// }
// func NewSource(seed int64) Source // 返回以给定值作为种子的新伪随机源。与顶级函数使用的默认 Source 不同，此源对于多个 goroutine 并发使用是不安全的。

// 3.Source64: 是一个 Source，它也可以直接在 [0, 1<<64) 范围内生成均匀分布的伪随机 uint64 值。如果 Rand r 的底层 Source s 实现 Source64，则 r.Uint64 将返回一次调用 s.Uint64 的结果，而不是两次调用 s.Int63。
// type Source64 interface {
//    Source
//    Uint64() uint64
// }

// 4.Zipf: 生成 Zipf 分布式变量
// type Zipf struct {
//    // contains filtered or unexported fields
// }
// func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf // 返回一个 Zipf 变量生成器
// func (z *Zipf) Uint64() uint64                                 // 返回从 Zipf 对象描述的 Zipf 分布中提取的值
