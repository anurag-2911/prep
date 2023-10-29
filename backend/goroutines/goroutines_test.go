package goroutines

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println()
}

func TestSum(t *testing.T) {

	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr2 := []int{10, 20, 30, 40, 50, 60, 70, 80}

	callSum(arr1, arr2)
}
func BenchMarkSum(b *testing.B) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr2 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	n:=b.N
	fmt.Println("benchmarking times ",n)
	b.ResetTimer()
	for i := 0; i < n; i++ {
		callSum(arr1, arr2)
	}
}
func callSum(arr1 []int, arr2 []int) {
	ch := make(chan int, 2)
	go summe(arr1, ch)
	go summe(arr2, ch)

	<-ch
	<-ch
}

func summe(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	ch <- sum
}
