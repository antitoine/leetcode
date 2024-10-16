package main

import (
	"container/heap"
	"fmt"
)

type Char struct {
	c    rune
	left int
}

type CharHeap []Char

func (h CharHeap) Len() int           { return len(h) }
func (h CharHeap) Less(i, j int) bool { return h[i].left > h[j].left }
func (h CharHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CharHeap) Push(x any) {
	*h = append(*h, x.(Char))
}

func (h *CharHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	chars := &CharHeap{}
	heap.Init(chars)
	if a > 0 {
		heap.Push(chars, Char{'a', a})
	}
	if b > 0 {
		heap.Push(chars, Char{'b', b})
	}
	if c > 0 {
		heap.Push(chars, Char{'c', c})
	}
	var result string
	var lastChar rune
	lastCharOcc := 0
	for chars.Len() > 0 {
		var invalidChars []Char
		char := heap.Pop(chars).(Char)
		for chars.Len() > 0 && char.c == lastChar && lastCharOcc >= 2 {
			invalidChars = append(invalidChars, char)
			char = heap.Pop(chars).(Char)
		}
		if char.c != lastChar || lastCharOcc < 2 {
			result += string(char.c)
			char.left--
			if char.left > 0 {
				heap.Push(chars, char)
			}
			if char.c != lastChar {
				lastChar = char.c
				lastCharOcc = 1
			} else {
				lastCharOcc++
			}
		} else {
			return result
		}
		for _, invalidChar := range invalidChars {
			heap.Push(chars, invalidChar)
		}
	}
	return result
}

func main() {
	a, b, c := 1, 1, 7
	result := longestDiverseString(a, b, c)
	fmt.Printf("Result: %#v\n", result)
	if result == "ccaccbcc" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
