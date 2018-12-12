/*
Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
before all nodes greater than or equal to x. If x is contained within the list, the values of x only need
to be after the elements less than x (see below). The partition element x can appear anywhere in the
"right partition"; it does not need to appear between the left and right partitions.
EXAMPLE
Input:  3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition = 5]
Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8

we create two different linked lists: one for elements less than x, and one for elements greater than or equal to x.
We iterate through the linked list, inserting elements into our before list or our after list.
Once we reach the end of the linked list and have completed this splitting, we merge the two lists.

*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
)

func partition(l *linkedList.LinkedList, partition int) linkedList.LinkedList {
	current := l.Head
	less := linkedList.LinkedList{}
	more := linkedList.LinkedList{}
	for current != nil {
		if current.Key < partition {
			less.Append(current.Key)
		} else if current.Key >= partition {
			more.Append(current.Key)
		}
		current = current.Next
	}
	less.Tail.Next = more.Head
	less.Traversal()
	return less
}

func main() {
	l := linkedList.LinkedList{}
	values := []int{1, 4, 3, 2, 5, 2}
	for _, v := range values {
		l.Append(v)
	}
	partition(&l, 3)
}
