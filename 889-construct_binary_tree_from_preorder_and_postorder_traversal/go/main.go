package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("{Val: %d, Left: %s, Right: %s}", t.Val, t.Left, t.Right)
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postOrderIndexMap := make(map[int]int)
	for i, val := range postorder {
		postOrderIndexMap[val] = i
	}
	return buildTree(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1, postOrderIndexMap)
}

func buildTree(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int, postOrderIndexMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[preStart]}
	if preStart == preEnd {
		return root
	}
	leftRootIndex := postOrderIndexMap[preorder[preStart+1]]
	leftSubtreeLength := leftRootIndex - postStart + 1
	root.Left = buildTree(preorder, preStart+1, preStart+leftSubtreeLength, postorder, postStart, leftRootIndex, postOrderIndexMap)
	root.Right = buildTree(preorder, preStart+leftSubtreeLength+1, preEnd, postorder, leftRootIndex+1, postEnd-1, postOrderIndexMap)
	return root
}

func main() {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	result := constructFromPrePost(preorder, postorder)
	fmt.Printf("Result: %s\n", result)
	if result.String() == "{Val: 1, Left: {Val: 2, Left: {Val: 4, Left: <nil>, Right: <nil>}, Right: {Val: 5, Left: <nil>, Right: <nil>}}, Right: {Val: 3, Left: {Val: 6, Left: <nil>, Right: <nil>}, Right: {Val: 7, Left: <nil>, Right: <nil>}}}" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
