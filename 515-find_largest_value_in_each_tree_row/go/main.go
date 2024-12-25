package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	currentNodes := []*TreeNode{root}
	for len(currentNodes) > 0 {
		var nextNodes []*TreeNode
		maxValue := currentNodes[0].Val
		for _, node := range currentNodes {
			if node.Val > maxValue {
				maxValue = node.Val
			}
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		result = append(result, maxValue)
		currentNodes = nextNodes
	}

	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	result := largestValues(root)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[1 3 9]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
