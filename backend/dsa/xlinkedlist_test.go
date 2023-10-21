package dsa

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	fmt.Println("tests")
	node1:=&ZNode{data:1}
	linkedList:=&ZLinkedList{head:node1}
	linkedList.Insert(ZNode{data:2})
}

type ZNode struct {
	data int
	next *ZNode
}

type ZLinkedList struct {
	head *ZNode
}

func (this *ZLinkedList) Insert(node ZNode) {
	current := this.head
	for current.next != nil {
		current = current.next
	}
	current.next=&node


}

func (this *ZLinkedList) Remove(node ZNode) {

}
func (this *ZLinkedList) Traverse() {

}

// // Given the head of a singly linked list, reverse the list, and return the reversed list.
// func reverseList(head *ListNode) *ListNode {

// }
