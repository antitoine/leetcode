package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	pos1 := head
	pos2 := head.Next
	for i := 1; pos2.Next != nil && pos1 != pos2; i++ {
		pos2 = pos2.Next
		if i%3 == 0 {
			pos1 = pos1.Next
		}
	}
	return pos1 == pos2
}

func main() {
	input := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	input.Next.Next.Next.Next = input.Next
	result := hasCycle(input.Next)
	fmt.Printf("Result: %#v\n", result)
	if result == true {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
