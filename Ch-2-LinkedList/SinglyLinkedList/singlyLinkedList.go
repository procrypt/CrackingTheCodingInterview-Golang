/*
Singly LinkedList Implementation
*/

package SinglyLinkedList

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Key  int
	Next *Node
}

func(l *LinkedList) Length() int {
	return l.Size
}

func(l *LinkedList) IsEmpty() bool {
	if l.Length() != 0 {
		return false
	}
	return true
}

func(l *LinkedList) Append(value int) {
	node := &Node{Key:value}
	if l.Size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size++
}

func(l *LinkedList) Traversal() {
	for l.Head != nil {
			fmt.Println(l.Head.Key)
			l.Head = l.Head.Next
	}
}
