package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Result struct {
	balanced bool
	height   int
}

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) Result
	dfs = func(root *TreeNode) Result {
		if root == nil {
			return Result{true, 0}
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		balanced := left.balanced && right.balanced && intAbs(left.height-right.height) <= 1
		return Result{balanced, 1 + intMax(left.height, right.height)}
	}
	return dfs(root).balanced
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 4}

	fmt.Println(isBalanced(root))
}

//TODO в obsidian записать отгадку
