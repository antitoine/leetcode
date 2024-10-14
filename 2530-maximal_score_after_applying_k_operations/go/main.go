package main

import (
	"container/heap"
	"fmt"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func maxKelements(nums []int, k int) int64 {
	numsHeap := IntHeap(nums)
	heap.Init(&numsHeap)
	var score int64
	for i := 0; i < k; i++ {
		val := heap.Pop(&numsHeap).(int)
		score += int64(val)
		newVal := int(math.Ceil(float64(val) / 3))
		heap.Push(&numsHeap, newVal)
	}
	return score
}

func main() {
	input1, input2 := []int{10, 10, 10, 10, 10}, 5
	result := maxKelements(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == 50 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
