package ds

import (
	"fmt"
)

// single element in queue
type Node struct {
	value int
}

func createNode(val int) Node {
	node := Node{val}
	return node
}

// list of elements
type Queue struct {
	nodeList []Node
}

// stack constructor for empty stack
func CreateQueue() Queue {
	newQueue := Queue{}
	return newQueue
}

// add an element at the end of the queue
func (q *Queue) Push(data int) {
	node := createNode(data)
	q.nodeList = append(q.nodeList, node)
	//fmt.Println(q.nodeList)
}

// remove first element from the queue
func (q *Queue) Pop() {
	q.nodeList = q.nodeList[1:]
}

func (q *Queue) IsEmpty() bool {
	if len(q.nodeList) == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) Front() int {
	return q.nodeList[0].value
}

func (q Queue) String() string {
	return fmt.Sprintf("%v", q.nodeList)
}

func (q Queue) Printer() {
	fmt.Println(q.nodeList)
}

func main1() {
	fmt.Println(CreateQueue())
}
