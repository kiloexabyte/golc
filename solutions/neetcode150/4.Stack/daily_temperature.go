package main

// Define a struct to represent key-value pairs (temperature, index)
type KeyValuePair struct {
	Temperature int
	Index       int
}

// Leetcode #739
func dailyTemperatures(temperatures []int) []int {
	stack := []KeyValuePair{} // Stack to hold key-value pairs
	res := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		// While the stack is not empty and the current temperature is greater than the temperature at the top of the stack
		for len(stack) > 0 && temperatures[i] > stack[len(stack)-1].Temperature {
			// Pop the top element from the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Remove the last element

			// Calculate the number of days until a warmer temperature
			res[top.Index] = i - top.Index
		}

		// Push the current temperature and its index onto the stack
		stack = append(stack, KeyValuePair{Temperature: temperatures[i], Index: i})
	}

	return res
}
