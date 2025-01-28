package main

import (
	"container/list"
	"fmt"
	"image"
)

func findMaxFish(grid [][]int) int {
	visited := make(map[image.Point]bool)
	direction := []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	space := image.Rect(0, 0, len(grid[0]), len(grid))
	computeFish := func(x, y int) int {
		startPt := image.Pt(x, y)
		if visited[startPt] {
			return 0
		}
		queue := list.New()
		queue.PushBack(startPt)
		fish := 0
		for queue.Len() > 0 {
			currentPt := queue.Remove(queue.Front()).(image.Point)
			if visited[currentPt] {
				continue
			}
			visited[currentPt] = true
			fish += grid[currentPt.Y][currentPt.X]
			for _, dir := range direction {
				nextPt := currentPt.Add(dir)
				if nextPt.In(space) && grid[nextPt.Y][nextPt.X] > 0 {
					queue.PushBack(nextPt)
				}
			}
		}
		return fish
	}

	maxFish := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				continue
			}
			possibleFish := computeFish(x, y)
			if possibleFish > maxFish {
				maxFish = possibleFish
			}
		}
	}

	return maxFish
}

func main() {
	grid := [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}
	result := findMaxFish(grid)
	fmt.Printf("Result: %#v\n", result)
	if result == 7 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
