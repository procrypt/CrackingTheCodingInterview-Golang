/*
Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but
the first and last node, not necessarily the exact middle) of a singly linked list, given only access to
that node.

EXAMPLE
Input:  the node c from the linked list a->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a->b->d->e->f

Solution:
Since we are only given the node to be deleted, what we can do here is replace the next node with the node
to be deleted.
eg.
Delete c
a->b->c->d->e->f
Replace c with d
	c.data = c.next.data
	c.next = c.next.next
a->b->d->e->f
*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
)

type LinkedList struct {
	list linkedList.LinkedList
}

func (l *LinkedList) deleteMiddleNode(node *linkedList.Node) {
	if node != nil && node.Next != nil {
		node.Key = node.Next.Key
		node.Next = node.Next.Next
	} else {
		return
	}
}

func main() {
	l := LinkedList{}
	values := []int{2, 1, 3, 6, 3, 4, 5}
	for _, v := range values {
		l.list.Append(v)

	}
	l.deleteMiddleNode(l.list.Head.Next.Next.Next)
	l.list.Traversal()
}
