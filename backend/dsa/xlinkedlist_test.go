package dsa

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	fmt.Print()

}

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
	tail *Node
}

func (this *LinkedList) Append(data int) {
	node := &Node{data: data}
	if this.head == nil { //first node
		node.next = nil
		this.head = node
		this.tail = node
	}
	current := this.head
	for current.next != nil {
		current = current.next
	}
}
