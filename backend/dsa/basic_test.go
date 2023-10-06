package dsa

import (
	"fmt"
	"testing"
)

func TestDSA(t *testing.T) {
	fmt.Println("tests")
	stack := &PStack{data: make([]int, 0), count: 0}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	res := stack.Pop()
	res = stack.Pop()
	res = stack.Pop()
	res = stack.Pop()
	fmt.Println(res)
}

type PStack struct {
	data  []int
	count int
}

func (this *PStack) Push(n int) {
	this.data = append(this.data, n)
	this.count++
}
func (this *PStack) Pop() int {
	if !this.IsEmpty() {
		element := this.data[this.count-1]
		this.count--
		this.data = this.data[0:this.count]
		return element
	} else {
		return -1
	}
}
func (this *PStack) IsEmpty() bool {
	return this.count == 0
}

type PQueue struct {
	data  []int
	front int
	rear  int
}

type MyHashMap struct {
	keys   []int
	values []int
}

const bucketCount =1000

type zode struct{
	key int
	value int
	next *zode
}
type JHashMap struct{
	buckets []*zode
}

func Constructor()JHashMap{
	return JHashMap{buckets: make([]*zode, bucketCount)}
}

func(this *JHashMap)Put(key int, value int){
	index:=key % bucketCount

	if(this.buckets[index]==nil){
		this.buckets[index]=&zode{key:key,value:value}
		return
	}
	current:=this.buckets[index]
	for current!=nil{
		
		current=current.next
	}
}
