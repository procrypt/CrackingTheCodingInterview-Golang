/*
One Away: There are three types of edits that can be performed on strings: insert a character,
remove a character, or replace a character. Given two strings, write a function to check if they are
one edit (or zero edits) away.
EXAMPLE
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false
*/


package main

import (
	"fmt"
)

func stringToRune(s string) []rune {
	r := []rune{}
	for _, v := range s {
		r = append(r, v)
	}
	return r
}

func oneAway(a,b string) bool {

	// Insert
	if len(a) < len(b) {
		return true
	}
	// Remove
	if len(a) > len(b) {
		return true
	}

	// replace

	return false
}

func main() {

		fmt.Println(oneAway("pale", "bake"))
}
