package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	calculate("1+2*3")
}

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	numbers := regexp.MustCompile(`\d+`).FindAllString(s, -1)
	operators := regexp.MustCompile(`\D+`).FindAllString(s, -1)

	oIndex := 0
	i := 0
	for i < len(numbers) {
		currentOperator := operators[oIndex]
		nextOperator := operators[oIndex]
	}

	fmt.Println(numbers, operators)
	return 0
}
