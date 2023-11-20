package goroutines

import (
	"fmt"
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
