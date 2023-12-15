// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph
package graph

import (
	"fmt"
	"testing"
)

func TestGraphBFS(t *testing.T) {
	/* 初始化无向图 */
	vets := valsToVets([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	edges := [][]Vertex{
		{vets[0], vets[1]}, {vets[0], vets[3]}, {vets[1], vets[2]}, {vets[1], vets[4]},
		{vets[2], vets[5]}, {vets[3], vets[4]}, {vets[3], vets[6]}, {vets[4], vets[5]},
		{vets[4], vets[7]}, {vets[5], vets[8]}, {vets[6], vets[7]}, {vets[7], vets[8]}}
	graph := newGraphAdjList(edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 广度优先遍历 BFS */
	res := graphBFS(graph, vets[0])
	fmt.Println("广度优先遍历（BFS）顶点序列为:")
	fmt.Println(vetsToVals(res))

	// Output:
	// 初始化后，图为:
	//	领接表 =
	//		0: 1 3
	//		1: 0 2 4
	//		2: 1 5
	//		3: 0 4 6
	//		4: 1 3 5 7
	//		5: 2 4 8
	//		6: 3 7
	//		7: 4 6 8
	//		8: 5 7
	// 广度优先遍历（BFS）顶点序列为:
	// [0 1 3 2 4 6 5 7 8]
}
