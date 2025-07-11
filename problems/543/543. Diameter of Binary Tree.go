package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := height(root.Left)
		rightHeight := height(root.Right)

		if leftHeight+rightHeight > diameter {
			diameter = leftHeight + rightHeight
		}

		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}
	height(root)
	return diameter
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

	fmt.Print(diameterOfBinaryTree(root))
}
