package lc

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	// Starter();
	haystack := "sadbutsad"
	needle := "sad"
	res:=findIndexOfFirstOccurence(haystack,needle)
	fmt.Println(res)
	haystack = "leetcode"
	needle = "leeto"
	res=findIndexOfFirstOccurence(haystack,needle)
	
}
func findIndexOfFirstOccurence(heystack string,needle string)(int){
	z:=0
	for i := 0; i < len(heystack); i++ {
		if(needle[i]==heystack[i]){
			allmatched:=false
			for k := i; i < len(needle); i++ {
				if(needle[k]==heystack[k]){
					allmatched=true
				}else{
					allmatched=false
					z++
					break
				}
			}
			if(allmatched){
				return z
			}
		}
	}
	return -1
}



