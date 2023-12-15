// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package graph

import "fmt"

/* 基于邻接矩阵实现的无向图类 */
type graphAdjMat struct {
	// 顶点列表，元素代表“顶点值”，索引代表“顶点索引”
	vertices []int
	// 邻接矩阵，行列索引对应“顶点索引”
	adjMat [][]int
}

/* 构造函数 */
func newGraphAdjMat(vertices []int, edges [][]int) *graphAdjMat {
	// 初始化
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	g := &graphAdjMat{
		vertices: vertices,
		adjMat:   adjMat,
	}

	// 添加边
	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])
	}
	return g
}

/* 添加边: i,j对应vertices索引 */
func (g *graphAdjMat) addEdge(i int, j int) {
	if (i < 0 || i >= len(g.vertices)) || (j < 0 || j >= len(g.vertices)) {
		panic(fmt.Sprintf("index out of range [%d, %d]", i, j))
	}
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

/* 删除边: i,j对应vertices索引 */
func (g *graphAdjMat) removeEdge(i int, j int) {
	if (i < 0 || i >= len(g.vertices)) || (j < 0 || j >= len(g.vertices)) {
		panic(fmt.Sprintf("index out of range [%d, %d]", i, j))
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

/* 添加顶点: val为顶点值 */
func (g *graphAdjMat) addVertex(val int) {
	n := len(g.vertices)
	g.vertices = append(g.vertices, val)
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

/* 删除顶点: index为vertices索引 */
func (g *graphAdjMat) removeVertex(index int) {
	if index < 0 || index >= len(g.vertices) {
		panic(fmt.Sprintf("index out of range [%d]", index))
	}
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

/* 获取顶点数量 */
func (g *graphAdjMat) size() int {
	return len(g.vertices)
}

/* 打印 */
func (g *graphAdjMat) print() {
	fmt.Printf("\t顶点列表 = %v\n", g.vertices)
	fmt.Printf("\t领接矩阵 = \n")
	for i := range g.adjMat {
		fmt.Printf("\t\t\t%v\n", g.adjMat[i])
	}
}
