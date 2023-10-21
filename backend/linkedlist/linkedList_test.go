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
