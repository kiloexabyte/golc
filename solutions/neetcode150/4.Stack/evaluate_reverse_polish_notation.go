package main

// Leetcode #150
func EvalRPN(tokens []string) int {
	stack := []int{}

	// Helper function to convert string to int without external packages
	stringToInt := func(s string) int {
		num := 0
		sign := 1
		for i, c := range s {
			if c == '-' && i == 0 {
				sign = -1
				continue
			}
			num = num*10 + int(c-'0')
		}
		return num * sign
	}

	for _, token := range tokens {
		switch token {
		case "+":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b+a)
		case "-":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b-a)
		case "*":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b*a)
		case "/":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b/a)
		default:
			stack = append(stack, stringToInt(token))
		}
	}

	return stack[len(stack)-1]
}