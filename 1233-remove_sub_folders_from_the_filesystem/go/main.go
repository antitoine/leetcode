package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(paths []string) []string {
	sort.Strings(paths)
	newPaths := []string{paths[0]}
	for i := 1; i < len(paths); i++ {
		if !strings.HasPrefix(paths[i], newPaths[len(newPaths)-1]+"/") {
			newPaths = append(newPaths, paths[i])
		}
	}
	return newPaths
}

func main() {
	input := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	result := removeSubfolders(input)
	fmt.Printf("Result: %#v\n", result)
	if result[0] == "/a" && result[1] == "/c/d" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
