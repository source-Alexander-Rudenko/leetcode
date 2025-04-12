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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next, list1 = list1, list1.Next
		} else {
			tail.Next, list2 = list2, list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return head.Next
}

func main() {
	head1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	head2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	printList(mergeTwoLists(head1, head2))
}
