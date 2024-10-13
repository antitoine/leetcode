package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middleNode := head
	currentNode := head
	for i := 0; currentNode.Next != nil; i++ {
		currentNode = currentNode.Next
		if i%2 == 0 {
			middleNode = middleNode.Next
		}
	}
	return middleNode
}

func main() {
	// [1,2,3,4,5]
	input := &ListNode{Val: 1}
	input.Next = &ListNode{Val: 2}
	input.Next.Next = &ListNode{Val: 3}
	input.Next.Next.Next = &ListNode{Val: 4}
	input.Next.Next.Next.Next = &ListNode{Val: 5}
	result := middleNode(input)
	fmt.Printf("Result: %#v\n", result)
	// [3,4,5]
	if result.Val == 3 && result.Next.Val == 4 && result.Next.Next.Val == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
