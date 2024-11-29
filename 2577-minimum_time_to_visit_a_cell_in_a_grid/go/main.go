package main

import (
	"container/heap"
	"fmt"
)

type Position struct {
	i int
	j int
}

type State struct {
	pos  Position
	dist int
}

type StateHeap []State

func (sh StateHeap) Len() int           { return len(sh) }
func (sh StateHeap) Less(i, j int) bool { return sh[i].dist < sh[j].dist }
func (sh StateHeap) Swap(i, j int)      { sh[i], sh[j] = sh[j], sh[i] }

func (sh *StateHeap) Push(s any) {
	*sh = append(*sh, s.(State))
}

func (sh *StateHeap) Pop() any {
	old := *sh
	n := len(old)
	x := old[n-1]
	*sh = old[0 : n-1]
	return x
}

func minimumTime(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	startVal := grid[0][0] + 1
	if grid[0][1] > startVal && grid[1][0] > startVal {
		return -1
	}

	directions := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	currentState := State{Position{0, 0}, 0}
	sh := &StateHeap{}
	heap.Init(sh)
	alreadyExplored := make(map[Position]struct{})
	alreadyExplored[currentState.pos] = struct{}{}
	for currentState.pos.i != m-1 || currentState.pos.j != n-1 {
		for _, dir := range directions {
			nextI := currentState.pos.i + dir[0]
			if nextI < 0 || nextI >= m {
				continue
			}
			nextJ := currentState.pos.j + dir[1]
			if nextJ < 0 || nextJ >= n {
				continue
			}
			nextPosition := Position{nextI, nextJ}
			if _, found := alreadyExplored[nextPosition]; found {
				continue
			}
			alreadyExplored[nextPosition] = struct{}{}

			waitTime := max(grid[nextI][nextJ]-(currentState.dist+1), 0)
			if (grid[nextI][nextJ]-(currentState.dist+1))%2 == 1 {
				waitTime++
			}
			nextDistance := currentState.dist + 1 + waitTime
			nextState := State{nextPosition, nextDistance}
			heap.Push(sh, nextState)
		}
		currentState = heap.Pop(sh).(State)
	}

	return currentState.dist
}

func main() {
	grid := [][]int{
		{0, 1, 3, 2},
		{5, 1, 2, 5},
		{4, 3, 8, 6},
	}
	result := minimumTime(grid)
	fmt.Printf("Result: %#v\n", result)
	if result == 7 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
