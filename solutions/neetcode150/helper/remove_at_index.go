package main

func removeAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		// Return the original slice if the index is out of bounds
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}