package linkedlist

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func (this *LinkedList) Append(val int) {
	newNode := &Node{Data: val}
	if this.Head == nil {
		this.Head = newNode
	} else {
		current := this.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		this.Size++
	}
}
func (this *LinkedList) Remove(val int) {
	if this.Head == nil {
		return
	}
	if this.Head.Data == val { // first node case
		this.Size--
		this.Head = this.Head.Next
		return
	}
	current := this.Head
	for current.Next != nil && current.Next.Data != val {
		current = current.Next
	}
	if current.Next == nil {
		return // no node with such value found
	}

	current.Next = current.Next.Next
	this.Size--

}

func (this *LinkedList) Contains(data int) bool {
	current := this.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

func TestLinkedList() {
	ll := &LinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Remove(4)
}
