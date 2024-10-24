package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subNodesFlipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	subNodesRoot1 := []*TreeNode{}
	if root1.Left != nil {
		subNodesRoot1 = append(subNodesRoot1, root1.Left)
	}
	if root1.Right != nil {
		subNodesRoot1 = append(subNodesRoot1, root1.Right)
	}
	subNodesRoot2 := []*TreeNode{}
	if root2.Left != nil {
		subNodesRoot2 = append(subNodesRoot2, root2.Left)
	}
	if root2.Right != nil {
		subNodesRoot2 = append(subNodesRoot2, root2.Right)
	}
	if len(subNodesRoot1) != len(subNodesRoot2) {
		return false
	}
	sort.SliceStable(subNodesRoot1, func(i, j int) bool {
		return subNodesRoot1[i].Val < subNodesRoot1[j].Val
	})
	sort.SliceStable(subNodesRoot2, func(i, j int) bool {
		return subNodesRoot2[i].Val < subNodesRoot2[j].Val
	})
	for i := 0; i < len(subNodesRoot1); i++ {
		if subNodesRoot1[i].Val != subNodesRoot2[i].Val {
			return false
		}
		if !subNodesFlipEquiv(subNodesRoot1[i], subNodesRoot2[i]) {
			return false
		}
	}
	return true
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if !(root1 != nil && root2 != nil) {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return subNodesFlipEquiv(root1, root2)
}

func main() {
	input1 := &TreeNode{Val: 1}
	input1.Left = &TreeNode{Val: 2}
	input1.Right = &TreeNode{Val: 3}
	input1.Left.Left = &TreeNode{Val: 4}
	input1.Left.Right = &TreeNode{Val: 5}
	input1.Right.Left = &TreeNode{Val: 6}
	input1.Left.Right.Left = &TreeNode{Val: 7}
	input1.Left.Right.Right = &TreeNode{Val: 8}
	input2 := &TreeNode{Val: 1}
	input2.Left = &TreeNode{Val: 3}
	input2.Right = &TreeNode{Val: 2}
	input2.Left.Right = &TreeNode{Val: 6}
	input2.Right.Left = &TreeNode{Val: 4}
	input2.Right.Right = &TreeNode{Val: 5}
	input2.Right.Right.Left = &TreeNode{Val: 8}
	input2.Right.Right.Right = &TreeNode{Val: 7}
	result := flipEquiv(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
