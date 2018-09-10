package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var next *ListNode

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				next = l1
				l1 = l1.Next
			} else {
				next = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			next = l1
			l1 = l1.Next
		} else {
			next = l2
			l2 = l2.Next
		}
		if tail != nil {
			tail.Next = next
			tail = next
		} else {
			head = next
			tail = next
		}
	}

	return head
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := mergeTwoLists(l1, l2)
}
