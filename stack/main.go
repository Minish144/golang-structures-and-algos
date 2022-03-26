package main

import "fmt"

type Stack struct {
	values []int
	Size   int
}

func (s *Stack) Append(val int) {
	s.values = append(s.values, val)
	s.Size++
}

func (s *Stack) Pop() int {
	last := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	s.Size--
	return last
}

func (s *Stack) Traverse() {
	for index, val := range s.values {
		fmt.Printf("[%d]: %d\n", index, val)
	}
}

func main() {
	st := Stack{make([]int, 0), 0}
	st.Append(5)
	st.Append(10)
	st.Append(15)

	fmt.Println(st.Pop())

	st.Traverse()
}
