package ds

import "fmt"

type Graph struct {
	Vertices int
	Edges    [][]int
}

func CreateGraph(v int) Graph {
	return Graph{Vertices: v, Edges: make([][]int, v)}
}

func (g *Graph) CreateGraph(v int) {
	g.Vertices = v
	g.Edges = make([][]int, v)
	for i := range g.Edges {
		g.Edges[i] = make([]int, v)
		fmt.Print(i, "--")
	}
}

func (g *Graph) AddEdge(u int, v int) {
	g.Edges[u] = append(g.Edges[u], v)
	g.Edges[v] = append(g.Edges[v], u)
}

func (g *Graph) Print() {
	fmt.Println("Graph")
	for i := range g.Edges {
		fmt.Println(i, "-->", g.Edges[i])
	}
	//fmt.Println(g.Edges)
}

func (g *Graph) PrintBfs() {

}
