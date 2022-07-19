package algo

import (
	"fmt"

	"github.com/deepakpunjabi/my-first-go/ds"
)

func PrintDfs(g ds.Graph) {
	fmt.Println("Printing DFS traversal")

	v := g.Vertices
	//edges := g.Edges

	visited := make([]bool, v)
	source := 0
	dfs(g, visited, source)
}

func dfs(g ds.Graph, visited []bool, node int) {
	visited[node] = true
	fmt.Println(node)

	adj := g.Edges[node]

	for i := range adj {
		vertex := g.Edges[node][i]
		if visited[vertex] == false {
			dfs(g, visited, vertex)
		}
	}
}
