// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
package main

import (
	"bytes"
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listToInt(l *ListNode) *big.Int {
	var digits []string

	for l != nil {
		digits = append(digits, strconv.Itoa(l.Val))
		l = l.Next
	}

	var number bytes.Buffer
	length := len(digits) - 1
	for i := length; i >= 0; i-- {
		number.WriteString(digits[i])
	}

	n := new(big.Int)
	n.SetString(number.String(), 10)
	return n
}

func intToList(n *big.Int) *ListNode {
	node := &ListNode{}

	digits := n.String()
	length := len(digits)
	for i := 0; i < len(digits); i++ {
		node.Val = int(digits[i] - '0')
		if i+1 < length {
			previous := &ListNode{Next: node}
			node = previous
		}
	}
	return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// big.Int are used to support huge inputs.
	n1 := listToInt(l1)
	sum := n1.Add(n1, listToInt(l2))
	return intToList(sum)
}

func main() {
	d1 := &ListNode{Val: 3}
	d2 := &ListNode{Val: 4, Next: d1}
	d3 := &ListNode{Val: 2, Next: d2}

	d4 := &ListNode{Val: 4}
	d5 := &ListNode{Val: 6, Next: d4}
	d6 := &ListNode{Val: 5, Next: d5}

	result := addTwoNumbers(d3, d6)
	fmt.Println(listToInt(result))
}
