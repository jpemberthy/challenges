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
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// String returns a string representation of the list
func (l *ListNode) String() string {
	var string bytes.Buffer
	for l != nil {
		string.WriteString(strconv.Itoa(l.Val))
		l = l.Next
	}

	return string.String()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultList := &ListNode{}
	head := resultList

	carry := 0
	for l1 != nil || l2 != nil {
		d1 := 0
		d2 := 0

		if l1 != nil {
			d1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			d2 = l2.Val
			l2 = l2.Next
		}

		sum := d1 + d2 + carry
		if sum >= 10 {
			sum = sum % 10
			carry = 1
		} else {
			carry = 0
		}

		resultList.Val = sum
		if l1 == nil && l2 == nil {
			if carry > 0 {
				resultList.Next = &ListNode{Val: 1}
			}
			return head
		}

		resultList.Next = &ListNode{}
		resultList = resultList.Next
	}

	return head
}

func main() {
	d1 := &ListNode{Val: 8}
	d2 := &ListNode{Val: 1, Next: d1}

	d3 := &ListNode{Val: 0}

	fmt.Println(addTwoNumbers(d2, d3))
}
