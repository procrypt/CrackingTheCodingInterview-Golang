/*
Loop Detection: Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.

DEFINITION
Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so
as to make a loop in the linked list.

Solution:
Create two runner traversing the linked list, with one runner running two nodes ahead, i.e.,
If slow is node1 then fast should be on node3, if they collide then there is a loop and we return that node.
If they are only one node head of each other they will never be able to collide.

Fast moves two nodes each time slow moves one node.


eg:
1->2->3->4->5->6->7
            \____/

Let us assume 7 is pointing to 5 the previous node, not to NIL. The runner will traverse the linked list as follows

R1  R2
1	2
2	5
3	7
4	5
5	7
6	5
7	7

When R1 will point to node6 R2 will also be pointing to node6 because of the loop.

If R1 and R2 are just one node ahead they will never collide and we will never be able to detect the loop.
R1	R2
1	2
2	3
3	4
4	5
5	6
6	7
*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
	"fmt"
)

// Method 2
func Loop(l *linkedList.LinkedList) *linkedList.Node {
	if l.Head == nil {
		return nil
	}
	slow := l.Head

	// Initially start Fast one node ahead for the condition on line 63 doesn't match,
	// since initially both will point to the same node
	fast := l.Head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return nil
		}
		slow = slow.Next      // Moves one jump
		fast = fast.Next.Next // Moves two jumps
	}

	// There is mathematical solution for it. I'm not sure if I understand that :/
	// https://stackoverflow.com/questions/2936213/explain-how-finding-cycle-start-node-in-cycle-linked-list-work
	// https://www.youtube.com/watch?v=-YiQZi3mLq0

	traverse := l.Head
	for slow != traverse {
		slow = slow.Next
		traverse = traverse.Next
	}
	return slow
}

func main() {
	l := linkedList.LinkedList{}
	fmt.Println(Loop(&l))
}
