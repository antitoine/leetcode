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

func minimumOperations(root *TreeNode) int {
	levels := make(map[int][]int)
	level := 0
	currentNodes := []*TreeNode{root}
	for len(currentNodes) > 0 {
		var nextNodes []*TreeNode
		for _, node := range currentNodes {
			levels[level] = append(levels[level], node.Val)
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		currentNodes = nextNodes
		level++
	}

	var result int
	for _, values := range levels {
		result += minimumSwapsToSort(values)
	}

	return result
}

func minimumSwapsToSort(arr []int) int {
	n := len(arr)
	type Pair struct {
		Value, Index int
	}

	indexedArr := make([]Pair, n)
	for i, v := range arr {
		indexedArr[i] = Pair{Value: v, Index: i}
	}

	sort.Slice(indexedArr, func(i, j int) bool {
		return indexedArr[i].Value < indexedArr[j].Value
	})

	visited := make([]bool, n)
	swaps := 0
	for i := 0; i < n; i++ {
		if visited[i] || indexedArr[i].Index == i {
			continue
		}
		cycleSize := 0
		x := i
		for !visited[x] {
			visited[x] = true
			x = indexedArr[x].Index
			cycleSize++
		}
		if cycleSize > 1 {
			swaps += cycleSize - 1
		}
	}

	return swaps
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 10,
				},
			},
		},
	}
	result := minimumOperations(root)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
