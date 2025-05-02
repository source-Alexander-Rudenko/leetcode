/*
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋
denotes the largest integer less than or equal to x.
For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
*/

package main

import "fmt"

func deleteMiddle(head *ListNode) *ListNode {
	dummyHead := &ListNode{-1, head}
	slow, fast := dummyHead, dummyHead.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next

}

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
	lis1 := ListNode{1, &ListNode{2, nil}}
	res := deleteMiddle(&lis1)
	printList(res)
}
