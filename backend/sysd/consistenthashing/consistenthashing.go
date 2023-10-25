package consistenthashing

import (
	"backend/sysd/kvstore"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
)

const virtualnodes = 10

type Node struct {
	ID    string
	Store kvstore.Store
}
type ConsistentHashingRing struct {
	hashKeys  []uint64
	nodes     map[uint64]Node
	realnodes map[uint64]Node
}

func NewConsistentHashRing() *ConsistentHashingRing {
	return &ConsistentHashingRing{
		hashKeys:  make([]uint64, 0),
		nodes:     make(map[uint64]Node),
		realnodes: make(map[uint64]Node),
	}
}
func (this *ConsistentHashingRing) AddNode(node Node) {
	for i := 0; i < virtualnodes; i++ {
		virtualID := fmt.Sprintf("%s#%d", node.ID, i)
		hashkey := this.hash(virtualID)
		this.nodes[hashkey] = node
		this.hashKeys = append(this.hashKeys, hashkey)

	}
	this.realnodes[this.hash(node.ID)] = node
	sort.Slice(this.hashKeys, func(i int, j int) bool {
		return this.hashKeys[i] < this.hashKeys[j]
	})
}

func (this *ConsistentHashingRing) GetNodeForKey(key string) Node {
	hashkey := this.hash(key)
	index := sort.Search(len(this.hashKeys), func(i int) bool {
		return this.hashKeys[i] >= hashkey
	})
	// Wrap around if needed
	if index == len(this.hashKeys) {
		index = 0
	}
	return this.nodes[this.hashKeys[index]]
}

func (this *ConsistentHashingRing) hash(key string) uint64 {
	hashedData := sha256.Sum256([]byte(key))
	return binary.BigEndian.Uint64(hashedData[:8])
}
