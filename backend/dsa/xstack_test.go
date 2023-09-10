package dsa

import (
	"fmt"
	"testing"	
)

func TestStack(t *testing.T) {
	fmt.Print()
	var arr []int
	xstack := &XStack{xstack: arr, count: 0}
	xstack.Push(0)
	xstack.Push(1)
	xstack.Push(6)
	xstack.Push(8)
	for !xstack.isEmpty() {
		fmt.Println(xstack.Pop())
	}

}

type XStack struct {
	xstack []int
	count  int
}

func (this *XStack) Push(val int) {
	this.xstack = append(this.xstack, val)
	this.count++

}
func (this *XStack) Pop() int {
	if this.count == 0 {
		return 0
	}
	ln := len(this.xstack)
	last := this.xstack[ln-1]
	this.xstack = this.xstack[0 : ln-1]
	this.count--
	return last
}
func (this *XStack) isEmpty() bool {
	return this.count == 0
}
