package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

func main() {
	fmt.Println("main")

	logfile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("log file is not initialzed")
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	log.Println("starting the application")

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "redis-service:6379", // use the kubernetes service name
			Password: "",                   // no password set
			DB:       0,                    // use default DB
		})

	const maxRetries = 5
	const baseDelay = time.Second
	for i := 0; i < maxRetries; i++ {
		err = rdb.Ping(ctx).Err()
		if err == nil {
			break
		}
		fmt.Printf("failed to connect to redis %v ,retrying in %v", err, baseDelay<<i)
		time.Sleep(baseDelay << i) // exponential backoff
	}
	if err != nil {
		fmt.Println("count not connect to redis")
	}

	err = rdb.Set(ctx, "presname", "murmu", 0).Err()
	if err != nil {
		log.Printf("error setting key %s\n", err)
	}

	val, err := rdb.Get(ctx, "presname").Result()
	if err != nil {
		log.Println("error in getting vaalue ", err)
	} else {
		fmt.Println("value from redis for the key ", val)
		log.Println("value read is ", val)
	}
	log.Println("over")
	fmt.Println("over")

	done := make(chan bool)
	go worker(done)
	select {}

}
func worker(ch chan bool) {
	fmt.Println("I am not done")
}
