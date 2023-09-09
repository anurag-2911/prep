package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	question10()
}
func question10() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		fmt.Println("going to write 1")
		ch <- 1
		fmt.Println("going to write 2")
		ch <- 2
		fmt.Println("writing done")
	}()
	
	for v := range ch {
		fmt.Println("reading ",v)
		// fmt.Println(v)
		fmt.Println("read ",v)
	}
}
func question09() {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from panic ", r)
			}
		}()

		ch <- 2
	}()
	v, ok := <-ch
	fmt.Println(v, ok)
}
func test10() {
	ch := make(chan int)
	go func() {
		close(ch)
	}()
	v, ok := <-ch
	fmt.Println(v, ok)
}
func test09() {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)
	v, ok := <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
}
func test07() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()
	fmt.Println(<-ch)
}

func test06() {
	ch := make(chan int)

	select {
	case <-ch:
	default:
		fmt.Println("default case")
	}
}

func test05() {
	ch := make(chan int, 2)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}
func test04() {
	ch := make(chan int)
	ch <- 1
	close(ch)
	ch <- 2
	fmt.Println(<-ch)
}
func test03() {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 42
	}()
	wg.Wait()
	fmt.Println(<-ch)

}
func test02() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
func test01() {
	ch := make(chan int)
	go func() {
		fmt.Println("I am about to write")
		time.Sleep(10 * time.Hour)
		ch <- 42

		fmt.Println("I have written")
	}()
	fmt.Println("I am about to read")
	fmt.Println(<-ch)
	fmt.Println("I have read now")
}
func newFunction() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("main ")
}
func producer(ch chan int, wg *sync.WaitGroup) {
	count := 1
	for count <= 10 {
		ch <- count
		time.Sleep(time.Millisecond * 7)
		count++
	}
	fmt.Println("done with produce")
	close(ch)
	wg.Done()
}
func consumer(ch chan int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("done with consuming")
	wg.Done()

}

// type MMap struct {
// 	key string
// 	value string
// 	count int

// }
// var hashIndex []int
// var hashValues []string

// var myMap []MMap
// func init(){
// 	myMap=make([]MMap, 500)
// }
// func(mp *MMap)Add(key string,value string){
// 	current := &MMap{key: getHashedKey(key),value:value}
// 	if(len(myMap)<=mp.count){
// 	index:=getIndex(key)
// 	myMap[index]=value
// 	}

// }
// func(mp *MMap)Remove(key string,value string){

// }
// func(mp *MMap)Contains(key string,value string)(bool){

// 	return false
// }
// func getHashedKey(k string)(string){

// 	return k // todo:
// }
// func (mp *MMap)GetAllEntries()(){

// }
