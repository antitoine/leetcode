package main

import (
	"fmt"
	"math"
	"sort"
)

type BinaryIndexedTree struct {
	size int
	tree []int
}

func NewBinaryIndexedTree(size int) *BinaryIndexedTree {
	tree := make([]int, size+1)
	for i := range tree {
		tree[i] = math.MaxInt32
	}
	return &BinaryIndexedTree{
		size: size,
		tree: tree,
	}
}

func (bit *BinaryIndexedTree) Update(index, value int) {
	for index <= bit.size {
		bit.tree[index] = min(bit.tree[index], value)
		index += index & -index
	}
}

func (bit *BinaryIndexedTree) Query(index int) int {
	minValue := math.MaxInt32
	for index > 0 {
		minValue = min(minValue, bit.tree[index])
		index -= index & -index
	}
	if minValue == math.MaxInt32 {
		return -1
	}
	return minValue
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	numBuildings := len(heights)
	numQueries := len(queries)

	for i := range queries {
		if queries[i][0] > queries[i][1] {
			queries[i][0], queries[i][1] = queries[i][1], queries[i][0]
		}
	}

	sortedQueryIndices := make([]int, numQueries)
	for i := range sortedQueryIndices {
		sortedQueryIndices[i] = i
	}
	sort.Slice(sortedQueryIndices, func(i, j int) bool {
		return queries[sortedQueryIndices[j]][1] < queries[sortedQueryIndices[i]][1]
	})

	answers := make([]int, numQueries)
	for i := range answers {
		answers[i] = -1
	}
	bit := NewBinaryIndexedTree(numBuildings)

	uniqueSortedHeights := append([]int{}, heights...)
	sort.Ints(uniqueSortedHeights)
	uniqueSortedHeights = unique(uniqueSortedHeights)

	lastProcessedBuilding := numBuildings - 1

	for _, queryIndex := range sortedQueryIndices {
		leftBound, rightBound := queries[queryIndex][0], queries[queryIndex][1]

		for lastProcessedBuilding > rightBound {
			heightIndex := numBuildings - lowerBound(uniqueSortedHeights, heights[lastProcessedBuilding]) + 1
			bit.Update(heightIndex, lastProcessedBuilding)
			lastProcessedBuilding--
		}

		if leftBound == rightBound || heights[leftBound] < heights[rightBound] {
			answers[queryIndex] = rightBound
		} else {
			heightIndex := numBuildings - lowerBound(uniqueSortedHeights, heights[leftBound])
			answers[queryIndex] = bit.Query(heightIndex)
		}
	}

	return answers
}

func lowerBound(arr []int, target int) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
}

func unique(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	result := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result
}

func main() {
	heights := []int{6, 4, 8, 5, 2, 7}
	queries := [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}
	result := leftmostBuildingQueries(heights, queries)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%#v", result) == "[]int{2, 5, -1, 5, 2}" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
