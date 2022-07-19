package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/deepakpunjabi/my-first-go/algo"
	"github.com/deepakpunjabi/my-first-go/ds"
)

func toInt(data string) int {
	convData, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	} else {
		return convData
	}
}

func main() {

	fmt.Println("Running...")

	id := 1
	// get parameters
	args := os.Args

	//fmt.Println(args[1])
	numVertices := toInt(args[id])
	fmt.Println("numVertices", numVertices)
	id++
	numEdges := toInt(args[id])
	id++

	// initialize
	g := ds.CreateGraph(numVertices)
	for i := 0; i < numEdges; i++ {
		u := toInt(args[id])
		id++
		v := toInt(args[id])
		id++
		g.AddEdge(u, v)
	}

	g.Print()

	operation := toInt(args[id])
	switch operation {
	case 1:
		algo.PrintBfs(g)
	case 2:
		algo.PrintDfs(g)
	default:
		fmt.Println("incorrect operation")
	}
	//q := ds.CreateQueue()
	//q.Push(11)
	//q.Printer()
	//q.Push(12)
	//q.Printer()
	//q.Pop()
	//q.Printer()
	//fmt.Println(q)

	//fmt.Println("------")
	//g := ds.CreateGraph(6)
	//g.Print()
	//g.AddEdge(0,1)
	//g.AddEdge(0,2)
	//g.AddEdge(1,3)
	//g.AddEdge(4,3)
	//g.AddEdge(5,3)
	//g.AddEdge(2,5)

	//algo.PrintBfs(g)
	//algo.PrintDfs(g)
}
