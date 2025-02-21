package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	visited map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{visited: make(map[int]struct{})}
	if root != nil {
		root.Val = 0
		fe.recoverTree(root)
	}
	return fe
}

func (fe *FindElements) recoverTree(node *TreeNode) {
	fe.visited[node.Val] = struct{}{}

	if node.Left != nil {
		node.Left.Val = node.Val*2 + 1
		fe.recoverTree(node.Left)
	}

	if node.Right != nil {
		node.Right.Val = node.Val*2 + 2
		fe.recoverTree(node.Right)
	}
}

func (fe *FindElements) Find(target int) bool {
	_, found := fe.visited[target]
	return found
}

func main() {
	root := &TreeNode{
		Val: -1,
		Right: &TreeNode{
			Val: -1,
		},
	}
	obj := Constructor(root)
	if obj.Find(1) {
		fmt.Println("Find 1 failed, expected false, got true")
		return
	}
	if !obj.Find(2) {
		fmt.Println("Find 2 failed, expected true, got false")
		return
	}
	fmt.Println("Passed")
}
