/*
Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a
rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
EXAMPLE
Input: Tact Coa
Output: True (permutations: "taco cat", "atco eta", etc.)

What does it take to be able to write a set of characters the same way forwards and backwards?
We need to have an even number of almost all characters, so that half can be on one side and half can be on the other
side.
Very Important :- At most one character (the middle character) can have an odd count.
eg: racecar
Palindrome := racecar
Permutaion := rac e rac
This one odd letter e will always be placed in the middle to make a permutation,
All letter on the left and right side of it will be same but arranged differently.


To be more precise, strings with even length (after removing all non-letter characters, like spaces)
must have all even counts of characters.
Strings of an odd length must have exactly one character with an odd count.
Of course, an "even" string can't have an odd number of exactly one character,
otherwise it wouldn't be an even-length string (an odd number+ many even numbers= an odd number).
Likewise, a string with odd length can't have all characters with even counts (sum of evens is even).
It's therefore sufficient to say that, to be a permutation ot a palindrome, a string can have
no more than one character that is odd. This will cover both the odd and the even cases.
*/

package main

import (
	"fmt"
	"strings"
)

func palindromePermutation(s string) bool {
	s = strings.ToLower(strings.Replace(s, " ", "", -1))
	oddCount := 0
	m := make(map[string]int)
	for _, v := range s {
		if _,ok := m[string(v)]; ok {
			m[string(v)]++
		} else {
			m[string(v)] = 1
		}
	}
	for _, v := range m {
		if v % 2 != 0 && oddCount == 1 {
			oddCount++
			if oddCount > 1 {
				return false
			}
		}
	}
	return true
}

func main() {

	fmt.Println(palindromePermutation("Tact Coa "))
}
