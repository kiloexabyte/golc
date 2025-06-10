package main

import "strings"

var (
	res   []string
	stack []string
	N     int
)

// leetcode #22
func generateParenthesis(n int) []string {
	res = []string{}
	stack = []string{}

	N = n
	backtrack(0,0)
	return res
}

// backtrack performs the recursive generation of parentheses.
func backtrack(openN, closeN int) {
	// Base case: if we've used all open and close parentheses
	if openN == closeN && closeN == N {
		// Create a valid combination from the stack and append it to the result
		res = append(res, strings.Join(stack, ""))
		return
	}

	// Add an open parenthesis if possible
	if openN < N {
		stack = append(stack, "(")   // Push to stack
		backtrack(openN+1, closeN)   // Recur with one more open parenthesis
		stack = stack[:len(stack)-1] // Pop from stack
	}

	// Add a close parenthesis if it doesn't exceed the open count
	if closeN < openN {
		stack = append(stack, ")")   // Push to stack
		backtrack(openN, closeN+1)   // Recur with one more close parenthesis
		stack = stack[:len(stack)-1] // Pop from stack
	}
}