/*
Palindrome: Implement a function to check if a linked list is a palindrome.


Method 1 -> Reverse and Compare
Convert the link elements into a string, reverse that string and compare it to the original string.

Method 2 -> Using Stack
Iterate over the linked list and push the elements onto a STACK.
Then we are going to iterate through once more and check whether or not the POPPED element is equal to the
element we are encountering in our second iteration.
*/

package main

import (
	linkedList "algorithms-in-golang/CTCI/Ch-2-LinkedList/SinglyLinkedList"
	"fmt"
)

type Stack struct {
	value interface{}
	container []interface{}
}

func(s *Stack) push(value interface{}) {
	s.container = append(s.container, value)
}

func(s *Stack) pop() interface{} {
	popped := s.container[len(s.container)-1]
	s.container = append(s.container[:len(s.container)-1])
	return popped
}


// Method 1, Reverse and Compare
func palindrome(l *linkedList.LinkedList) bool {
	// convert the link elements into a strings
	// node pointer
	current := l.Head
	var s string
	for current != nil {
		s+= current.Key.(string)
		current = current.Next
	}
	// reverse the string and compare it to the original string
	rvrse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}(s)

	if rvrse != s {
		return false
	}

	return true
}


// Method 2, Using Stack
func stackPalindome(l*linkedList.LinkedList) bool  {
	s := Stack{}
	current :=  l.Head
	// Push the elements into the stack
	for current != nil {
		s.push(current.Key)
		current = current.Next
	}

	// reinitialize current again to the head, because current will point to NIL after the first iteration and we will
	// never enter the second loop
	current = l.Head

	// Pop the element and compare it to the linked list elements
	for current != nil {
		pop := s.pop()
		if pop != current.Key {
			return false
		}
		current = current.Next
	}

	return true
}


func main() {
	l := linkedList.LinkedList{}
	values := []string{"a","b","c","d","c","b","a"}
	for _, v := range values {
		l.Append(string(v))
	}
	//Method 1
	fmt.Println(palindrome(&l))

	//Method 2
	fmt.Println(stackPalindome(&l))


}