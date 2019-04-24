/*
Write a method to replace all spaces in a string with '%20'. You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string. (Note: If implementing in Java, please use a character array so that you can
perform this operation in place.)
EXAMPLE
Input: "Mr John Smith ", 13
Output: "Mr%20John%20Smith"

A common approach in string manipulation problems is to edit the string starting from the end and working
backwards.This is useful because we have an extra buffer at the end, which allows us to change characters
without worrying about what we're overwriting. We will use this approach in this problem.

The algorithm employs a two-scan approach. In the first scan, we count the number of spaces.
By tripling this number, we can compute how many extra characters we will have in the final string.
In the second pass, which is done in reverse order, we actually edit the string. When we see a space,
we replace it with %20. If there is no space, then we copy the original character.
*/

package main

import (
	"fmt"
)

/*
Strings are immutable in go so convert it into a rune.
*/
func urlify(a string, truelength int) string {
	space := 0
	r := []rune{}
	for i := 0; i < truelength; i++ {
		if string(a[i]) == " " {
			space++
		}
	}
	for _, v := range a {
		r = append(r, v)
	}
	originalLength := space*2 + truelength
	for i := truelength - 1; i >= 0; i-- {
		if r[i] == ' ' {
			r[originalLength-1] = '0'
			r[originalLength-2] = '2'
			r[originalLength-3] = '%'
			originalLength = originalLength - 3
		} else {
			r[originalLength-1] = r[i]
			originalLength--
		}
	}

	return string(r)
}

func main() {
	fmt.Println(urlify("Mr John Smith         ", 13))
}
