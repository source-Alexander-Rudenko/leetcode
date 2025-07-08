package main

import "fmt"

/*
A linked list of length n is given such that each node contains an additional random pointer,
 which could point to any node in the list, or null.
Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes,
 where each new node has its value set to the value of its corresponding original node.
  Both the next and random pointer of the new nodes should point to new nodes in the copied list
   such that the pointers in the original list and copied list represent the same list state.
    None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y,
 then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	oldToCopy := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		copy := &Node{Val: cur.Val}
		oldToCopy[cur] = copy
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		copy := oldToCopy[cur]
		copy.Next = oldToCopy[cur.Next]
		copy.Random = oldToCopy[cur.Random]
		cur = cur.Next
	}

	return oldToCopy[head]
}

func printList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ->", curr.Val)
		curr = curr.Next
	}
}

func createTestList() *Node {
	nodes := []*Node{
		{Val: 7},
		{Val: 13},
		{Val: 11},
		{Val: 10},
		{Val: 1},
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	nodes[1].Random = nodes[0] // 13 → 7
	nodes[2].Random = nodes[4] // 11 → 1
	nodes[3].Random = nodes[2] // 10 → 11
	nodes[4].Random = nodes[0] // 1  → 7

	return nodes[0]
}

func main() {
	head := createTestList()

	newHead := copyRandomList(head)

	printList(newHead)

}
