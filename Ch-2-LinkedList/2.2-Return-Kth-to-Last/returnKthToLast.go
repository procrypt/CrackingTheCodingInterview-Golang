/*
Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.

Method 1:
1) Get the length of the LinkedList
2) Count "down"(count in reverse order) the total length until "kth to last element" is reached

Method 2:
Have two pointers with one at distance of k from other one.
eg, if k = 2, then
a points to 1 and b points to 3
Now whenever b points to NULL, a will be at Kth to the last element of the list.
*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
	"fmt"
)
// Method 1
func kthToLastElement1(l *linkedList.LinkedList, k int) int {
	current := l.Head
	lenght := l.Length()
	for current != nil {
		if lenght == k {
			return current.Key
		} else {
			lenght--
		}
		current = current.Next
	}
	return 0
}

// Method 2
func kthToLastElement2(l *linkedList.LinkedList, k int) int {
	current := l.Head
	a := current
	b := current


	// This is how we increment b by k nodes from a
	count := 0
	for b != nil && count < k {
		b = b.Next
		count++
	}
	for a != nil {
		if b == nil {
			return a.Key
		} else {
			b = b.Next
		}
		a = a.Next
	}
	return 0
}


func main() {
	l := linkedList.LinkedList{}
	values := []int{2, 1, 3, 6, 3, 4, 5}
	for _, v := range values {
		l.Append(v)

	}
	fmt.Println(kthToLastElement1(&l, 2))
	fmt.Println(kthToLastElement2(&l, 4))
}
