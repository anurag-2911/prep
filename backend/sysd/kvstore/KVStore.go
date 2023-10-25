package kvstore

import (
	"errors"
	"fmt"
	"sync"
)

type Store interface {
	Put(key string, value string) error
	Get(key string) (value string, err error)
	Delete(key string) (err error)
}

type SimpleKVStore struct {
	store map[string]string
	mu    sync.Mutex
}

var kvStore *SimpleKVStore

func init() {
	fmt.Println("initializing the Key Value store")
	kvStore = &SimpleKVStore{store: make(map[string]string)}
}
func (this *SimpleKVStore) Put(key string, value string) error {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.store[key] = value
	return nil
}
func (this *SimpleKVStore) Get(key string) (string, error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	val, exists := this.store[key]
	if exists {
		return val, nil
	}
	return "", errors.New("key doesn't exists")
}
func (this *SimpleKVStore) Delete(key string) error {
	this.mu.Lock()
	defer this.mu.Unlock()
	delete(this.store, key)
	return nil
}
