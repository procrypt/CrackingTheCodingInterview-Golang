/*
Sort Stack: Write a program to sort a stack such that the smallest items are on the top.
You can use an additional temporary stack, but you may not copy the elements into any other data structure
(such as an array).

The stack supports the following operations: push, pop, peek, and isEmpty.
*/

package main

import (
	"algorithms-in-golang/CTCI/Ch3-Stacks-and-Queues/Stack"
	"fmt"
)

var newStack Stack.Stack

type sortStack struct {
	stack *Stack.Stack
}

func Sort(stack *Stack.Stack) *Stack.Stack {
	for !stack.IsEmpty() {
		temp := stack.Pop()
		for !newStack.IsEmpty() && newStack.Peek().(int) > temp.(int) {
			stack.Push(newStack.Pop())
		}
		newStack.Push(temp)
	}
	stack = &newStack
	return stack
}

func main() {
	s := &Stack.Stack{}
	values := []int{1, 2, 9, 6, 5}
	for _, v := range values {
		s.Push(v)
	}
	fmt.Println()
	s = Sort(s)
	for _, v := range s.Container {
		fmt.Print(v)
	}
}
