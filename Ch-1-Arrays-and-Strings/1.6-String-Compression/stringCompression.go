/*
String Compression: Implement a method to perform basic string compression using the counts
of repeated characters. For example, the string aabcccccaaa would become a2blc5a3.
If the "compressed" string would not become smaller than the original string, your method should return
the original string. You can assume the string has only uppercase and lowercase letters (a - z).
*/

package main

import (
	"fmt"
	"strconv"
)

func compression(s string) string {
	count := 1
	str := ""
	for i:=0; i<len(s)-1; i++ {
		if string(s[i]) == string(s[i+1]) {
			count++
		} else {
			c := strconv.Itoa(count)
			str+= string(s[i])+c
			count = 1
		}
	}
	c := strconv.Itoa(count)
	// This is very important step as out loop exited before and we didn't hit the else part, we manually have to
	// append the last set of characters in the string.
	str+= string(s[len(s)-1])+c

	// If the "compressed" string would not become smaller than the original string, your method should return
	//the original string.
	if len(str) > len(s) {
		return s
	}
	return str
}



func main() {

	fmt.Println(compression("aaaabccdddeeee"))
}
