package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func main() {
	a := &TreeNode{Val: 1}
	b := &TreeNode{Val: 1}
	a.Left = &TreeNode{Val: 2}
	a.Right = &TreeNode{Val: 3}
	b.Left = &TreeNode{Val: 2}
	b.Right = &TreeNode{Val: 3}
	fmt.Println(isSameTree(a, b))
}

//TODO тоже в obsisian
