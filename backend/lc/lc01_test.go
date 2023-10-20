package lc

import (
	"fmt"
	"strings"

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

func generatePascalTriangle(numOfRows int) [][]int {
	if numOfRows == 0 {
		return [][]int{}
	}
	triangle := make([][]int, numOfRows)

	return triangle
}
func TestPascalTriangle(t *testing.T) {
	numRows := 5
	fmt.Println(generatePascalTriangle(numRows))
	// Expected output: [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]

	numRows = 1
	fmt.Println(generatePascalTriangle(numRows))
	// Expected output: [[1]]
}

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = nums[v]
	}
	return ans
}
func TestBuildArray(t *testing.T) {
	nums := []int{0, 2, 1, 5, 3, 4}
	res := buildArray(nums)
	fmt.Println(res)
}
func removeVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	result := make([]rune, 0)
	for _, v := range s {
		if !strings.ContainsRune(string(vowels), v) {
			result = append(result, v)
		}
	}
	return string(result)
}
func TestRemoveVowels(t *testing.T) {
	s := "leetcodeisacommunityforcoders"
	res := removeVowels(s)
	fmt.Println(res)
}
func defangIPaddr(address string) string {
	result := make([]rune, 0)
	for _, v := range address {
		if v == '.' {
			result = append(result, '[', '.', ']')
		} else {
			result = append(result, v)
		}
	}
	return string(result)
}
func TestDefangIPAd(t *testing.T) {
	address := "1.1.1.1"
	result := defangIPaddr(address)
	fmt.Println(result)
}
