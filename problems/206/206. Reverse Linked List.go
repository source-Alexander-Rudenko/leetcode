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

func main() {
	node4 := &ListNode{5, nil}
	node3 := &ListNode{4, node4}
	node2 := &ListNode{3, node3}
	node1 := &ListNode{2, node2}
	head := &ListNode{1, node1}

	printList(head)
	fmt.Println()
	printList(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
		//prev, head, head.Next = head, head.Next, prev
	}
	return prev
}
