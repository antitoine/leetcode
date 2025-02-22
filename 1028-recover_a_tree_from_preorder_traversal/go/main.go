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

func recoverFromPreorder(traversal string) *TreeNode {
	var stack []*TreeNode
	currentDepth, currentValue := 0, 0
	parsingValue := false

	for i, char := range traversal {
		if char == '-' {
			if parsingValue {
				newNode := &TreeNode{Val: currentValue}

				for len(stack) > currentDepth {
					stack = stack[:len(stack)-1]
				}

				if len(stack) > 0 {
					if stack[len(stack)-1].Left == nil {
						stack[len(stack)-1].Left = newNode
					} else {
						stack[len(stack)-1].Right = newNode
					}
				}

				stack = append(stack, newNode)

				currentDepth = 0
				currentValue = 0
				parsingValue = false
			}
			currentDepth++
		} else {
			currentValue = currentValue*10 + int(char-'0')
			parsingValue = true
		}

		if i == len(traversal)-1 {
			newNode := &TreeNode{Val: currentValue}

			for len(stack) > currentDepth {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				if stack[len(stack)-1].Left == nil {
					stack[len(stack)-1].Left = newNode
				} else {
					stack[len(stack)-1].Right = newNode
				}
			}

			stack = append(stack, newNode)
		}
	}

	var result *TreeNode
	if len(stack) > 0 {
		result = stack[0]
	}
	return result
}

func main() {
	traversal := "1-2--3--4-5--6--7"
	result := recoverFromPreorder(traversal)
	fmt.Printf("Result: %s\n", result)
	if result.String() == "{Val: 1, Left: {Val: 2, Left: {Val: 3, Left: <nil>, Right: <nil>}, Right: {Val: 4, Left: <nil>, Right: <nil>}}, Right: {Val: 5, Left: {Val: 6, Left: <nil>, Right: <nil>}, Right: {Val: 7, Left: <nil>, Right: <nil>}}}" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
