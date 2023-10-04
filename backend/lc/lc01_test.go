package lc

import (
	"fmt"

	"sort"
	"testing"
)

func TestAll(t *testing.T) {
	fmt.Println("tests")
	r := repeatedCharacter("abccbaacz")
	fmt.Println(r)
	nums := []int{8, 1, 2, 2, 3}
	res := smallerNumbersThanCurrent(nums)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	sum := 0
	current := head
	entries := make([]int, 0)
	for current.Next == nil {
		entries = append(entries, current.Val)
		current = current.Next
	}
	//  for i:=len(entries)-1;i>=0;i--{
	// 	sum=sum+ int(math.Pow(2,i))
	//  }
	return sum
}

func repeatedCharacter(s string) byte {
	var b byte
	charmap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		_, exists := charmap[s[i]]
		if exists {

			return s[i]
		} else {
			charmap[s[i]] = 0
		}
	}
	return b
}
func smallerNumbersThanCurrent(nums []int) []int {
	sort.Ints(nums)
	return nums
}
