package main

import (
	"math/big"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// leetcode #2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1, n2 string

	for l1 != nil {
		if l1 != nil {
			n1 = strconv.Itoa(l1.Val) + n1
			l1 = l1.Next
		}
	}

	// Traverse the second linked list and construct the number as a string
	for l2 != nil {
		n2 = strconv.Itoa(l2.Val) + n2 // Prepend the digit
		l2 = l2.Next
	}

	// Parse the strings into big integers
	num1 := new(big.Int)
	num1.SetString(n1, 10)

	num2 := new(big.Int)
	num2.SetString(n2, 10)
	
	// Add the two big integers
	num3 := new(big.Int)
	num3.Add(num1, num2)

	// Convert the result back to a string
	str3 := num3.String()

	// Construct the result linked list
	dummy := &ListNode{}
	current := dummy

	for i := len(str3) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(string(str3[i]))
		current.Val = val

		if i != 0 { // Create a new node for the next digit
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return dummy
}