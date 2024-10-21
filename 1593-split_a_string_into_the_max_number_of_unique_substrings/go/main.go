package main

import (
	"fmt"
	"strings"
)

type Set struct {
	items map[string]struct{}
}

func (s *Set) Has(item string) bool {
	_, found := s.items[item]
	return found
}

func (s *Set) Add(item string) {
	s.items[item] = struct{}{}
}

func (s *Set) Clone() *Set {
	newSet := NewSet()
	for item := range s.items {
		newSet.Add(item)
	}
	return newSet
}

func (s *Set) CloneWith(item string) *Set {
	newSet := s.Clone()
	newSet.Add(item)
	return newSet
}

func (s *Set) String() string {
	keys := make([]string, len(s.items))
	i := 0
	for key := range s.items {
		keys[i] = key
		i++
	}
	return "[" + strings.Join(keys, ",") + "]"
}

func NewSet() *Set {
	return &Set{make(map[string]struct{})}
}

func maxUniqueSplitRec(s string, used *Set) int {
	if len(s) == 0 {
		return 0
	}
	maxSplits := -1
	for i := 1; i <= len(s); i++ {
		substr := s[0:i]
		if !used.Has(substr) {
			nextUsed := used.CloneWith(substr)
			next := maxUniqueSplitRec(s[i:], nextUsed)
			if next != -1 {
				maxSplits = max(maxSplits, 1+next)
			}
		}
	}
	return maxSplits
}

func maxUniqueSplit(s string) int {
	return maxUniqueSplitRec(s, NewSet())
}

func main() {
	input := "ababccc"
	result := maxUniqueSplit(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 5 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
