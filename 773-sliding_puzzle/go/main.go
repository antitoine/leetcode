package main

import (
	"fmt"
)

type Board struct {
	l1 uint8
	l2 uint8
	l3 uint8
	l4 uint8
	l5 uint8
	l0 uint8
}

func (b Board) Valid() bool {
	return b.l1 == 1 && b.l2 == 2 && b.l3 == 3 && b.l4 == 4 && b.l5 == 5 && b.l0 == 0
}

func (b Board) Next() []Board {
	if b.l1 == 0 {
		return []Board{
			Board{b.l4, b.l2, b.l3, b.l1, b.l5, b.l0},
			Board{b.l2, b.l1, b.l3, b.l4, b.l5, b.l0},
		}
	}
	if b.l2 == 0 {
		return []Board{
			Board{b.l2, b.l1, b.l3, b.l4, b.l5, b.l0},
			Board{b.l1, b.l3, b.l2, b.l4, b.l5, b.l0},
			Board{b.l1, b.l5, b.l3, b.l4, b.l2, b.l0},
		}
	}
	if b.l3 == 0 {
		return []Board{
			Board{b.l1, b.l3, b.l2, b.l4, b.l5, b.l0},
			Board{b.l1, b.l2, b.l0, b.l4, b.l5, b.l3},
		}
	}
	if b.l4 == 0 {
		return []Board{
			Board{b.l4, b.l2, b.l3, b.l1, b.l5, b.l0},
			Board{b.l1, b.l2, b.l3, b.l5, b.l4, b.l0},
		}
	}
	if b.l5 == 0 {
		return []Board{
			Board{b.l1, b.l2, b.l3, b.l5, b.l4, b.l0},
			Board{b.l1, b.l2, b.l3, b.l4, b.l0, b.l5},
			Board{b.l1, b.l5, b.l3, b.l4, b.l2, b.l0},
		}
	}
	if b.l0 == 0 {
		return []Board{
			Board{b.l1, b.l2, b.l0, b.l4, b.l5, b.l3},
			Board{b.l1, b.l2, b.l3, b.l4, b.l0, b.l5},
		}
	}
	return []Board{}
}

func slidingPuzzle(board [][]int) int {
	b := Board{
		l1: uint8(board[0][0]),
		l2: uint8(board[0][1]),
		l3: uint8(board[0][2]),
		l4: uint8(board[1][0]),
		l5: uint8(board[1][1]),
		l0: uint8(board[1][2]),
	}
	if b.Valid() {
		return 0
	}
	existingBoards := make(map[Board]struct{})
	existingBoards[b] = struct{}{}
	nextBoards := b.Next()
	for i := 1; len(nextBoards) > 0; i++ {
		var newNextBoards []Board
		for _, nextBoard := range nextBoards {
			if nextBoard.Valid() {
				return i
			}
			if _, found := existingBoards[nextBoard]; found {
				continue
			}
			existingBoards[nextBoard] = struct{}{}
			newNextBoards = append(newNextBoards, nextBoard.Next()...)
		}
		nextBoards = newNextBoards
	}
	return -1
}

func main() {
	board := [][]int{{4, 1, 2}, {5, 0, 3}}
	result := slidingPuzzle(board)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
