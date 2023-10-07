package progs

import (
	"fmt"
	"testing"
)

func TestProgs(t *testing.T) {
	fmt.Println("test")
	si := searchInsert([]int{1, 3, 5, 6, 11, 15, 21}, 5)
	si = searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Println(si)
}
// searchInsert searches for the given target in the sorted slice nums.
//
// If the target is found, it returns its index.
// If the target is not found, it returns the index at which the target can be 
// inserted in order to maintain the slice's sorted order.
//
// This function uses a binary search algorithm, achieving O(log n) runtime complexity.
//
// Parameters:
// - nums: A sorted slice of distinct integers.
// - target: The integer value to search for.
//
// Returns:
// - The index of the target if found, or the index where the target could be inserted.
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	// While there are elements between left and right inclusive
	for left <= right {
		 // Find the middle index
		mid := (right+left) / 2
		// Check if the target is found
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			// Target might be on the left side
			right = mid-1

		} else {
			// Target might be on the right side
			
			left = mid+1

		}
	}
	// If not found, left now points to the index where the target can be inserted
	return left
}
