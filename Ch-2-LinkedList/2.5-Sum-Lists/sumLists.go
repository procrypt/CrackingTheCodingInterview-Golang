/*
Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.
The digits are stored in reverse order, such that the 1 's digit is at the head of the list.
Write a function that adds the two numbers and returns the sum as a linked list.
EXAMPLE
Input: (7-> 1 -> 6) + (5 -> 9 -> 2).That is,617 + 295.
Output: 2 -> 1 -> 9. That is, 912.
FOLLOW UP
Suppose the digits are stored in forward order. Repeat the above problem.
EXAMPLE
lnput:(6 -> 1 -> 7) + (2 -> 9 -> 5).That is,617 + 295.
Output: 9 -> 1 -> 2. That is, 912.

Define two pointers p and q for list l1 and l2.
Check is both they are not nil, if one of the list is shorter than the other one make sure to add two to it,
to make both of them equal in length.
If there is carry add it to the next sum just like classical addition.
If the carry is > 0
carry = sum/10
remainder = sum % 10
*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
)

func sumList(l1, l2 *linkedList.LinkedList) linkedList.LinkedList {
	sumList := linkedList.LinkedList{}
	p := l1.Head
	q := l2.Head
	carry, i, j := 0,0,0

	for p != nil || q != nil {
		// to make sure both the list are equal in length
		if p == nil {
			i = 0
		} else {
			i = p.Key
			p = p.Next
		}
		if q == nil {
			j = 0
		} else {
			j = q.Key
			q = q.Next
		}
		sum := i + j + carry
		if sum >= 10 {
			// will give the carry tens position of the number, 30/10 = 3
			carry = sum/10
			// will give the ones position of the number, 12 % 10 = 2
			remainder := sum % 10
			sumList.Append(remainder)
		} else {
			carry = 0
			sumList.Append(sum)
		}
	}

	return sumList
}

func main() {
	l1 := linkedList.LinkedList{}
	l1.Append(5)
	l1.Append(6)
	l1.Append(3)
	l2 := linkedList.LinkedList{}
	l2.Append(8)
	l2.Append(4)
	l2.Append(2)
	result := sumList(&l1, &l2)
	result.Traversal()
}

