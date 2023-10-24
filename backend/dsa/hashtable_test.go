package dsa

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println("tests")
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}

// A pair (i, j) is called identical if nums[i] == nums[j] and i < j.
func numIdenticalPairs(nums []int) int {

	return 0
}
