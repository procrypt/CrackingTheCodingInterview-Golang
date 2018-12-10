package main

import (
	"fmt"
	"strings"
)

/*
URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string.
 */

func URLify(s string) string {
	urlify := strings.Replace(strings.TrimRight(s, " ")," ", "%20", -1)
	fmt.Println(len(urlify))
	return urlify
}


func main() {
	fmt.Println(URLify("abhishek pratap singh    "))
}
