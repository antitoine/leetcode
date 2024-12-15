package main

import (
	"container/heap"
	"fmt"
)

type Class struct {
	PassCount         int
	TotalCount        int
	PotentialIncrease float64
}

type PriorityQueue []Class

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].PotentialIncrease > pq[j].PotentialIncrease
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Class))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, cls := range classes {
		passCount := cls[0]
		totalCount := cls[1]
		potentialIncrease := float64(passCount+1)/float64(totalCount+1) - float64(passCount)/float64(totalCount)
		heap.Push(pq, Class{passCount, totalCount, potentialIncrease})
	}

	for extraStudents > 0 {
		extraStudents--
		top := heap.Pop(pq).(Class)
		top.PassCount++
		top.TotalCount++
		top.PotentialIncrease = float64(top.PassCount+1)/float64(top.TotalCount+1) - float64(top.PassCount)/float64(top.TotalCount)
		heap.Push(pq, top)
	}

	totalRatio := 0.0
	for pq.Len() > 0 {
		top := heap.Pop(pq).(Class)
		totalRatio += float64(top.PassCount) / float64(top.TotalCount)
	}

	return totalRatio / float64(len(classes))
}

func main() {
	classes, extraStudents := [][]int{{1, 2}, {3, 5}, {2, 2}}, 2
	result := maxAverageRatio(classes, extraStudents)
	fmt.Printf("Result: %#v\n", result)
	if result == 0.7833333333333333 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
