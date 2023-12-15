// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package graph

import (
	"fmt"
	"sort"
	"strings"
)

/* 基于邻接表实现的无向图类 */
type graphAdjList struct {
	// 邻接表，key：顶点，value：该顶点的所有邻接顶点
	adjList map[Vertex][]Vertex
}

/* 构造函数 */
func newGraphAdjList(edges [][]Vertex) *graphAdjList {
	g := &graphAdjList{
		adjList: make(map[Vertex][]Vertex),
	}
	for i := range edges {
		vet1 := edges[i][0]
		vet2 := edges[i][1]
		g.addVertex(vet1)
		g.addVertex(vet2)
		g.addEdge(vet1, vet2)
	}
	return g
}

/* 添加边 (未处理边已存在情况) */
func (g *graphAdjList) addEdge(vet1 Vertex, vet2 Vertex) {
	g.validEdge(vet1, vet2)
	if !g.existEdge(vet1, vet2) {
		g.adjList[vet1] = append(g.adjList[vet1], vet2)
		g.adjList[vet2] = append(g.adjList[vet2], vet1)
	}
}

/* 删除边 */
func (g *graphAdjList) removeEdge(vet1 Vertex, vet2 Vertex) {
	g.validEdge(vet1, vet2)
	g.adjList[vet1] = removeSliceItem(g.adjList[vet1], vet2)
	g.adjList[vet2] = removeSliceItem(g.adjList[vet2], vet1)
}

/* 校验边 */
func (g *graphAdjList) validEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 {
		panic(fmt.Sprintf("vertex [%d] not exists", vet1.Val))
	}
	if !ok2 {
		panic(fmt.Sprintf("vertex [%d] not exists", vet2.Val))
	}
	if vet1 == vet2 {
		panic(fmt.Sprintf("invalid edge: [%d] - [%d]", vet1.Val, vet2.Val))
	}
}

/* 是否存在边 */
func (g *graphAdjList) existEdge(vet1 Vertex, vet2 Vertex) bool {
	return isSliceItemExist(g.adjList[vet1], vet2) &&
		isSliceItemExist(g.adjList[vet2], vet1)
}

/* 添加顶点 */
func (g *graphAdjList) addVertex(vet Vertex) {
	if _, ok := g.adjList[vet]; !ok {
		g.adjList[vet] = []Vertex{}
	}
}

/* 删除顶点 */
func (g *graphAdjList) removeVertex(vet Vertex) {
	if _, ok := g.adjList[vet]; !ok {
		panic(fmt.Sprintf("vertex [%d] not exist", vet.Val))
	}
	delete(g.adjList, vet)
	for k := range g.adjList {
		g.adjList[k] = removeSliceItem(g.adjList[k], vet)
	}
}

/* 获取顶点数量 */
func (g *graphAdjList) size() int {
	return len(g.adjList)
}

/* 打印 */
func (g *graphAdjList) print() {
	vets := make([]Vertex, 0, g.size())
	for k := range g.adjList {
		vets = append(vets, k)
	}
	sort.Slice(vets, func(i, j int) bool {
		return vets[i].Val < vets[j].Val
	})

	fmt.Printf("\t领接表 = \n")
	var sb strings.Builder
	for i := range vets {
		vet := vets[i]
		adjVets := g.adjList[vet]
		sb.WriteString(fmt.Sprintf("\t\t%d: ", vet.Val))
		for j := range adjVets {
			sb.WriteString(fmt.Sprintf("%d ", adjVets[j].Val))
		}
		fmt.Println(sb.String())
		sb.Reset()
	}
}
