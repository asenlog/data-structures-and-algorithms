package main

import "fmt"

// O(nlogn) time | O(1) space
// Binary Search is always used in an ordered array.
// Performs best when we have large ordered datasets.
func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high{
		// find the middle element
		median := (low + high) / 2

		// if the middle element is lower
		// than the target we are looking for
		// discard the left side
		// else discard the right side
		if haystack[median] < needle {
			low = median + 1
		}else{
			high = median - 1
		}
	}

	// if the length of the array equals to low value
	// or our target in not in array[low] value
	// the element does not exist in the array
	if low == len(haystack) || haystack[low] != needle {
		return false
	}


	return true
}

func main(){
	items := []int{5,7, 9, 21, 35, 47, 69, 84, 92}
	fmt.Println(binarySearch(47, items))
}