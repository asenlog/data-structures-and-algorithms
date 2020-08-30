/**
Given an integer array nums sorted in ascending order, and an integer target.

Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You should search for target in nums and if you found return its index, otherwise return -1.

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:

Input: nums = [1], target = 0
Output: -1



Constraints:
    1 <= nums.length <= 5000
    -10^4 <= nums[i] <= 10^4
    All values of nums are unique.
    nums is guranteed to be rotated at some pivot.
    -10^4 <= target <= 10^4
*/

package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	l := 0
	h := n - 1

	for l <= h {
		mid := (l+h)/2
		if nums[mid] == target {
			return mid
		}

		if l == h {
			return -1
		}

		if nums[mid] <= nums[h] {
			if target > nums[mid] && target <= nums[h] {
				l = mid+1
			} else {
				h = mid-1
			}
		} else {
			if target >= nums[l] && target < nums[mid] {
				h = mid-1
			} else {
				l = mid+1
			}
		}
	}

	return -1
}

func main() {
	in := []int{4,5,6,7,0,1,2}
	fmt.Println(search(in,4))
}
