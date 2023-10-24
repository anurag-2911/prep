package sysd

import (
	"fmt"
)

func TestKV() {
	fmt.Println("")
	hashTable := &XHashTable{}
	hashTable.Put("name", "john")
	hashTable.Put("age", "25")
	hashTable.Put("eag", "draupadi")
	hashTable.Put("gae", "murmu")
	hashTable.Put("eag", "kovind")
	hashTable.Put("anme", "ramnath")
	hashTable.Put("amne", "james")
	hashTable.Put("mena", "xalxo")

	val, exists := hashTable.Get("anme")
	if exists {
		fmt.Println(val)
		hashTable.Delete("anme")
	}

}

type KVP struct {
	Key   string
	Value string
}

const BucketCount = 10

type XHashTable struct {
	Buckets [BucketCount][]KVP
}

func (this *XHashTable) hash(key string) int {
	sum := 0
	for _, val := range key {
		sum += int(val)
	}
	return sum % BucketCount
}

// handle collisions by chaining
func (this *XHashTable) Put(key string, value string) {
	index := this.hash(key)
	bucket := &this.Buckets[index]

	for _, bk := range *bucket {
		if bk.Key == key {
			//key already exists update the new value
			bk.Value = value
			return
		}
	}
	*bucket = append(*bucket, KVP{Key: key, Value: value})
}

func (this *XHashTable) Get(key string) (string, bool) {
	index := this.hash(key)
	kvp := this.Buckets[index]

	for _, v := range kvp {
		if v.Key == key {
			return v.Value, true
		}
	}
	return "", false

}
func (this *XHashTable) Delete(key string) {
	index := this.hash(key)
	bucket := &this.Buckets[index]
	for i, k := range *bucket {
		if k.Key == key {
			*bucket = append((*bucket)[:i], (*bucket)[i+1:]...)
			return
		}
	}

}
