package progs

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	LRU()
}

func LRU() {
	lru := Construct(20)
	fmt.Println(lru.capacity)
}

type Node struct {
	Key   int
	Value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func (this *LRUCache) Get(key int) int {
	if node, found := this.cache[key]; found {
		this.moveToFront(node)
		return node.Value
	}

	return -1
}
func (this *LRUCache) Put(key int, value int) {
	if node, found := this.cache[key]; found {
		node.Value = value
		this.moveToFront(node)
	} else {
		node = &Node{Key: key, Value: value}
		if len(this.cache) >= this.capacity {
			delete(this.cache, this.tail.prev.Key)
			this.removeNode(this.tail.prev)
		}
		this.cache[key] = node
		this.addToFront(node)
	}

}
func (this *LRUCache) addToFront(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node

}
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addToFront(node)
}
func Construct(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}

}
