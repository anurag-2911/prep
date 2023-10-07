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

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right+left) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			right = mid-1

		} else {
			left = mid+1

		}
	}

	return left
}
