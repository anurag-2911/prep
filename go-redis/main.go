package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
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
}
