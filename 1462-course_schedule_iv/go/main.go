package main

import (
	"container/list"
	"fmt"
)

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	preReqMatrix := make([][]bool, numCourses)
	for i := range preReqMatrix {
		preReqMatrix[i] = make([]bool, numCourses)
	}
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, preReq := range prerequisites {
		course, nextCourse := preReq[0], preReq[1]
		graph[course] = append(graph[course], nextCourse)
		inDegree[nextCourse]++
	}
	queue := list.New()
	for i, count := range inDegree {
		if count == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		currentCourse := queue.Remove(queue.Front()).(int)
		for _, nextCourse := range graph[currentCourse] {
			preReqMatrix[currentCourse][nextCourse] = true
			for course := 0; course < numCourses; course++ {
				if preReqMatrix[course][currentCourse] {
					preReqMatrix[course][nextCourse] = true
				}
			}
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue.PushBack(nextCourse)
			}
		}
	}
	result := make([]bool, len(queries))
	for i, query := range queries {
		startCourse, endCourse := query[0], query[1]
		result[i] = preReqMatrix[startCourse][endCourse]
	}
	return result
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	queries := [][]int{{0, 1}, {1, 0}}
	result := checkIfPrerequisite(numCourses, prerequisites, queries)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[false true]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
