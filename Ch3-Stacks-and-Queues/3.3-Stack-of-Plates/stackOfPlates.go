package main

import (
	"algorithms-in-golang/CTCI/Ch3-Stacks-and-Queues/Stack"
)

type StackList struct {
	value *Stack.Stack
	next  *StackList
}

type StackSet struct {
	maxCapacity int
	itemCount   int
	stackCount  int
	stacks       *StackList
}

func NewStackSet(max int) *StackSet {
	stack := &Stack.Stack{}
	s := &StackSet{
		stacks: &StackList{
			value:stack,
		},
		stackCount: 1,
		maxCapacity: max,
	}
	return s
}

func(s *StackSet) Push(value interface{}) {
	// Total Capacity is maxCapacity * StackCount
	if (s.maxCapacity * s.stackCount) == s.itemCount {
		stack := &Stack.Stack{}
		stack.Push(value)

		list := &StackList{
			value: stack,
			next: s.stacks,
		}
		s.stacks = list
		s.stackCount++
	} else {
			s.stacks.value.Push(value)
	}
	s.itemCount++
}

func(s *StackSet) Pop() interface{} {
	if s.itemCount == 0 {
		return nil
	}
	pop := s.stacks.value.Pop()
	s.itemCount--

	if s.itemCount == ((s.stackCount-1)*s.maxCapacity) {
		s.stacks = s.stacks.next
		s.stackCount--
	}
	return pop
}