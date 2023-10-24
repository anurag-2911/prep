package linkedlist

import (
	"fmt"
	"testing"
)

var linkedlist *LinkedList

func init() {
	linkedlist = &LinkedList{}
	linkedlist.Append(100)
	linkedlist.Append(200)
	linkedlist.Append(300)
	linkedlist.Append(400)
	linkedlist.Append(500)
	linkedlist.Append(600)
	linkedlist.Append(700)

}

func TestLinked(t *testing.T) {
	TestLinkedList()
}
func TestFindMiddleElement(t *testing.T) {
	middleElement := FindMiddleElement()
	fmt.Println(middleElement)
}
func FindMiddleElement() *Node {
	slow, fast := linkedlist.Head, linkedlist.Head
	if linkedlist.Head == nil {
		return nil
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
func hasCycle(head *Node) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
func TestLinkedListLoop(t *testing.T) {
	node5 := &Node{5, nil}
	node4 := &Node{4, node5}
	node3 := &Node{3, node4}
	node2 := &Node{2, node3}
	node1 := &Node{1, node2}
	fmt.Println(hasCycle(node1))
	node5.Next = node3
	fmt.Println(hasCycle(node1))

}
func reverseLinkList() *LinkedList {
	current := linkedlist.Head
	var revlist LinkedList
	if current == nil {
		return nil
	}
	arr := make([]int, 0)
	for current != nil {
		arr = append(arr, current.Data)
		current = current.Next
	}
	for i := len(arr) - 1; i >= 0; i-- {
		revlist.Append(arr[i])
	}

	return &revlist
}

func TestRevLinkList(t *testing.T) {
	ll := reverseLinkList()
	ll.Traverse()
}
