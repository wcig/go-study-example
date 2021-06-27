package ch35_sort

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// sort: 提供对切片和用户自定义切片的排序

// 函数:
// func Float64s(x []float64): float64类型切片x按升序排序
func TestFloat64s(t *testing.T) {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s)
	fmt.Println(s)
	// output:
	// [-3.8 -1.3 0.7 2.6 5.2]
	// [NaN -Inf 0 +Inf]
}

// func Float64sAreSorted(x []float64) bool: 报告切片x是否按升序排序（非数字NaN排最前面）
func TestFloat64AreSorted(t *testing.T) {
	s := []float64{0.7, 1.3, 2.6, 3.8, 5.2} // sorted ascending
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 3.8, 2.6, 1.3, 0.7} // sorted descending
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 1.3, 0.7, 3.8, 2.6} // unsorted
	fmt.Println(sort.Float64sAreSorted(s))
	// output:
	// true
	// false
	// false
}

// func Ints(x []int): int类型切片x按升序排序
func TestInts(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4 5 6]
}

// func IntsAreSorted(x []int) bool: 报告切片x是否按升级排序
func TestIntsAreSorted(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6} // sorted ascending
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{6, 5, 4, 3, 2, 1} // sorted descending
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{3, 2, 4, 1, 5} // unsorted
	fmt.Println(sort.IntsAreSorted(s))
	// output:
	// true
	// false
	// false
}

// func IsSorted(data Interface) bool: 报告data是否已排序
func TestIsSorted(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6} // sorted ascending
	fmt.Println(sort.IsSorted(sort.IntSlice(s)))

	s = []int{6, 5, 4, 3, 2, 1} // sorted descending
	fmt.Println(sort.IsSorted(sort.IntSlice(s)))

	s = []int{3, 2, 4, 1, 5} // unsorted
	fmt.Println(sort.IsSorted(sort.IntSlice(s)))
	// output:
	// true
	// false
	// false
}

// func Search(n int, f func(int) bool) int: 使用二分搜索查找并返回[0,n)中f(i)为true的最小索引i
func TestSearch(t *testing.T) {
	{
		a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
		x := 6

		i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
		if i < len(a) && a[i] == x {
			fmt.Printf("found %d at index %d in %v\n", x, i, a)
		} else {
			fmt.Printf("%d not found in %v\n", x, a)
		}
	}
	{
		a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
		x := 6

		i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
		if i < len(a) && a[i] == x {
			fmt.Printf("found %d at index %d in %v\n", x, i, a)
		} else {
			fmt.Printf("%d not found in %v\n", x, a)
		}
	}
	// output:
	// found 6 at index 2 in [1 3 6 10 15 21 28 36 45 55]
	// found 6 at index 7 in [55 45 36 28 21 15 10 6 3 1]
}

// func SearchFloat64s(a []float64, x float64) int：在已排序的 float64s 切片中搜索 x 并返回 Search 指定的索引。如果 x 不存在，则返回值是插入 x 后的索引（可能是 len(a)）。切片必须按升序排序。
func TestSearchFloat64s(t *testing.T) {
	s := []float64{0.7, 1.3, 2.6, 3.8, 5.2}
	fmt.Println(sort.SearchFloat64s(s, 2.6))
	fmt.Println(sort.SearchFloat64s(s, 3.6))
	// output:
	// 2
	// 3
}

// func SearchInts(a []int, x int) int: 和SearchFloat64一样，类型为int
func TestSearchInts(t *testing.T) {
	s := []int{1, 3, 5, 7, 9}
	fmt.Println(sort.SearchInts(s, 5))
	fmt.Println(sort.SearchInts(s, 6))
	// output:
	// 2
	// 3
}

// func SearchStrings(a []string, x string) int: 和SearchFloat64一样，类型为string
func TestSearchStrings(t *testing.T) {
	s := []string{"a", "c", "e", "g"}
	fmt.Println(sort.SearchStrings(s, "e"))
	fmt.Println(sort.SearchStrings(s, "f"))
	// output:
	// 2
	// 3
}

// func Slice(x interface{}, less func(i, j int) bool): 给定切片和less比较函数，对切片x进行排序
// 如果 x 不是切片，它会发生恐慌。 排序不能保证稳定：相等的元素可能会从它们的原始顺序颠倒过来。对于稳定排序，请使用 SliceStable。
func TestSlice(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
	// output:
	// By name: [{Alice 55} {Bob 75} {Gopher 7} {Vera 24}]
	// By age: [{Gopher 7} {Vera 24} {Alice 55} {Bob 75}]
}

