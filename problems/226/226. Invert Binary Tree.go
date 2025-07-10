package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	PrintTree(root.Left)
	PrintTree(root.Right)
	fmt.Printf("%d ", root.Val)
}

func main() {
	r := &TreeNode{Val: 4}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 7}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Right = &TreeNode{Val: 9}
	r.Left.Left = &TreeNode{Val: 1}
	r.Left.Right = &TreeNode{Val: 3}
	PrintTree(r)
	fmt.Println()
	r = invertTree(r)
	PrintTree(r)
}
