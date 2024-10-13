package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Friend struct {
	id      int
	arrival int
	leaving int
	chair   int
}

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

func smallestChair(times [][]int, targetFriend int) int {
	arrivalTimes := make([]*Friend, len(times))
	leavingTimes := make([]*Friend, len(times))
	for id, details := range times {
		arrivalTimes[id] = &Friend{id, details[0], details[1], -1}
		leavingTimes[id] = arrivalTimes[id]
	}
	sort.Slice(arrivalTimes, func(i, j int) bool {
		return arrivalTimes[i].arrival < arrivalTimes[j].arrival
	})
	sort.Slice(leavingTimes, func(i, j int) bool {
		return leavingTimes[i].leaving < leavingTimes[j].leaving
	})
	var friend *Friend
	arrivalIdx, leavingIdx := 0, 0
	availableChairs := &IntHeap{}
	heap.Init(availableChairs)
	lastChairIdx := 0
	for friend == nil || friend.id != targetFriend {
		if arrivalIdx < len(arrivalTimes) && arrivalTimes[arrivalIdx].arrival < leavingTimes[leavingIdx].leaving {
			friend = arrivalTimes[arrivalIdx]
			arrivalIdx++
			if len(*availableChairs) > 0 {
				friend.chair = heap.Pop(availableChairs).(int)
			} else {
				friend.chair = lastChairIdx
				lastChairIdx++
			}
		} else {
			friend = leavingTimes[leavingIdx]
			leavingIdx++
			heap.Push(availableChairs, friend.chair)
		}
	}
	return friend.chair
}

func main() {
	input1, input2 := [][]int{{1, 4}, {2, 3}, {4, 6}}, 1
	result := smallestChair(input1, input2)
	fmt.Printf("Result: %#v\n", result)
	if result == 1 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
