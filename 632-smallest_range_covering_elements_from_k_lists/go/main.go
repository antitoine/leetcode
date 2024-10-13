package main

import (
	"container/heap"
	"fmt"
)

type Constraint struct {
	listIdx int
	itemIdx int
	itemVal int
}

type MinConstraints []Constraint

func (h MinConstraints) Len() int           { return len(h) }
func (h MinConstraints) Less(i, j int) bool { return h[i].itemVal < h[j].itemVal }
func (h MinConstraints) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinConstraints) Push(x any) {
	*h = append(*h, x.(Constraint))
}

func (h *MinConstraints) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxConstraints []Constraint

func (h MaxConstraints) Len() int           { return len(h) }
func (h MaxConstraints) Less(i, j int) bool { return h[i].itemVal > h[j].itemVal }
func (h MaxConstraints) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxConstraints) Push(x any) {
	*h = append(*h, x.(Constraint))
}

func (h *MaxConstraints) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	minConstraints := &MinConstraints{}
	heap.Init(minConstraints)
	maxConstraints := &MaxConstraints{}
	heap.Init(maxConstraints)

	for idx, list := range nums {
		heap.Push(minConstraints, Constraint{idx, 0, list[0]})
		heap.Push(maxConstraints, Constraint{idx, 0, list[0]})
	}

	minConstraint := heap.Pop(minConstraints).(Constraint)
	maxConstraint := heap.Pop(maxConstraints).(Constraint)
	smallestRangeLength := maxConstraint.itemVal - minConstraint.itemVal
	smallestRangeMin := minConstraint.itemVal
	smallestRangeMax := maxConstraint.itemVal
	for {
		if minConstraint.itemIdx < len(nums[minConstraint.listIdx])-1 {
			heap.Push(minConstraints, Constraint{minConstraint.listIdx, minConstraint.itemIdx + 1, nums[minConstraint.listIdx][minConstraint.itemIdx+1]})
			if nums[minConstraint.listIdx][minConstraint.itemIdx+1] > maxConstraint.itemVal {
				heap.Push(maxConstraints, Constraint{minConstraint.listIdx, minConstraint.itemIdx + 1, nums[minConstraint.listIdx][minConstraint.itemIdx+1]})
				maxConstraint = heap.Pop(maxConstraints).(Constraint)
			}
			minConstraint = heap.Pop(minConstraints).(Constraint)
			if maxConstraint.itemVal-minConstraint.itemVal < smallestRangeLength {
				smallestRangeLength = maxConstraint.itemVal - minConstraint.itemVal
				smallestRangeMin = minConstraint.itemVal
				smallestRangeMax = maxConstraint.itemVal
			}
		} else {
			break
		}
	}

	return []int{smallestRangeMin, smallestRangeMax}
}

func main() {
	nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	result := smallestRange(nums)
	fmt.Printf("Result: %#v\n", result)
	if result[0] == 20 && result[1] == 24 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
