package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Gifts []int

func (h Gifts) Len() int           { return len(h) }
func (h Gifts) Less(i, j int) bool { return h[i] > h[j] }
func (h Gifts) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Gifts) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Gifts) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	g := Gifts(gifts)
	piles := &g
	heap.Init(piles)

	for i := 0; i < k; i++ {
		selectedPile := heap.Pop(piles).(int)
		newPile := int(math.Floor(math.Sqrt(float64(selectedPile))))
		heap.Push(piles, newPile)
	}

	var result int64
	for _, pile := range *piles {
		result += int64(pile)
	}
	return result
}

func main() {
	gifts := []int{25, 64, 9, 4, 100}
	k := 4
	result := pickGifts(gifts, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 29 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
