package goroutines

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var counter int
var mu sync.Mutex

func TestAtomic(t *testing.T) {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println("number of go routines executed ", counter)

}
func TestClosingchannel(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "halo"
	ch <- "hello"
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- "hola"
}

func TestEmptyChannel(t *testing.T) {
	ch := make(chan int, 0)

	go func() {
		i := <-ch
		fmt.Println("value of i ", i)
	}()
	// ch<-45
	select {}
}
func TestPanic(t *testing.T) {
	withpanic()
	fmt.Println("after a panic")
}
func withpanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic ", r)
		}
	}()
	panic("bhago bhago")

}
func TestBufferedChannel(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- "halo"
	fmt.Println("done sending")
}

func TestSelectfn(t *testing.T) {
	ch := make(chan int, 0)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()
	for i := 0; i < 2; i++ {
		select {

		case <-ch:
			{
				fmt.Println("value received ", <-ch)
			}
		default:
			{
				fmt.Println("default case")
			}
		}
	}
	time.Sleep(4 * time.Second)
}
func TestWG(t *testing.T) {
	var wg sync.WaitGroup
	x := 0
	wg.Add(1)
	go worker009(&x, &wg)
	wg.Wait()
	fmt.Println("value of x ", x)
}
func worker009(i *int, wg *sync.WaitGroup) {
	defer wg.Done()
	(*i)++
	fmt.Println("value of i ", *i)
}
func TestMuLock(t *testing.T) {
	mu.Lock()
	defer mu.Unlock()
	counter++
}
func TestSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Second)
		ch1 <- "hello from ch1"
	}()
	go func() {
		time.Sleep(100 * time.Second)
		ch2 <- "hello from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			{
				fmt.Println(msg)
			}
		case msg := <-ch2:
			{
				fmt.Println(msg)
			}
		}
	}
}
func TestDataonClosedChannel(t *testing.T) {
	closedchannel()
	fmt.Println("done with this test")

}

func closedchannel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from painc", r)
		}
	}()
	ch := make(chan int, 0)

	close(ch)

	ch <- 1

	fmt.Println("after closing the channel")
}
func TestXxx(t *testing.T) {
	fmt.Println()
}

func TestSum(t *testing.T) {

	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr2 := []int{10, 20, 30, 40, 50, 60, 70, 80}

	callSum(arr1, arr2)
}
func BenchMarkSum(b *testing.B) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr2 := []int{10, 20, 30, 40, 50, 60, 70, 80}
	n := b.N
	fmt.Println("benchmarking times ", n)
	b.ResetTimer()
	for i := 0; i < n; i++ {
		callSum(arr1, arr2)
	}
}
func callSum(arr1 []int, arr2 []int) {
	ch := make(chan int, 2)
	go summe(arr1, ch)
	go summe(arr2, ch)
	fmt.Println("number of go routines ", runtime.NumGoroutine())
	<-ch
	<-ch
}

func summe(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	ch <- sum
}

// worker pool

// define the job
type Job struct {
	data   []int
	result chan int
}

func worker(jobs <-chan Job) {
	for job := range jobs {
		sum := 0
		for _, v := range job.data {
			sum += v
		}
		job.result <- sum
	}
}

func TestWorkerPool(t *testing.T) {
	data := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}
	//create channels for job and result
	jobs := make(chan Job, len(data))
	results := make(chan int, len(data))

	//start the workers
	for w := 0; w < len(data); w++ {
		go worker(jobs)
	}

	//submit jobs

	for _, d := range data {
		jobs <- Job{data: d, result: results}
	}
	close(jobs)

	for i := 0; i < len(data); i++ {
		sum := <-results
		fmt.Println(sum)
	}

}

func performSimpleTask(taskid int) {
	fmt.Printf(" task started %d", taskid)
	fmt.Print(" > ")
	time.Sleep(2 * time.Second)
	fmt.Println()
	fmt.Printf(" task ended %d", taskid)
	fmt.Print(" > ")
}

func runsimpleTask() {
	for i := 0; i < 5; i++ {
		go performSimpleTask(i)
	}

	time.Sleep(12 * time.Second)
	fmt.Printf("all tasks are done\n")
}

func TestSimpleTask(t *testing.T) {
	runsimpleTask()
}

func TestConcurrentPrimeNUmberFinder(t *testing.T) {
	startTime := time.Now()
	ConcurrentPrimeNumberFinder()
	endtime := time.Since(startTime)
	fmt.Println("time taken ", endtime)
}
func ConcurrentPrimeNumberFinder() {
	start, end := 1, 10000
	primes := make(chan int)
	var wg sync.WaitGroup
	for i := start; i <= end; i = i + 1000 {
		wg.Add(1)
		go findPrimes(i, i+999, &wg, primes)
	}

	go func() {
		wg.Wait()
		close(primes)
	}()

	count := 0
	for prime := range primes {
		fmt.Printf(" prime %d ", prime)
		count++
	}
	fmt.Println("number of prime numbers ", count)
}

func isPrime(n int) bool {
	if n <= 2 {
		return true
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func findPrimes(start int, end int, wg *sync.WaitGroup, primes chan<- int) {
	defer wg.Done()
	fmt.Println("find primes is called")
	for num := start; num < end; num++ {
		if isPrime(num) {
			primes <- num
		}
	}
}

func TestCocurrentDownload(t *testing.T){
	concurrentFileDownloader()
}

func concurrentFileDownloader() {
	var wg sync.WaitGroup
	urls := []string{
		"http://example.com/file1.jpg",
		"http://example.com/file2.jpg",
	}
	for _,url:=range urls{
		wg.Add(1)
		go downLoadFile(url,&wg)
	}
	wg.Wait()

	fmt.Println("all files are downloaded")
}
func downLoadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create("output_" + url[strings.LastIndex(url, "/")+1:])
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	fmt.Println("downloaded ", url)

}

func countWords(filePath string){
	
}

// log message processor
