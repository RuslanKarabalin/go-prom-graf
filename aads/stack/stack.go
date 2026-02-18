package main

import "fmt"

type IntStack struct {
	buf  []int
	tail int
	size int
}

func NewIntStack(cap int) IntStack {
	if cap < 1 {
		cap = 1
	}
	return IntStack{buf: make([]int, cap)}
}

func (s *IntStack) Push(v int) {
	if s.size == len(s.buf) {
		s.buf = append(s.buf, v)
	} else {
		s.buf[s.tail] = v
	}
	s.tail++
	s.size++
}

func (s *IntStack) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	s.tail--
	s.size--
	v := s.buf[s.tail]
	return v, true
}

func (s *IntStack) Print() {
	for i := 0; i < s.size; i++ {
		fmt.Println(s.buf[i])
	}
}

func main() {
	stack := NewIntStack(2)
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Print()
	fmt.Println("-")
	stack.Pop()
	stack.Print()
}
