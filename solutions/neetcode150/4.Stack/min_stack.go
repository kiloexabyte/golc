package main

type MinStack struct {
    arr   []int
    minArr []int
}

// Constructor initializes and returns a new MinStack
func Constructor() MinStack {
    return MinStack{arr: []int{}, minArr: []int{}}
}

// Push adds an element to the stack
func (minStack *MinStack) Push(val int) {
    // Push the value onto the main stack
    minStack.arr = append(minStack.arr, val)
    
    // Determine the new minimum and push it onto the min stack
    if len(minStack.minArr) == 0 {
        minStack.minArr = append(minStack.minArr, val)
    } else {
        minStack.minArr = append(minStack.minArr, MinInt(val, minStack.minArr[len(minStack.minArr)-1]))
    }
}

// Pop removes the top element from the stack
func (minStack *MinStack) Pop() {
    if len(minStack.arr) == 0 {
        return // Handle underflow as needed
    }
    minStack.arr = minStack.arr[:len(minStack.arr)-1]
    minStack.minArr = minStack.minArr[:len(minStack.minArr)-1] // Corrected length here
}

// Top returns the top element of the stack
func (minStack *MinStack) Top() int {
    if len(minStack.arr) == 0 {
        return -1 // Handle empty stack case
    }
    return minStack.arr[len(minStack.arr)-1]
}

// GetMin returns the minimum element from the stack
func (minStack *MinStack) GetMin() int {
    if len(minStack.minArr) == 0 {
        return -1 // Handle empty stack case
    }
    return minStack.minArr[len(minStack.minArr)-1]
}

// MinInt returns the minimum of two integers
func MinInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

