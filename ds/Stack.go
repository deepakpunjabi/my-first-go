package ds

type Stack struct {
	stack []int
}

func (s *Stack) Top() int {
	return s.stack[0]
}

func (s *Stack) Push(data int) {
	s.stack = append([]int {data}, s.stack...)
}

func (s *Stack) Pop() {
	s.stack = s.stack[1:]
}