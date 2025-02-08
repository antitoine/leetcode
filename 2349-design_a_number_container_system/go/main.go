package main

import (
	"fmt"
	"slices"
)

type NumberContainers struct {
	byIdx map[int]int
	byVal map[int][]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		byIdx: make(map[int]int),
		byVal: make(map[int][]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if val, ok := this.byIdx[index]; ok {
		if val == number {
			return
		}
		idx, _ := slices.BinarySearch(this.byVal[val], index)
		this.byVal[val] = append(this.byVal[val][:idx], this.byVal[val][idx+1:]...)
	}
	this.byIdx[index] = number
	if _, ok := this.byVal[number]; !ok {
		this.byVal[number] = []int{}
	}
	idx, _ := slices.BinarySearch(this.byVal[number], index)
	this.byVal[number] = append(this.byVal[number][:idx], append([]int{index}, this.byVal[number][idx:]...)...)
}

func (this *NumberContainers) Find(number int) int {
	if indexes, ok := this.byVal[number]; ok && len(indexes) > 0 {
		return indexes[0]
	}
	return -1
}

func main() {
	obj := Constructor()
	if val := obj.Find(10); val != -1 {
		fmt.Printf("1. Test failed, expected -1, got %d", val)
		return
	}
	obj.Change(2, 10)
	obj.Change(1, 10)
	obj.Change(3, 10)
	obj.Change(5, 10)
	if val := obj.Find(10); val != 1 {
		fmt.Printf("2. Test failed, expected 1, got %d", val)
		return
	}
	obj.Change(1, 20)
	if val := obj.Find(10); val != 2 {
		fmt.Printf("3. Test failed, expected 2, got %d", val)
		return
	}
	fmt.Println("Passed")
}
