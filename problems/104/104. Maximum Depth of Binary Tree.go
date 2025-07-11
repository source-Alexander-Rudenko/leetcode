package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1

}

func TreePrint(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	if node.Right != nil {
		TreePrint(node.Right, prefix+"│   ", false)
	}
	fmt.Print(prefix)

	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(node.Val)

	if node.Left != nil {
		TreePrint(node.Left, prefix+"|  ", true)
	}
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 17}
	root.Right.Right = &TreeNode{Val: 7}

	TreePrint(root, "", false)

	fmt.Print(maxDepth(root))

}
