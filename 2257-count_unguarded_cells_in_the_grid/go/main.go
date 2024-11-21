package main

import (
	"fmt"
)

type Cell int8

const (
	NotGuarded Cell = 0
	Guard           = 1
	Guarded         = 2
	Wall            = 3
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]Cell, m)
	for r := 0; r < m; r++ {
		grid[r] = make([]Cell, n)
	}

	totalUnguarded := m * n
	for _, wall := range walls {
		grid[wall[0]][wall[1]] = Wall
		totalUnguarded--
	}
	for _, guard := range guards {
		grid[guard[0]][guard[1]] = Guard
		totalUnguarded--
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, guard := range guards {
		for _, direction := range directions {
			r, c := guard[0]+direction[0], guard[1]+direction[1]
			for r >= 0 && r < m && c >= 0 && c < n {
				if grid[r][c] == Wall || grid[r][c] == Guard {
					break
				}
				if grid[r][c] == NotGuarded {
					grid[r][c] = Guarded
					totalUnguarded--
				}
				r, c = r+direction[0], c+direction[1]
			}
		}
	}

	return totalUnguarded
}

func main() {
	m, n, guards, walls := 4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}
	result := countUnguarded(m, n, guards, walls)
	fmt.Printf("Result: %#v\n", result)
	if result == 7 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
