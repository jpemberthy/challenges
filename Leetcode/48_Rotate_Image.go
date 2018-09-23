package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)

	// total number of "matrices" aka number of spiral cycles
	m := n / 2
	for i := 0; i < m; i++ {
		limit := n - i - 1
		for j := i; j < limit; j++ {
			// copy right and move top to right
			copyRight := matrix[j][n-i-1]
			matrix[j][n-i-1] = matrix[i][j]

			// copy bottom and move copy of right to bottom
			copyBottom := matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = copyRight

			// copy left and move copy of bottom to left
			copyLeft := matrix[n-j-1][i]
			matrix[n-j-1][i] = copyBottom

			// move copy of left to top
			matrix[i][j] = copyLeft
		}
	}
}

func main() {
	matrix := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}

	rotate(matrix)
	fmt.Println(matrix)

}
