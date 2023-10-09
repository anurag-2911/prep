package analytics

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func AnalyzerMain() {
	fmt.Println("analyzer")
	r := gin.Default()
	// Set up CORS middleware
	config:=cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Allow only frontend server
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	r.Use(cors.New(config))

	// Routes
	r.GET("/record", func(c *gin.Context) {
		e := &Event{TimeStamp: time.Now().Unix()}
		cb.AddEvents(e)
		eventMsg:=fmt.Sprintf("events recorded %d",cb.CountEventsLastMinute())
		c.String(200, eventMsg )
	})

	r.GET("/count", func(c *gin.Context) {
		count := cb.CountEventsLastMinute()
		c.String(200, "Number of events in the last minute: %d", count)
	})

	r.Run(":8080")
}

var cb = &CircularBuffer{}

func RecordEvent(w http.ResponseWriter, r *http.Request) {
	e := &Event{TimeStamp: time.Now().Unix()}
	cb.AddEvents(e)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Event recorded %d",cb.CountEventsLastMinute())

}
func GetEventCount(w http.ResponseWriter, r *http.Request) {
	count := cb.CountEventsLastMinute()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "number of events in last minute :%d", count)
}

const BufferSize = 60

type Event struct {
	TimeStamp int64
}

type CircularBuffer struct {
	Events [BufferSize]*Event
	Pos    int
	Mutex  sync.RWMutex
}

func (this *CircularBuffer) AddEvents(e *Event) {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	this.Events[this.Pos] = e
	this.Pos = (this.Pos + 1) % BufferSize
}

func (this *CircularBuffer) CountEventsLastMinute() int {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	now := time.Now().Unix()

	count := 0

	for _, event := range this.Events {
		if event != nil && now-event.TimeStamp < 60 {
			count++
		}
	}
	return count
}
