package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minOperations(nums []int, k int) int {
	numsHeap := IntHeap(nums)
	queue := &numsHeap
	heap.Init(queue)
	var nbOp int
	for queue.Len() > 1 && (*queue)[0] < k {
		x := heap.Pop(queue).(int)
		y := heap.Pop(queue).(int)
		heap.Push(queue, min(x, y)*2+max(x, y))
		nbOp++
	}
	return nbOp
}

func main() {
	nums := []int{2, 11, 10, 1, 3}
	k := 10
	result := minOperations(nums, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 2 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
