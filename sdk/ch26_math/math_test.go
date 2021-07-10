package ch26_math

// math: 提供基本常量和数学函数（不保证跨架构位相同结果）

// 常量
// const (
//    E   = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113
//    Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
//    Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // https://oeis.org/A001622
//
//    Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
//    SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
//    SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
//    SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339
//
//    Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
//    Log2E  = 1 / Ln2
//    Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
//    Log10E = 1 / Ln10
// )
// const (
//    MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
//    SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)
//
//    MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
//    SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
// )
// const (
//    MaxInt8   = 1<<7 - 1
//    MinInt8   = -1 << 7
//    MaxInt16  = 1<<15 - 1
//    MinInt16  = -1 << 15
//    MaxInt32  = 1<<31 - 1
//    MinInt32  = -1 << 31
//    MaxInt64  = 1<<63 - 1
//    MinInt64  = -1 << 63
//    MaxUint8  = 1<<8 - 1
//    MaxUint16 = 1<<16 - 1
//    MaxUint32 = 1<<32 - 1
//    MaxUint64 = 1<<64 - 1
// )

// 函数
// func Abs(x float64) float64                       // 求x绝对值
// func Acos(x float64) float64                      // 求x的反余弦值（x为角度对应的余弦值）
// func Acosh(x float64) float64                     // 求x的反向双曲余弦（arcosh(x)）
// func Asin(x float64) float64                      // 求x的反正弦值（x为角度对应的正弦值）
// func Asinh(x float64) float64                     // 求x的反向双曲正弦（arsinh(x)）
// func Atan(x float64) float64                      // 求x的反正切值（x为角度对应的正切值）
// func Atan2(y, x float64) float64                  // 返回y/x的弧切，使用两个迹象来确定返回值的象限
// func Atanh(x float64) float64                     // 求x的反向双曲正切值（artanh(x)）
// func Cbrt(x float64) float64                      // 求x的立方根
// func Ceil(x float64) float64                      // 返回大于或等于x的最小整数值
// func Copysign(x, y float64) float64               // 返回具有x大小和y符号的值
// func Cos(x float64) float64                       // 求x的余弦值（x为角度）
// func Cosh(x float64) float64                      // 求x的双曲余弦值（cosh(x)）
// func Dim(x, y float64) float64                    // 返回x-y和0的最大值
// func Erf(x float64) float64                       // 返回x的误差函数值
// func Erfc(x float64) float64                      // 返回x的互补误差函数值
// func Erfcinv(x float64) float64                   // 返回Erfc(x)的倒数
// func Erfinv(x float64) float64                    // 返回x的逆误差函数值
// func Exp(x float64) float64                       // 返回e的x次方
// func Exp2(x float64) float64                      // 返回2的x次方
// func Expm1(x float64) float64                     // 返回e的x次方-1
// func FMA(x, y, z float64) float64                 // 返回x*y+z
// func Float32bits(f float32) uint32                // 返回f的 IEEE 754 二进制表示
// func Float32frombits(b uint32) float32            // 返回对应于 IEEE 754 二进制表示 b 的浮点数
// func Float64bits(f float64) uint64                // 返回f的IEEE 754二进制表示
// func Float64frombits(b uint64) float64            // 返回对应于 IEEE 754 二进制表示 b 的浮点数
// func Floor(x float64) float64                     // 返回小于或等于x的最大整数值
// func Frexp(f float64) (frac float64, exp int)     // 将 f 分解为归一化分数和 2 的整数幂（f = frac * 2的exp次方）
// func Gamma(x float64) float64                     // 返回x的gamma函数值
// func Hypot(p, q float64) float64                  // 返回sqrt(p*p, q*q)
// func Ilogb(x float64) int                         // 返回x的二进制指数（2的返回值次方=x）
// func Inf(sign int) float64                        // 如果sign>=0返回正无穷大+Inf，如果sign<0返回负无穷大-Inf
// func IsInf(f float64, sign int) bool              // 如果sign>0报告f是否为正无穷大，如果sign<0报告f是否为负无穷大，如果sign=0报告f是否为无穷大
// func IsNaN(f float64) (is bool)                   // 报告 f 是否为 IEEE 754 “非数字”值
// func J0(x float64) float64                        // 返回第一类零阶贝塞尔函数
// func J1(x float64) float64                        // 返回第一类一阶贝塞尔函数
// func Jn(n int, x float64) float64                 // 返回第一类 n 阶贝塞尔函数
// func Ldexp(frac float64, exp int) float64         // 返回Frexp的倒数（frac x 2**exp）
// func Lgamma(x float64) (lgamma float64, sign int) // 返回 Gamma(x) 的自然对数和符号（-1 或 +1）
// func Log(x float64) float64                       // 返回x的自然对数
// func Log10(x float64) float64                     // 返回x的以10位底的对数
// func Log1p(x float64) float64                     // 返回1加x的自然对数
// func Log2(x float64) float64                      // 返回x的以2位底的对数
// func Logb(x float64) float64                      // 返回x的以2位底的对数
// func Max(x, y float64) float64                    // 返回x、y的最大值
// func Min(x, y float64) float64                    // 返回x、y的最小值
// func Mod(x, y float64) float64                    // 返回x/y的余数
// func Modf(f float64) (int float64, frac float64)  // 返回f的整数部分和小数部分
// func NaN() float64                                // 返回一个 IEEE 754 “非数字”值
// func Nextafter(x, y float64) (r float64)          // 返回x之后朝向y的下一个科表示的float64值
// func Nextafter32(x, y float32) (r float32)        // 返回x之后朝向y的下一个科表示的float32值
// func Pow(x, y float64) float64                    // 返回x**y，即x的y次方
// func Pow10(n int) float64                         // 返回10**n，即10的n次方
// func Remainder(x, y float64) float64              // 返回 x/y 的 IEEE 754 浮点余数
// func Round(x float64) float64                     // 返回x的四舍五入
// func RoundToEven(x float64) float64               // 返回最接近x的整数，四舍五入到偶数
// func Signbit(x float64) bool                      // 报告x是负数还是负零
// func Sin(x float64) float64                       // 返回x的正弦值
// func Sincos(x float64) (sin, cos float64)         // 返回Sin(x)和Cos(x)
// func Sinh(x float64) float64                      // 返回x的双曲正弦值
// func Sqrt(x float64) float64                      // 返回x的平方根
// func Tan(x float64) float64                       // 返回x的正切值
// func Tanh(x float64) float64                      // 返回x的双曲正切值
// func Trunc(x float64) float64                     // 返回x的整数值
// func Y0(x float64) float64                        // 返回第二类零阶贝塞尔函数
// func Y1(x float64) float64                        // 返回第二类一阶贝塞尔函数
// func Yn(n int, x float64) float64                 // 返回第二类 n 阶贝塞尔函数
