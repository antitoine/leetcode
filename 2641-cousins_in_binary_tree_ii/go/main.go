package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Node   *TreeNode
	Parent *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*NodeInfo{{Node: root, Parent: nil}}
	levelNodes := [][]*NodeInfo{}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []*NodeInfo{}

		for i := 0; i < levelSize; i++ {
			nodeInfo := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, nodeInfo)

			if nodeInfo.Node.Left != nil {
				queue = append(queue, &NodeInfo{Node: nodeInfo.Node.Left, Parent: nodeInfo.Node})
			}
			if nodeInfo.Node.Right != nil {
				queue = append(queue, &NodeInfo{Node: nodeInfo.Node.Right, Parent: nodeInfo.Node})
			}
		}

		levelNodes = append(levelNodes, currentLevel)
	}

	for _, level := range levelNodes[1:] {
		totalSum := 0
		parentSum := make(map[*TreeNode]int)

		for _, nodeInfo := range level {
			totalSum += nodeInfo.Node.Val
			parentSum[nodeInfo.Parent] += nodeInfo.Node.Val
		}

		for _, nodeInfo := range level {
			nodeInfo.Node.Val = totalSum - parentSum[nodeInfo.Parent]
		}
	}

	root.Val = 0
	return root
}

func main() {
	input := &TreeNode{Val: 5}
	input.Left = &TreeNode{Val: 4}
	input.Right = &TreeNode{Val: 9}
	input.Left.Left = &TreeNode{Val: 1}
	input.Left.Right = &TreeNode{Val: 10}
	input.Right.Right = &TreeNode{Val: 7}
	result := replaceValueInTree(input)
	fmt.Printf("Result: %#v\n", result)
	if result.Val == 0 && result.Left.Val == 0 && result.Right.Val == 0 && result.Left.Left.Val == 7 && result.Left.Right.Val == 7 && result.Right.Right.Val == 11 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
