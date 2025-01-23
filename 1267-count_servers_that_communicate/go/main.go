package main

import (
	"fmt"
)

func countServers(grid [][]int) int {
	computersOnRows := make([]int, len(grid))
	computersOnColumns := make([]int, len(grid[0]))
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 {
				computersOnRows[y]++
				computersOnColumns[x]++
			}
		}
	}
	var connectedServers int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 && (computersOnRows[y] > 1 || computersOnColumns[x] > 1) {
				connectedServers++
			}
		}
	}
	return connectedServers
}

func main() {
	grid := [][]int{
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	result := countServers(grid)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
