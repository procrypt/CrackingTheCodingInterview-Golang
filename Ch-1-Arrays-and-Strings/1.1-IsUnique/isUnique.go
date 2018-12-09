package main

import "fmt"

/*
Is Unique: Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/

// What if you cannot use additional data structures? O(n^2)
func isUniqueInefficient(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

// O(n)
func isUnique(s string) bool {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok {
			return false
		} else {
			m[string(s[i])]++
		}
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcdefghijklmnopqrstuvwxyza"))
}
