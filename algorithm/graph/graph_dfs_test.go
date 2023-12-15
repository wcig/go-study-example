// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package graph

import (
	"fmt"
	"testing"
)

func TestGraphDFS(t *testing.T) {
	/* 初始化无向图 */
	vets := valsToVets([]int{0, 1, 2, 3, 4, 5, 6})
	edges := [][]Vertex{
		{vets[0], vets[1]}, {vets[0], vets[3]}, {vets[1], vets[2]},
		{vets[2], vets[5]}, {vets[4], vets[5]}, {vets[5], vets[6]}}
	graph := newGraphAdjList(edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 深度优先遍历 DFS */
	res := graphDFS(graph, vets[0])
	fmt.Println("深度优先遍历（DFS）顶点序列为:")
	fmt.Println(vetsToVals(res))

	// Output:
	// 初始化后，图为:
	//	领接表 =
	//		0: 1 3
	//		1: 0 2
	//		2: 1 5
	//		3: 0
	//		4: 5
	//		5: 2 4 6
	//		6: 5
	// 深度优先遍历（DFS）顶点序列为:
	// [0 1 2 5 4 6 3]
}
