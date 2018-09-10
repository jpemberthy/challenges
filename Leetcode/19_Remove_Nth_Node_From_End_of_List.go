package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := make([]*ListNode, 0)

	for i := 0; head != nil; i++ {
		list = append(list, head)
		head = head.Next
	}

	lenght := len(list)
	nth := lenght - n
	if nth > 0 {
		head = list[0]
		list[nth-1].Next = list[nth].Next
	} else if lenght > 1 {
		head = list[1]
	} else {
		head = nil
	}

	return head
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	removeNthFromEnd(l1, 2)
}
