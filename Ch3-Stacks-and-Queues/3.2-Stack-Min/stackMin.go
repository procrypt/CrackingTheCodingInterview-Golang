package main

import "fmt"

type Stack struct {
	Key       interface{}
	Container []interface{}
	Minimum   []interface{}
}

func (s *Stack) Pop() {
	pop := s.Container[len(s.Container)-1]
	min := s.Minimum[len(s.Minimum)-1]
	if pop.(int) == min.(int) {
		s.PopMin()
	}
	s.Container = append(s.Container[:len(s.Container)-1])
}

func (s *Stack) Push(value interface{}) {
	s.Container = append(s.Container, value)
	// Add element in the min stack
	if len(s.Minimum) != 0 {
		min := s.PeekMin()
		if value.(int) > min.(int) {
			return
		} else {
			s.PushMin(value)
		}
	} else {
		s.PushMin(value)
	}
}

func (s *Stack) Peek() interface{} {
	return s.Container[len(s.Container)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.Container) != 0 {
		return false
	}
	return true
}

func (s *Stack) PeekMin() interface{} {
	return s.Minimum[len(s.Minimum)-1]
}

func (s *Stack) PushMin(value interface{}) {
	s.Minimum = append(s.Minimum, value)
}

func (s *Stack) PopMin() {
	s.Minimum = append(s.Minimum[:len(s.Minimum)-1])
}

func (s *Stack) Min() interface{} {
	min := s.Minimum[len(s.Minimum)-1]
	return min
}

func main() {
	s := Stack{}
	values := []int{7, 8, 5, 9, 5, 2}
	for _, v := range values {
		s.Push(v)
	}
	for _, v := range s.Container {
		fmt.Println(v)
	}
	fmt.Println("--- Minimum ---")
	for _, v := range s.Minimum {
		fmt.Println(v)
	}
	fmt.Println("--")
	fmt.Println(s.Min())
	s.Pop()
	fmt.Println("--")
	fmt.Println(s.Min())
	s.Pop()
	fmt.Println("--")
	fmt.Println(s.Min())
	fmt.Println("+++")
	for _, v := range s.Container {
		fmt.Print(v)
	}
	s.Pop()
	fmt.Println("--")
	fmt.Println(s.Min())
	fmt.Println("+++")
	for _, v := range s.Container {
		fmt.Print(v)
	}
	fmt.Println("===")
	for _, v := range s.Minimum {
		fmt.Print(v)
	}
}
