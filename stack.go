package main

import (
	"fmt"
)

type Stack struct {
	Arrays []int
}

func (s *Stack) push(i int) {
	s.Arrays = append(s.Arrays, i)
}

func (s *Stack) pop() int {
	if v := s.size(); v > 0 {
		c := s.Arrays[v-1]
		s.Arrays = s.Arrays[:v-1]
		return c
	}
	return 0
}

func (s *Stack) peek() int {
	if !s.isEmpty() {
		return s.Arrays[len(s.Arrays)-1]
	}
	return 0
}

func (s *Stack) isEmpty() bool {
	return len(s.Arrays) == 0
}

func (s *Stack) size() int {
	return len(s.Arrays)
}

func main() {
	s := Stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())

}
