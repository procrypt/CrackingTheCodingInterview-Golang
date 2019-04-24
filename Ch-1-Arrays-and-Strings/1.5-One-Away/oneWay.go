/*
There are three types of edits that can be performed on strings: insert a character,
remove a character, or replace a character. Given two strings, write a function to check if they are
one edit (or zero edits) away.
EXAMPLE
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false
*/

package main

import "fmt"

func oneWay(a, b string) bool {
	m := make(map[string]int)

	for _, v := range a {
		if _, ok := m[string(v)]; ok {
			m[string(v)]++
		} else {
			m[string(v)] = 1
		}
	}

	// case 1 insert
	if len(a) > len(b) {
		for i := 0; i < len(b); i++ {
			if _, ok := m[string(b[i])]; ok {
				m[string(b[i])]--
				if m[string(b[i])] == 0 {
					delete(m, string(b[i]))
				}
			} else {
				continue
			}
		}
		fmt.Println("case 1:", m)
		if len(m) == 1 {
			return true
		}
	}

	// case 2 remove
	if len(a) < len(b) {
		for i := 0; i < len(a); i++ {
			if _, ok := m[string(b[i])]; ok {
				m[string(b[i])]--
				if m[string(b[i])] == 0 {
					delete(m, string(b[i]))
				}
			} else {
				continue
			}
		}
		fmt.Println("case 2:", m)
		if len(m) == 0 {
			return true
		}
	}

	// case 3 replace
	if len(a) == len(b) {
		for i := 0; i < len(b); i++ {
			if _, ok := m[string(b[i])]; ok {
				m[string(b[i])]--
				if m[string(b[i])] == 0 {
					delete(m, string(b[i]))
				}
			} else {
				continue
			}
		}
		fmt.Println("case 3:", m)
		if len(m) == 1 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(oneWay("pale", "nake"))
}
