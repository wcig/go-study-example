// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package graph

/* 深度优先遍历 (以领接表表示图) */
func graphDFS(g *graphAdjList, startVet Vertex) []Vertex {
	// 结果顶点列表
	var res []Vertex
	// 哈希表记录顶点是否被访问过
	visited := map[Vertex]struct{}{}
	dfs(g, startVet, &res, visited)
	return res
}

func dfs(g *graphAdjList, vet Vertex, res *[]Vertex, visited map[Vertex]struct{}) {
	// 记录访问顶点
	*res = append(*res, vet)
	visited[vet] = struct{}{}
	// 遍历顶点所有的领接顶点
	for _, adjVet := range g.adjList[vet] {
		if _, ok := visited[adjVet]; !ok {
			// 递归访问领接顶点
			dfs(g, adjVet, res, visited)
		}
	}
}
