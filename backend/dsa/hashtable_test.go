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
	pairs:=make(map[int]int, 0)
	count:=0
	for _,num:=range nums{
		pairs[num]++
	}
	for _, freq:=range pairs {
		count+=freq*(freq-1)/2
	}
	return count
}
func TestSmallerNumThanCurr(t *testing.T){
	nums := []int{8,1,2,2,3}
	fmt.Println(smallerNumbersThanCurrent(nums))
	fmt.Println(smallerNumbersThanCurrent([]int{6,5,4,8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7,7,7,7}))

}
func smallerNumbersThanCurrent(nums []int) []int {
	result:=make([]int, 0)
	for i:=0;i<len(nums);i++{
		count:=0
		for k := 0; k < len(nums); k++ {
			if(i==k){
				continue
			}
			if(nums[i]>nums[k]){
				count++
			}
		}
		result = append(result, count)
	}
    return result
}