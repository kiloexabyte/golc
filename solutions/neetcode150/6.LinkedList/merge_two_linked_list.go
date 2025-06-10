package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Leetcode #21
 func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    // Create a dummy node to serve as the start of the merged list
    dummy := &ListNode{}
    tail := dummy

    // Traverse both lists until one is exhausted
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }

    // Attach the remaining elements of the non-exhausted list
    if list1 != nil {
        tail.Next = list1
    } else if list2 != nil {
        tail.Next = list2
    }

    return dummy.Next
}