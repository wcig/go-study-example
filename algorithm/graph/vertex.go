// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package graph

/* 顶点类 */
type Vertex struct {
	Val int
}

/* 构造函数 */
func NewVertex(val int) Vertex {
	return Vertex{
		Val: val,
	}
}

func valsToVets(vals []int) []Vertex {
	vets := make([]Vertex, len(vals))
	for i := range vals {
		vets[i] = NewVertex(vals[i])
	}
	return vets
}

func vetsToVals(vets []Vertex) []int {
	vals := make([]int, len(vets))
	for i := range vets {
		vals[i] = vets[i].Val
	}
	return vals
}

func removeSliceItem(vets []Vertex, vet Vertex) []Vertex {
	if len(vets) == 0 {
		return vets
	}
	index := -1
	for i := range vets {
		if vets[i] == vet {
			index = i
		}
	}
	if index == -1 {
		return vets
	}
	return append(vets[:index], vets[index+1:]...)
}

func isSliceItemExist(vets []Vertex, vet Vertex) bool {
	for i := range vets {
		if vets[i] == vet {
			return true
		}
	}
	return false
}
