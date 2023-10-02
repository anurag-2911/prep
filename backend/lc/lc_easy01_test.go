package lc

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	fmt.Println("test")
	res:=countBits(2)
	res=countBits(5)
	fmt.Println(res)
	newFunction()
}


func countBits(n int) []int {
    arr:=make([]int, n+1)
	for i:=0;i<=n;i++ {
		x:=0
		num:=i
		for num>0{
			remain:=num%2
			x=x+remain
			num=num/2
		}
		arr[i]=x
	}
	return arr
}

func newFunction() {
	res03 := numJewelsInStones("aA", "aAAbbbb")
	fmt.Println(res03)
	res02 := strStr("xhello", "hell")
	res02 = strStr("mississippi", "issipi")
	res02 = strStr("sadbutsad", "sad")
	res02 = strStr("leetcode", "leeto")
	res02 = strStr("a", "a")

	fmt.Println(res02)
	test := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(test)
}

func numJewelsInStones(jewels string, stones string) int {

	jewelMap := make(map[rune]int)
	for _, j := range jewels {
		jewelMap[j] = 0
	}
	count:=0
	for _,v:=range stones {
		_,ok:=jewelMap[v]
		if(ok){
			count++
		}
	}
	return count

}
func removeElement(nums []int, val int) int {
	totalElements := len(nums)

	valfound := 0
	arr := make([]int, len(nums))
	for i := 0; i < totalElements; i++ {
		arr[i] = 8888
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			valfound++
		} else {
			arr[count] = nums[i]
			count++
		}
	}
	return totalElements - valfound
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	found := false
	for i := 0; i < len(haystack); i++ {
		if needle[0] == haystack[i] {
			found = true
			currentHayStackCounter := i
			k := 0
			for ; k < len(needle) && currentHayStackCounter < len(haystack); k++ {
				if needle[k] == haystack[currentHayStackCounter] {
					found = true
					currentHayStackCounter++
				} else {
					found = false
					break
				}
			}
			if found == true && k == len(needle) {
				return i
			}
		}

	}
	return -1
}
