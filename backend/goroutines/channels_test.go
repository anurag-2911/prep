package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestXxx09(t *testing.T) {

}

type LogMessage struct {
	TimeStamp time.Time
	Message   string
}

type Logger struct {
}

func processLogMessage(logchannel <-chan LogMessage) {
	for log := range logchannel {
		fmt.Printf("[%v] %s\n", log.TimeStamp.Format("2006-01-02 15:04:05"), log.Message)
	}
}
func TestProcessLogMessage(t *testing.T) {
	logChannel := make(chan LogMessage, 10)
	go processLogMessage(logChannel)

	for i := 0; i < 5; i++ {
		log := &LogMessage{TimeStamp: time.Now(), Message: fmt.Sprintf("message : %d", i)}
		logChannel <- *log
		time.Sleep(time.Second)

	}
	close(logChannel)
}

func sendToChannel(sendOnly chan<- int){
	for i:=0;i<5;i++{
		fmt.Printf("sending %d\n",i)
		sendOnly<-i
		time.Sleep(time.Second)
	}
	close(sendOnly)
}
func receiveFromChannel(receiveOnly <-chan int){
	for value:=range receiveOnly{
		fmt.Printf("received %d\n",value)
	}
	fmt.Println("done receiving")
}
func TestSendRecieve(t *testing.T){
	mychan:=make(chan int)
	go sendToChannel(mychan)
	go receiveFromChannel(mychan)

	time.Sleep(time.Minute)
}

type DistTask struct{
	ID int
	Data string
}

func(task *DistTask)producer(id int,taskQueue chan<- DistTask,wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<5;i++{
		task:=&DistTask{ID:id*100+i,Data:fmt.Sprintf("task from the producer %d",i)}
		fmt.Printf("producer %d sending task %d\n",i,task.ID)
		taskQueue<-*task
		time.Sleep(time.Second) // simlulate some work

	}

}
func(task *DistTask)consumer(id int,taskQueue <-chan DistTask,wg *sync.WaitGroup){
	defer wg.Done()
	for task:=range taskQueue{
		fmt.Printf("consumer %d processed task %d\n",id,task.ID)
		time.Sleep(time.Second)
	}
}
func TestProCons(t *testing.T){
	task:=DistTask{}
	const producers=5
	const consumers =4
	var wg sync.WaitGroup

	taskQueue:=make(chan DistTask,10)

	//start producers
	for i:=0;i<producers;i++{
		wg.Add(1)
		go task.producer(i,taskQueue,&wg)
	}
	//start consumers
	for i:=0;i<consumers;i++{
		wg.Add(1)
		go task.consumer(i,taskQueue,&wg)

	}
	// wait for all the producers and close the task queue
	go func(){
		wg.Wait()
		close(taskQueue)
	}()

	wg.Wait()
	fmt.Println("all tasks done")
}

func startTaskSceduler(interval time.Duration, stopChan <- chan struct{}){
	ticker:=time.NewTicker(interval)

	defer ticker.Stop()
	count:=0
	for {
		select{
		case <-ticker.C:
			count++
			tickerTask(string(rune(count)))
		case <-stopChan:
			fmt.Println("scheduler stopped")
			return
			
		}
		
	}
	
}
func tickerTask(name string){
	fmt.Println("running task ",name)
}
func TestStartTaskScheduler(t *testing.T){
	stopchan:=make(chan struct{})

	go startTaskSceduler(time.Second,stopchan)

	time.Sleep(10*time.Second)

	stopchan<-struct{}{}

	time.Sleep(time.Minute)
}

func makeApiReq(id int){
	fmt.Println("requesting for id ",id)
	time.Sleep(time.Second*2)
}
func requestScheduler(requests <-chan int,ratelimit time.Duration){
	tiker:=time.NewTicker(ratelimit)

	defer tiker.Stop()

	for id:=range requests{
		<-tiker.C
		go makeApiReq(id)
	}

}
func TestRequestScheduler(t *testing.T){
	requests:=make(chan int, 5)

	go requestScheduler(requests,5*time.Second)

	for i:=0;i<10;i++{
		select{
		case requests<-i:
			fmt.Println("scheduled requests ",i)
		case <-time.After(5*time.Second):
			fmt.Println("request timesd out")
			return
		}

	}

	time.Sleep(time.Minute)
}

