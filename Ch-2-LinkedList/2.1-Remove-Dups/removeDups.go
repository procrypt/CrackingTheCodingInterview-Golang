/*
Write code to remove duplicates from an unsorted linked list.

Make sure you keep the track of previous node while working with unsorted linked list.
*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
)


// Make sure you keep the track of previous node
func removeDupsUnsorted(list *linkedList.LinkedList) *linkedList.LinkedList {
	m := make(map[int]int)
	current := list.Head
	prev := &linkedList.Node{} // initially set it to nil
	for current != nil {
			if _, ok := m[current.Key]; !ok {
				m[current.Key] = 1
				prev = current
			} else {
				prev.Next = current.Next
				list.Size--
			}
		current = current.Next
	}
	return list
}

func removeDupsSorted(list *linkedList.LinkedList) *linkedList.LinkedList {
	current := list.Head
	for current != nil {
		if current.Next != nil {
			if current.Key == current.Next.Key {
				current.Next = current.Next.Next
				list.Size--
			}
		}
		current = current.Next
	}
	return list
}

func main() {
	l := linkedList.LinkedList{}
	values := []int{2,1,3,6,3,4,5}
	for _, v := range values {
		l.Append(v)

	}
	removeDupsUnsorted(&l)
	l.Traversal()
}