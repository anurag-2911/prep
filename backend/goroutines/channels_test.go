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

func (logger *Logger) processLogMessage(logchannel <-chan LogMessage) {
	for log := range logchannel {
		fmt.Printf("[%v] %s\n", log.TimeStamp.Format("2006-01-02 15:04:05"), log.Message)
	}
}
