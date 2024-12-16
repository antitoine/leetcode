package main

import (
	"container/heap"
	"fmt"
)

type Num struct {
	value int
	index int
}

type NumHeap []Num

func (h NumHeap) Len() int { return len(h) }
func (h NumHeap) Less(i, j int) bool {
	if h[i].value == h[j].value {
		return h[i].index < h[j].index
	}
	return h[i].value < h[j].value
}
func (h NumHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *NumHeap) Push(x any) {
	*h = append(*h, x.(Num))
}

func (h *NumHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getFinalState(nums []int, k int, multiplier int) []int {
	numsHeap := &NumHeap{}
	for i, num := range nums {
		heap.Push(numsHeap, Num{value: num, index: i})
	}

	for i := 0; i < k; i++ {
		num := heap.Pop(numsHeap).(Num)
		num.value *= multiplier
		heap.Push(numsHeap, num)
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num := heap.Pop(numsHeap).(Num)
		result[num.index] = num.value
	}

	return result
}

func main() {
	nums := []int{2, 1, 3, 5, 6}
	k := 5
	multiplier := 2
	result := getFinalState(nums, k, multiplier)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[8 4 6 5 6]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
