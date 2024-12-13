package main

import (
	"container/heap"
	"fmt"
)

type Num struct {
	Val        int
	InitialIdx int
}

type Nums []*Num

func (h Nums) String() string {
	str := "["
	for i, num := range h {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", num)
	}
	str += "]"
	return str
}

func (h Nums) Len() int { return len(h) }
func (h Nums) Less(i, j int) bool {
	if h[i].Val == h[j].Val {
		return h[i].InitialIdx < h[j].InitialIdx
	} else {
		return h[i].Val < h[j].Val
	}
}
func (h Nums) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Nums) Push(x any) {
	*h = append(*h, x.(*Num))
}

func (h *Nums) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findScore(nums []int) int64 {
	numsHeap := make(Nums, len(nums))
	numsHeapP := &numsHeap
	for i, num := range nums {
		numsHeap[i] = &Num{Val: num, InitialIdx: i}
	}
	heap.Init(numsHeapP)

	var score int64
	numsIdxMarked := make(map[int]struct{})
	for len(numsIdxMarked) < len(nums) {
		item := heap.Pop(numsHeapP).(*Num)
		if _, ok := numsIdxMarked[item.InitialIdx]; ok {
			continue
		}
		score += int64(item.Val)
		numsIdxMarked[item.InitialIdx] = struct{}{}
		beforeIdx := item.InitialIdx - 1
		if beforeIdx >= 0 {
			numsIdxMarked[beforeIdx] = struct{}{}
		}
		afterIdx := item.InitialIdx + 1
		if afterIdx < len(nums) {
			numsIdxMarked[afterIdx] = struct{}{}
		}
	}
	return score
}

func main() {
	nums := []int{2, 1, 3, 4, 5, 2}
	result := findScore(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == 7 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
