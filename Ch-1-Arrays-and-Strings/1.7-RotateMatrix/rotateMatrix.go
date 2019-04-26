/*
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

1) Reverse the matrix
2) Take transpose of the reversed matrix
   In transpose row becomes column and column becomes row

1 2 3     7 8 9      7 4 1
4 5 6 --> 4 5 6  --> 8 5 2
7 8 9     1 2 3      9 6 3

*/

package main

import "fmt"

func rotateMatrix(matrix [][]int) [][]int {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i<j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func main() {
	fmt.Println(rotateMatrix([][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}))
}
