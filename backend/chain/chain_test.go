package chain

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	test01()
}

func test01(){
	ch:=make(chan int)
	go func(){
		ch<-42
	}()
	fmt.Println(<-ch)
}