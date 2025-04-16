package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	return
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	seccond := slow.Next
	slow.Next = nil
	var prev *ListNode
	for seccond != nil {
		tmp := seccond.Next
		seccond.Next = prev
		prev = seccond
		seccond = tmp
	}
	first := head
	seccond = prev
	for seccond != nil {
		tmp1, tmp2 := first.Next, seccond.Next
		first.Next = seccond
		seccond.Next = tmp1
		first, seccond = tmp1, tmp2
	}
}

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	printList(head)
}
