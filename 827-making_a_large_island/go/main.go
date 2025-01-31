package main

import (
	"fmt"
	"image"
)

func largestIsland(grid [][]int) int {
	n := len(grid)
	directions := [4]image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	space := image.Rect(0, 0, n, n)

	var maxIslandSize int
	var islands []int
	visited := make(map[image.Point]int)

	dfs := func(start image.Point, nextIslandIdx int) int {
		queue := []image.Point{start}
		var size int
		for len(queue) > 0 {
			currentPt := queue[0]
			queue = queue[1:]
			if _, found := visited[currentPt]; found {
				continue
			}
			size++
			visited[currentPt] = nextIslandIdx
			for _, dir := range directions {
				nextPt := currentPt.Add(dir)
				if nextPt.In(space) && grid[nextPt.X][nextPt.Y] == 1 {
					queue = append(queue, nextPt)
				}
			}
		}
		if size > maxIslandSize {
			maxIslandSize = size
		}
		return size
	}

	var possibilities []image.Point

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			pt := image.Pt(x, y)
			if grid[x][y] == 0 {
				possibilities = append(possibilities, pt)
			} else if _, found := visited[pt]; !found {
				size := dfs(pt, len(islands))
				islands = append(islands, size)
			}
		}
	}

	for _, pt := range possibilities {
		islandsIdx := make(map[int]struct{})
		for _, dir := range directions {
			nextPt := pt.Add(dir)
			if nextPt.In(space) && grid[nextPt.X][nextPt.Y] == 1 {
				islandsIdx[visited[nextPt]] = struct{}{}
			}
		}
		size := 1
		for idx := range islandsIdx {
			size += islands[idx]
		}
		if size > maxIslandSize {
			maxIslandSize = size
		}
	}

	return maxIslandSize
}

func main() {
	grid := [][]int{{1, 0}, {0, 1}}
	result := largestIsland(grid)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
