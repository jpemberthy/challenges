package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make(map[int]*ListNode)

	for i := 0; head != nil; i++ {
		nodes[i] = head
		head = head.Next
	}

	lenght := len(nodes)
	nth := lenght - n
	if nth > 0 {
		head = nodes[0]
		nodes[nth-1].Next = nodes[nth].Next
	} else if lenght > 1 {
		head = nodes[1]
	} else {
		head = nil
	}

	return head
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	removeNthFromEnd(l1, 2)
}
