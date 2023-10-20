package systemdesign

import (
	"fmt"
	"hash/crc32"
	"sort"
)

func TestConsistentHash(){
	ch := NewConsistentHash()

	ch.AddNode("Node1")
	ch.AddNode("Node2")
	ch.AddNode("Node3")

	fmt.Println(ch.Get("Apple"))
	fmt.Println(ch.Get("Banana"))
	fmt.Println(ch.Get("Cherry"))
	fmt.Println(ch.Get("Date"))
}



type HashRing []uint32

func (hr HashRing) Len() int {
	return len(hr)
}

func (hr HashRing) Less(i, j int) bool {
	return hr[i] < hr[j]
}

func (hr HashRing) Swap(i, j int) {
	hr[i], hr[j] = hr[j], hr[i]
}

type ConsistentHash struct {
	VirtualNodeNum int
	ring           HashRing
	nodes          map[uint32]string
}

func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		VirtualNodeNum: 20,
		ring:           HashRing{},
		nodes:          make(map[uint32]string),
	}
}

func (ch *ConsistentHash) AddNode(node string) {
	for i := 0; i < ch.VirtualNodeNum; i++ {
		virtualNodeName := fmt.Sprintf("%s#%d", node, i)
		hash := ch.hash(virtualNodeName)
		ch.nodes[hash] = node
		ch.ring = append(ch.ring, hash)
	}
	sort.Sort(ch.ring)
}

func (ch *ConsistentHash) Get(key string) string {
	hash := ch.hash(key)
	idx := sort.Search(len(ch.ring), func(i int) bool {
		return ch.ring[i] >= hash
	})

	if idx == len(ch.ring) {
		idx = 0
	}

	return ch.nodes[ch.ring[idx]]
}

func (ch *ConsistentHash) hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}







