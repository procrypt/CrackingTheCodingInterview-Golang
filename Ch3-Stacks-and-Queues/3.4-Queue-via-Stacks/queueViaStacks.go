/*
Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks.

The idea is to reverse the order of the elements of the stack by pushing them into a the stack.
*/

package main

import (
	"algorithms-in-golang/CTCI/Ch3-Stacks-and-Queues/Stack"
	"fmt"
)

type MyQueue struct {
	stack1 Stack.Stack
	stack2 Stack.Stack
}

func(m *MyQueue) front() interface{} {
	if m.stack2.IsEmpty() {
		for !m.stack1.IsEmpty() {
			m.stack2.Push(m.stack1.Pop())
		}
	}
	return m.stack2.Pop()
}

func (m *MyQueue) size() int {
	return len(m.stack1.Container) + len(m.stack2.Container)
}

func (m *MyQueue) enque(value interface{}) {
	m.stack1.Push(value)
}

func (m *MyQueue) dequeue() interface{} {
	if m.stack2.IsEmpty() {
		for !m.stack1.IsEmpty() {
			m.stack2.Push(m.stack1.Pop())
		}
	}
	return m.stack2.Pop()
}

func main() {
	m := MyQueue{}
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range value {
		m.enque(v)
	}
	for _, v := range m.stack1.Container {
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(m.dequeue())
	for _, v := range m.stack2.Container {
		fmt.Print(v)
	}
	fmt.Println(m.front())
}
