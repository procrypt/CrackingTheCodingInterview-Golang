/*
Check Permutation: Given two strings, write a method to decide if one is a permutation of the
other.

Observe first that strings of different lengths cannot be permutations of each other.
If two strings are permutations, then we know they have the same characters, but in different orders.
Therefore, sorting the strings will put the characters from two permutations in the same order. We just need to
compare the sorted versions of the strings.

How to sort a string in golang
Go 1.8 or before
The steps:
Convert string to []rune
Define a new type ByRune, which is actually []rune. Implement Len, Swap, and Less methods of type ByRune.
Call sort.Sort to sort the slice of runes.
Convert []rune back to string and return the string.

Go 1.8 or later
The sort package in Go 1.8 introduces new methods for sorting slices.
We can use sort.Slice method directly without defining a new type.
Convert string to []rune
Define a less method and call sort.Slice() with the slice of runes and the less method as parameters.
 sort.Slice(r, func(i, j int) bool {
              return r[i] < r[j]
      })
Convert []rune back to string and return the string.
*/

package main

import (
	"fmt"
	"sort"
)

func stringsToRunes(s string) []rune {
	r := []rune{}
	for _, v := range s {
			r = append(r, rune(v))
	}
	return r
}

func Permutation(a,b string) bool {
	if len(a) != len(b) {
		return false
	}
	c := stringsToRunes(a)
	d := stringsToRunes(b)
	sort.SliceStable(c, func(i, j int) bool {
		 return c[i] < c[j]
	})
	sort.SliceStable(d, func(i, j int) bool {
		 return d[i] < d[j]
	})
	if string(c) != string(d) {
		return false
	}
	return true
}

func main() {
	fmt.Println(Permutation("abcd","bcda"))
}