// func SliceIsSorted(x interface{}, less func(i, j int) bool) bool: 报告切片 x 是否根据提供的 less 函数进行排序。如果 x 不是切片，它会发生恐慌。
func TestSliceIsSorted(t *testing.T) {
	type people struct {
		Name string
		Age  int
	}
	p1 := []people{
		{"Alice", 55},
		{"Bob", 75},
		{"Gopher", 7},
		{"Vera", 24},
	}
	p2 := []people{
		{"Bob", 75},
		{"Alice", 55},
		{"Vera", 24},
		{"Gopher", 7},
	}
	fmt.Println(sort.SliceIsSorted(p1, func(i, j int) bool { return p1[i].Name < p1[j].Name }))
	fmt.Println(sort.SliceIsSorted(p2, func(i, j int) bool { return p2[i].Name < p2[j].Name }))
	// output:
	// true
	// false
}

// func SliceStable(x interface{}, less func(i, j int) bool): 使用提供的 less 函数对切片 x 进行排序，以原始顺序保持相等的元素。如果 x 不是切片，它会发生恐慌。
func TestSliceStable(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// Sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	// Sort by age preserving name order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age,name:", people)
	// output:
	// By name: [{Alice 25} {Alice 75} {Alice 75} {Bob 75} {Bob 25} {Colin 25} {Elizabeth 75} {Elizabeth 25}]
	// By age,name: [{Alice 25} {Bob 25} {Colin 25} {Elizabeth 25} {Alice 75} {Alice 75} {Bob 75} {Elizabeth 75}]
}

// func Sort(data Interface): 对数据按升序进行排序。它对 data.Len 进行一次调用以确定对 data.Less 和 data.Swap 的 n 次调用和 O(n*log(n)) 次调用。不能保证排序是稳定的。
func TestSort(t *testing.T) {
	s := sort.IntSlice{2, 3, 1, 4, 5}
	sort.Sort(s)
	fmt.Println(s) // [1 2 3 4 5]
}

// func Stable(data Interface): 对数据按升序进行排序，同时保持相等元素的原始顺序。
func TestStable(t *testing.T) {
	s := sort.IntSlice{2, 3, 1, 4, 5}
	sort.Stable(s)
	fmt.Println(s) // [1 2 3 4 5]
}

// func Strings(x []string)： 对字符串切片x按升序进行排序
func TestStrings(t *testing.T) {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s) // [Alpha Bravo Delta Go Gopher Grin]
}

// func StringsAreSorted(x []string) bool: 报告x是否已按升序排序
func TestStringsAreSorted(t *testing.T) {
	s1 := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	s2 := []string{"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin"}
	fmt.Println(sort.StringsAreSorted(s1))
	fmt.Println(sort.StringsAreSorted(s2))
	// output:
	// false
	// true
}

// 接口:
// type Interface: 自定义排序切片类型需实现该接口
// type Interface interface {
//    // Len is the number of elements in the collection.
//    Len() int
//
//    // Less reports whether the element with index i
//    // must sort before the element with index j.
//    //
//    // If both Less(i, j) and Less(j, i) are false,
//    // then the elements at index i and j are considered equal.
//    // Sort may place equal elements in any order in the final result,
//    // while Stable preserves the original input order of equal elements.
//    //
//    // Less must describe a transitive ordering:
//    //  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//    //  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//    //
//    // Note that floating-point comparison (the < operator on float32 or float64 values)
//    // is not a transitive ordering when not-a-number (NaN) values are involved.
//    // See Float64Slice.Less for a correct implementation for floating-point values.
//    Less(i, j int) bool
//
//    // Swap swaps the elements with indexes i and j.
//    Swap(i, j int)
// }

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func TestInterface(t *testing.T) {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
	// output:
	// [Bob: 31 John: 42 Michael: 17 Jenny: 26]
	// [Michael: 17 Jenny: 26 Bob: 31 John: 42]
	// [John: 42 Bob: 31 Jenny: 26 Michael: 17]
}

// func Reverse(data Interface) Interface: 翻转数据
func TestReverse(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s) // [6 5 4 3 2 1]
}

// 类型：（实现sort.Interface接口的自定义基础类型）
// type Float64Slice
//    func (x Float64Slice) Len() int
//    func (x Float64Slice) Less(i, j int) bool
//    func (p Float64Slice) Search(x float64) int
//    func (x Float64Slice) Sort()
//    func (x Float64Slice) Swap(i, j int)
// type IntSlice
//    func (x IntSlice) Len() int
//    func (x IntSlice) Less(i, j int) bool
//    func (p IntSlice) Search(x int) int
//    func (x IntSlice) Sort()
//    func (x IntSlice) Swap(i, j int)
// type StringSlice
//    func (x StringSlice) Len() int
//    func (x StringSlice) Less(i, j int) bool
//    func (p StringSlice) Search(x string) int
//    func (x StringSlice) Sort()
//    func (x StringSlice) Swap(i, j int)
