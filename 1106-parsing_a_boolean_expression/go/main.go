package main

import (
	"fmt"
	"regexp"
)

func reduceNot(expression string) string {
	if expression[2] == 'f' {
		return "t"
	} else {
		return "f"
	}
}

func reduceAnd(expression string) string {
	for i := 2; i < len(expression)-1; i += 2 {
		if expression[i] == 'f' {
			return "f"
		}
	}
	return "t"
}

func reduceOr(expression string) string {
	for i := 2; i < len(expression)-1; i += 2 {
		if expression[i] == 't' {
			return "t"
		}
	}
	return "f"
}

func reduceExpression(expression string) string {
	switch expression[0] {
	case '!':
		return reduceNot(expression)
	case '&':
		return reduceAnd(expression)
	case '|':
		return reduceOr(expression)
	}
	panic("Invalid char in reduceExpression: " + expression)
	return ""
}

func parseBoolExpr(expression string) bool {
	groupPattern := regexp.MustCompile(`([|&!]\((?:[^)(]+)+\))`)
	matchIndexes := groupPattern.FindStringIndex(expression)
	for matchIndexes != nil {
		expression = expression[0:matchIndexes[0]] + reduceExpression(expression[matchIndexes[0]:matchIndexes[1]]) + expression[matchIndexes[1]:]
		matchIndexes = groupPattern.FindStringIndex(expression)
	}
	return expression[0] == 't'
}

func main() {
	input := "&(|(f))"
	result := parseBoolExpr(input)
	fmt.Printf("Result: %#v\n", result)
	if result == false {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
