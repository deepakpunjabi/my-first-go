package algo

import (
	"fmt"
	"../ds"
)

func PrintBfs(g ds.Graph) {
	fmt.Println("Printing...")

	v := g.Vertices
	edges := g.Edges
	//fmt.Println(edges)

	queue := ds.CreateQueue()
	//fmt.Println(queue.IsEmpty())

	visited := make([]bool, v)
	//fmt.Println("visited --> ",visited)

	source := 0
	visited[source] = true
	queue.Push(source)
	//fmt.Println(queue)

	for queue.IsEmpty() == false {
		front := queue.Front()
		queue.Pop()

		fmt.Println(front)
		//fmt.Println("edges.front --> ",edges[front])

		for i := range edges[front] {
			//fmt.Println("i",i)
			vertex := edges[front][i]
			if visited[vertex] == false {
				visited[vertex] = true
				queue.Push(vertex)
				//fmt.Println("visited -->",visited)
				//fmt.Println("queue -->",queue)
			}
		}
	}
}