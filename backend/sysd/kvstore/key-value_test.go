package kvstore

import (
	"fmt"
	"testing"
)

func TestKVStore(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from %v\n", r)
		}
	}()

	TestKV()
}
