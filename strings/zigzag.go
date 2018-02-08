// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string text, int nRows);
// convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	n := len(s)
	matrix := make([][]string, numRows)
	for i := range matrix {
		matrix[i] = make([]string, 0)
	}

	k := 0
	for k < n {
		for i := 0; i < numRows && k < n; i++ {
			matrix[i] = append(matrix[i], string(s[k]))
			k += 1
		}

		for i := numRows - 2; i > 0 && k < n; i-- {
			matrix[i] = append(matrix[i], string(s[k]))
			k += 1
		}
	}

	var output bytes.Buffer
	for _, level := range matrix {
		for _, char := range level {
			output.WriteString(string(char))
		}
	}

	return output.String()
}

// main ...
func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
