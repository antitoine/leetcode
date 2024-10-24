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

type MaxSum []int64

func (h MaxSum) Len() int           { return len(h) }
func (h MaxSum) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxSum) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func computeSubSum(nodeLevel int, node *TreeNode, levelSum map[int]int64) int {
	levelSum[nodeLevel] += int64(node.Val)
	deepestLevel := nodeLevel
	if node.Left != nil {
		leftDeepestLevel := computeSubSum(nodeLevel+1, node.Left, levelSum)
		if leftDeepestLevel > deepestLevel {
			deepestLevel = leftDeepestLevel
		}
	}
	if node.Right != nil {
		rightDeepestLevel := computeSubSum(nodeLevel+1, node.Right, levelSum)
		if rightDeepestLevel > deepestLevel {
			deepestLevel = rightDeepestLevel
		}
	}
	return deepestLevel
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levelSum := make(map[int]int64)
	deepestLevel := computeSubSum(1, root, levelSum)

	if k > deepestLevel {
		return -1
	}

	maxSum := make(MaxSum, deepestLevel)
	for level, sum := range levelSum {
		maxSum[level-1] = sum
	}
	sort.Sort(maxSum)

	return maxSum[k-1]
}

func main() {
	input1 := &TreeNode{Val: 5}
	input1.Left = &TreeNode{Val: 8}
	input1.Right = &TreeNode{Val: 9}
	input1.Left.Left = &TreeNode{Val: 2}
	input1.Left.Right = &TreeNode{Val: 1}
	input1.Right.Left = &TreeNode{Val: 3}
	input1.Right.Right = &TreeNode{Val: 7}
	input1.Left.Left.Left = &TreeNode{Val: 4}
	input1.Left.Left.Right = &TreeNode{Val: 6}
	input2 := 2
	result := kthLargestLevelSum(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == 13 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
