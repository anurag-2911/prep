package dsa

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	fmt.Print()

}

type XNode struct {
	key        int
	value      int
	prev, next *XNode
}

type XLRU struct {
	capacity   int
	cache      map[int]*XNode
	head, tail *XNode
}

func constructor(capacity int) *XLRU {
	head := &XNode{}
	tail := &XNode{}
	head.next = tail
	tail.prev = head

	return &XLRU{
		capacity: capacity,
		cache:make(map[int]*XNode),
		head: head,
		tail: tail,
}

}
