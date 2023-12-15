// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package graph

import (
	"fmt"
	"testing"
)

func TestGraphAdjList(t *testing.T) {
	/* 初始化无向图 */
	v := valsToVets([]int{1, 3, 2, 5, 4})
	edges := [][]Vertex{{v[0], v[1]}, {v[0], v[3]}, {v[1], v[2]}, {v[2], v[3]}, {v[2], v[4]}, {v[3], v[4]}}
	graph := newGraphAdjList(edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 添加边 */
	// 顶点 1, 2 即 v[0], v[2]
	graph.addEdge(v[0], v[2])
	fmt.Println("\n添加边 1-2 后，图为")
	graph.print()

	/* 删除边 */
	// 顶点 1, 3 即 v[0], v[1]
	graph.removeEdge(v[0], v[1])
	fmt.Println("\n删除边 1-3 后，图为")
	graph.print()

	/* 添加顶点 */
	v5 := NewVertex(6)
	graph.addVertex(v5)
	fmt.Println("\n添加顶点 6 后，图为")
	graph.print()

	/* 删除顶点 */
	// 顶点 3 即 v[1]
	graph.removeVertex(v[1])
	fmt.Println("\n删除顶点 3 后，图为")
	graph.print()

	// Output:
	// 初始化后，图为:
	//	领接表 =
	//		1: 3 5
	//		2: 3 5 4
	//		3: 1 2
	//		4: 2 5
	//		5: 1 2 4
	//
	// 添加边 1-2 后，图为
	//	领接表 =
	//		1: 3 5 2
	//		2: 3 5 4 1
	//		3: 1 2
	//		4: 2 5
	//		5: 1 2 4
	//
	// 删除边 1-3 后，图为
	//	领接表 =
	//		1: 5 2
	//		2: 3 5 4 1
	//		3: 2
	//		4: 2 5
	//		5: 1 2 4
	//
	// 添加顶点 6 后，图为
	//	领接表 =
	//		1: 5 2
	//		2: 3 5 4 1
	//		3: 2
	//		4: 2 5
	//		5: 1 2 4
	//		6:
	//
	// 删除顶点 3 后，图为
	//	领接表 =
	//		1: 5 2
	//		2: 5 4 1
	//		4: 2 5
	//		5: 1 2 4
	//		6:
}
