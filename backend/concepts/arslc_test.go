package concepts

import (
	"fmt"
	"testing"
)

func TestFunctions(t *testing.T) {
	q5()
}

func q5() {
	slice := []int{1, 2, 3}
	slice = append(slice, slice...)
	fmt.Println(slice)
}

func q4() {
	slice := make([]int, 3, 5)
	fmt.Println(cap(slice), len(slice))
}
func q1() {
	arr := [3]int{1, 2, 3}
	arr[1] = 4
	fmt.Println(arr)
}
func q2() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)
}
func q3() {
	arr := [...]int{1, 2, 3}
	fmt.Println(len(arr))
}
