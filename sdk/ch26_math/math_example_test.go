package ch26_math

import (
	"fmt"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	x := math.Abs(-2)
	y := math.Abs(2)
	fmt.Printf("%.1f, %.1f\n", x, y) // 2.0, 2.0
}

func TestAcos(t *testing.T) {
	fmt.Printf("%.2f\n", math.Acos(1)) // 0.00
}

func TestAcosh(t *testing.T) {
	fmt.Printf("%.2f\n", math.Acosh(1)) // 0.00
}

func TestAsin(t *testing.T) {
	fmt.Printf("%.2f\n", math.Asin(0)) // 0.00
}

func TestAtan(t *testing.T) {
	fmt.Printf("%.2f\n", math.Atan(0)) // 0.00
}

func TestAtan2(t *testing.T) {
	fmt.Printf("%.2f\n", math.Atan2(0, 0)) // 0.00
}

func TestAtanh(t *testing.T) {
	fmt.Printf("%.2f\n", math.Atanh(0)) // 0.00
}

func TestCbrt(t *testing.T) {
	fmt.Printf("%.2f\n", math.Cbrt(8))  // 2.00
	fmt.Printf("%.2f\n", math.Cbrt(27)) // 3.00
}

func TestCeil(t *testing.T) {
	fmt.Printf("%.2f\n", math.Ceil(1.49)) // 2.00
}

func TestCopySign(t *testing.T) {
	fmt.Printf("%.2f\n", math.Copysign(3.21, -1)) // -3.21
}

func TestCos(t *testing.T) {
	fmt.Printf("%.2f\n", math.Cos(math.Pi/2)) // 0.00
}

func TestCosh(t *testing.T) {
	fmt.Printf("%.2f\n", math.Cosh(0)) // 1.00
}

func TestDim(t *testing.T) {
	fmt.Printf("%.2f\n", math.Dim(4, -2)) // 6.00
	fmt.Printf("%.2f\n", math.Dim(-4, 2)) // 0.00
}

func TestExp(t *testing.T) {
	fmt.Printf("%.2f\n", math.Exp(1)) // 2.72
}

func TestExp2(t *testing.T) {
	fmt.Printf("%.2f\n", math.Exp2(1))  // 2.00
	fmt.Printf("%.2f\n", math.Exp2(-1)) // 0.50
}

func TestExpm1(t *testing.T) {
	fmt.Printf("%.2f\n", math.Expm1(1)) // 1.72
}

func TestFMA(t *testing.T) {
	fmt.Printf("%.2f\n", math.FMA(2, 3, 4)) // 10
}

func TestFloor(t *testing.T) {
	fmt.Printf("%.2f\n", math.Floor(1.51)) // 1.00
}

func TestIlogb(t *testing.T) {
	fmt.Println(math.Ilogb(2)) // 1
	fmt.Println(math.Ilogb(8)) // 3
}

func TestLog(t *testing.T) {
	fmt.Printf("%.1f\n", math.Log(1))      // 0.0
	fmt.Printf("%.1f\n", math.Log(2.7183)) // 1.0
}

func TestLog10(t *testing.T) {
	fmt.Printf("%.1f\n", math.Log10(100)) // 2.0
}

func TestLog2(t *testing.T) {
	fmt.Printf("%.1f\n", math.Log2(8)) // 3.0
}

func TestLogb(t *testing.T) {
	fmt.Printf("%.1f\n", math.Logb(16)) // 3.0
}

func TestMax(t *testing.T) {
	fmt.Printf("%.2f\n", math.Max(1.11, 2.22)) // 2.22
}

func TestMin(t *testing.T) {
	fmt.Printf("%.2f\n", math.Min(1.11, 2.22)) // 1.11
}

func TestMod(t *testing.T) {
	fmt.Printf("%.2f\n", math.Mod(7, 4)) // 3.00
}

func TestModf(t *testing.T) {
	int, frac := math.Modf(3.14)
	fmt.Printf("%.2f, %.2f\n", int, frac) // 3.00, 0.14

	int, frac = math.Modf(-2.71)
	fmt.Printf("%.2f, %.2f\n", int, frac) // -2.00, -0.71
}

func TestPow(t *testing.T) {
	fmt.Printf("%.2f\n", math.Pow(2, 3)) // 8.00
}

func TestPow10(t *testing.T) {
	fmt.Printf("%.2f\n", math.Pow10(2)) // 100.00
}

func TestRound(t *testing.T) {
	fmt.Printf("%.1f\n", math.Round(1.5))  // 2.0
	fmt.Printf("%.1f\n", math.Round(-1.5)) // -2.0
}

func TestRoundToEven(t *testing.T) {
	fmt.Printf("%.1f\n", math.RoundToEven(1.5)) // 2.0
	fmt.Printf("%.1f\n", math.RoundToEven(2.5)) // 2.0
}

func TestSignbit(t *testing.T) {
	fmt.Println(math.Signbit(-1)) // true
	fmt.Println(math.Signbit(-0)) // false
	fmt.Println(math.Signbit(0))  // false
	fmt.Println(math.Signbit(1))  // false
}

func TestTrunc(t *testing.T) {
	fmt.Printf("%.2f\n", math.Trunc(1.51))  // 1.00
	fmt.Printf("%.2f\n", math.Trunc(-1.51)) // -1.00
}
