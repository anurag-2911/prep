package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestFanInFanout(t *testing.T) {
	in:=make(chan int, 0)
	numofWorkers:=4
	//start workers
	workerOuts:=fanout(in,numofWorkers)

	//merge the results from workers
	results:=fanin(workerOuts)

	//send data to workers
	go func ()  {
		for i:=0;i<10;i++{
			in<-i
		}
		close(in)
	}()

	for r:=range results{
		fmt.Println("processed ",r)
	}
}

func xworker(n int, chin <-chan int, chout chan<- int) {
	for n := range chin {
		chout <- n % 10
	}

}
func fanout(in <-chan int, numofWorkers int) []chan int {
	outs := make([]chan int, numofWorkers)
	for i := range outs {
		outs[i] = make(chan int)
		go xworker(i, in, outs[i])
	}
	return outs
}

func fanin(ins []chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(ins))
	for _, in := range ins {
		go output(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
