// refer: https://github.com/krahets/hello-algo/tree/main/codes/go/chapter_graph

package graph

/* 广度优先遍历 (以领接表表示图) */
func graphBFS(g *graphAdjList, startVet Vertex) []Vertex {
	// 结果顶点列表
	var res []Vertex
	// 辅助队列
	queue := []Vertex{startVet}
	// 哈希表记录顶点是否被访问过, 入队即被访问过
	visited := map[Vertex]struct{}{
		startVet: {},
	}
	for len(queue) != 0 {
		// 队头顶点出队
		vet := queue[0]
		queue = queue[1:]
		// 记录访问顶点
		res = append(res, vet)
		// 遍历顶点所有的领接顶点
		for _, adjVet := range g.adjList[vet] {
			if _, ok := visited[adjVet]; !ok {
				// 未访问顶点入队
				queue = append(queue, adjVet)
				visited[adjVet] = struct{}{}
			}
		}
	}
	return res
}
