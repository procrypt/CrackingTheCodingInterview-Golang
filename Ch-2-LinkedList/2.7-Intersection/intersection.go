/*
Intersection: Given two (singly) linked lists, determine if the two lists intersect.
Return the intersecting node. Note that the intersection is defined based on reference, not value.
That is, if the kth node of the first linked list is the exact same node (by reference) as the
jth node of the second linked list, then they are intersecting.

The two intersecting linked lists will always have the same last node.
Therefore, we can just traverse to the end of each linked list and compare the last nodes.

If the linked lists were the same length, you could just traverse through them at the same time.
When they collide, that's your intersection.

When they're not the same length, we'd like to just"chop off ignore-those excess nodes.
How can we do this?
If we know the lengths of the two linked lists, then the difference between those
two linked lists will tell us how much to chop off.

1. Run through each linked list to get the lengths and the tails.
2. Compare the tails. If they are different (by reference, not by value), return immediately.
   There is no inter section.
3. Set two pointers to the start of each linked list.
4. On the longer linked list, advance its pointer by the difference in lengths.
5. Now, traverse on each linked list until the pointers are the same.

*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
	"fmt"
	"math"
)

func Tail(l *linkedList.LinkedList) *linkedList.Node {
	current := l.Head
	tail := &linkedList.Node{}
	for current != nil {
		if current.Next == nil {
			tail = current
		}
		current = current.Next
	}
	return tail
}

func intersection(l1 *linkedList.LinkedList, l2 *linkedList.LinkedList) *linkedList.Node {
	list1, list2 := l1.Head, l2.Head
	lengthL1, tail1 := lengthAndTail(l1)
	lengthL2, tail2 := lengthAndTail(l2)
	if &tail1 != &tail2 {
		return nil
	}
	difference := abs(lengthL1,lengthL2)
	if lengthL1 > lengthL2 {
		counter := 0
		for list1 != nil && counter < difference {
			list1 = list1.Next
			counter++
		}
		for list1 != nil && list2 != nil {
			if &list1 == &list2 {
				return list1
			}
			list1 = list1.Next
			list2 = list2.Next
		}
	} else if lengthL2 > lengthL1 {
		counter := 0
		for list2 != nil && counter < difference {
			list2 = list2.Next
			counter++
		}
		for list1 != nil && list2 != nil {
			if &list1 == &list2 {
				return list1
			}
			list1 = list1.Next
			list2 = list2.Next
		}
	}
	return nil
}


func abs(a,b int) int {
	out := math.Abs(float64(a)-float64(b))
	return int(out)
}

func lengthAndTail( l *linkedList.LinkedList) (length int, tail *linkedList.Node) {
	current := l.Head
	prev := &linkedList.Node{}
	for current != nil {
		prev = current
		current = prev.Next
		if current == nil {
			tail = prev
		}
		length++
	}
	return length,tail
}

func main() {
	l1 :=  linkedList.LinkedList{}
	l2 := linkedList.LinkedList{}
	values1 := []int{1,2,3,4,5,6,7,8,9}
	values2 := []int{10,20,30,40,5,6,7,8,9}

	for _, v := range values1 {
		l1.Append(v)
	}
	for _, v := range values2 {
		l2.Append(v)
	}
	fmt.Print(intersection(&l1, &l2))
}