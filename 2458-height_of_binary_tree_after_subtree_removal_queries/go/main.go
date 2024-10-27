package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	depthMap := make(map[*TreeNode]int)

	var calculateDepth func(node *TreeNode) int
	calculateDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := calculateDepth(node.Left)
		rightDepth := calculateDepth(node.Right)
		depthMap[node] = 1 + max(leftDepth, rightDepth)
		return depthMap[node]
	}
	calculateDepth(root)

	maximumRest := make([]int, len(depthMap)+1)
	var dfs func(node *TreeNode, depth, rest int)
	dfs = func(node *TreeNode, depth, rest int) {
		if node == nil {
			return
		}
		depth++
		maximumRest[node.Val] = rest

		if node.Left != nil {
			dfs(node.Left, depth, max(rest, depth+depthMap[node.Right]))
		}
		if node.Right != nil {
			dfs(node.Right, depth, max(rest, depth+depthMap[node.Left]))
		}
	}
	dfs(root, -1, 0)

	answers := make([]int, len(queries))
	for i, value := range queries {
		answers[i] = maximumRest[value]
	}
	return answers
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 7}
	queries := []int{4}
	result := treeQueries(root, queries)
	fmt.Printf("Result: %#v\n", result)
	if result[0] == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
