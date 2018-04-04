// Given a digit string, return all possible letter combinations that the number could represent.
// A mapping of digit to letters (just like on the telephone buttons) is given below.
// http://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png
//
// Input:Digit string "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

package main

import (
	"bytes"
	"fmt"
)

var keypad = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

type Node struct {
	Value    string
	Children []*Node
}

func (n *Node) String() string {
	var s bytes.Buffer
	if len(n.Children) > 0 {
		s.WriteString(fmt.Sprintf("(%s) --> ", n.Value))
		s.WriteString("[")
		for _, child := range n.Children {
			s.WriteString(child.String())
		}
		s.WriteString("]")
	} else {
		s.WriteString(fmt.Sprintf(" (%s) ", n.Value))
	}

	return s.String()
}

func buildTree(root *Node, digits string, currentDigitPosition int) {
	if currentDigitPosition >= len(digits) {
		return
	}

	keys := keypad[string(digits[currentDigitPosition])]

	for _, key := range keys {
		root.Children = append(root.Children, &Node{Value: string(key)})
	}
	for _, child := range root.Children {
		buildTree(child, digits, currentDigitPosition+1)
	}
}

func visitNodeLeaves(root *Node, prefix string) []string {
	result := make([]string, 0)

	_prefix := fmt.Sprintf("%s%s", prefix, root.Value)
	if len(root.Children) == 0 {
		if _prefix != "" {
			result = append(result, _prefix)
		}

		return result
	}

	for _, child := range root.Children {
		for _, r := range visitNodeLeaves(child, _prefix) {
			result = append(result, r)
		}
	}

	return result
}

func letterCombinations(digits string) []string {
	root := Node{}
	buildTree(&root, digits, 0)
	return visitNodeLeaves(&root, "")
}

func main() {
	result := letterCombinations("23323434982")
	fmt.Println(result)
}
