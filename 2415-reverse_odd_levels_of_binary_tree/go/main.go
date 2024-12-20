package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	if n == nil {
		return ""
	}
	if n.Left == nil && n.Right == nil {
		return fmt.Sprintf("%d", n.Val)
	}
	return fmt.Sprintf("%d{%s,%s}", n.Val, n.Left, n.Right)
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		values := make([]int, 0, levelSize)
		nodes := make([]*TreeNode, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if level%2 == 1 {
				values = append(values, node.Val)
			}
			nodes = append(nodes, node)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if level%2 == 1 {
			for i := 0; i < len(nodes); i++ {
				nodes[i].Val = values[len(values)-1-i]
			}
		}

		level++
	}

	return root
}

func main() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
	}
	result := reverseOddLevels(root)
	fmt.Printf("Result: %s\n", result)
	if result.String() == "0{2{0{2,2},0{2,2}},1{0{1,1},0{1,1}}}" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
