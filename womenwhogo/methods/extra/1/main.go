package main

import "fmt"

type Stack struct {
	positions []int
}

func main() {
	stackOne := Stack{
		positions: []int{1, 2, 3, 4, 5},
	}
	stackOne.Top()
	stackOne.Push(10)
	stackOne.Top()
	fmt.Println(stackOne.positions)
	fmt.Println("POP:", stackOne.Pop())
	fmt.Println(stackOne.positions)

}

func (s Stack) Top() int {
	fmt.Println("Top: ", s.positions[len(s.positions)-1])
	return s.positions[0]
}

func (s *Stack) Push(i int) {

	s.positions = append(s.positions, i)
	fmt.Println(s.positions)
	//no need to return s.positions
}

func (s *Stack) Pop() int {
	last := s.positions[len(s.positions)-1]
	s.positions = s.positions[:len(s.positions)-1]
	return last
}
