/*
String Rotation: Assume you have a method isSubstring which checks if one word is a substring
of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one
call to isSubstring (e.g., "waterbottle" is a rotation of"erbottlewat").

The trick is S1+S1 containers all the rotations which can be found by rotating S1.
So we just have to check is S2 is in S1+S2
eg.
S1    = waterbottle
S2    = erbottlewat
S1+S1 = waterbottlewaterbottle
We can find S2 in S1+S1.
*/

package main

import (
	"fmt"
	"strings"
)

func isSubstring(s1, s2 string) bool {
	if !strings.Contains(s1, s2) {
		return false
	}
	return true
}

func rotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		s1 = s1+s1
		return isSubstring(s1, s2)
	}
	 return true
}

func main() {

	fmt.Println(rotation("abhishek", "shekabhi"))
}