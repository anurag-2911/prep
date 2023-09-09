package lc

import (
	"fmt"
)

func TestMap() {
	kv1 := KeyValue{key: "1", value: 123}
	kv2 := KeyValue{key: "2", value: "123"}
	var mmap MyMap
	mmap.Set(kv1.key, kv2.value)
	mmap.Set(kv2.key, kv2.value)
	fmt.Println(mmap.Get("1"))
	val,exists:=mmap.Get("2")
	if(exists){
		fmt.Println(val)
	}
	for _, kv := range mmap.data {
		fmt.Println(kv.key,kv.value)
	}
}

type KeyValue struct {
	key   string
	value interface{}
}

type MyMap struct {
	data []KeyValue
}

func (m *MyMap) Set(key string, value interface{}) {
	for i, kv := range m.data {
		if kv.key == key {
			m.data[i].value = value
			return
		}

	}
	m.data = append(m.data, KeyValue{key: key, value: value})
}
func (m *MyMap) Get(key string) (interface{},bool) {
	for _, kv := range m.data {
		if kv.key == key {
			return kv.value,true
		}
	}
	return nil,false
}
func (m *MyMap) Delete(key string) {
	for i, kv := range m.data {
		if kv.key == key {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return
		}

	}

}
