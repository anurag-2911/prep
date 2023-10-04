package lc

import (
	"fmt"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	ln := lengthOfLastWord("Hello World") //5
	ln=lengthOfLastWord("   fly me   to   the moon  ") //4
	ln=lengthOfLastWord("luffy is still joyboy")//6
	fmt.Println(ln)
	// Starter();
	haystack := "sadbutsad"
	needle := "sad"
	res := findIndexOfFirstOccurence(haystack, needle)
	fmt.Println(res)
	haystack = "leetcode"
	needle = "leeto"
	res = findIndexOfFirstOccurence(haystack, needle)

}
func findIndexOfFirstOccurence(heystack string, needle string) int {
	z := 0
	for i := 0; i < len(heystack); i++ {
		if needle[i] == heystack[i] {
			allmatched := false
			for k := i; i < len(needle); i++ {
				if needle[k] == heystack[k] {
					allmatched = true
				} else {
					allmatched = false
					z++
					break
				}
			}
			if allmatched {
				return z
			}
		}
	}
	return -1
}


func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	lastword := 0
	s=strings.TrimSpace(s)
	for _, v := range s {
		
		if v != ' ' {
			lastword++
		}else{
			lastword=0
		}
		
	}
	return lastword
}
func plusOne(digits []int) []int {
	rarray:=make([]int,len(digits))

	

	return rarray
}


