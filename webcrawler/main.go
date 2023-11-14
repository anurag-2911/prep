package main

import (
	"fmt"
	"sync"
	"webcrawler/crawler"
)

func main() {
	urls := []string{
		"www.google.com",
		"http://news.google.com",
		
	}
	var wg sync.WaitGroup

	results := make(chan crawler.CrawlResult, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go crawler.CrawlURL(url, &wg, results)
	}

	//close the channel in the background
	go func() {
		fmt.Println("waiting for the go routines")
		wg.Wait()
		fmt.Println("closing the channels")
		close(results)
		fmt.Println("done closing the channels")
	}()

	//process results

	for result := range results {
		if result.Error != nil {
			fmt.Printf("error crawling %v:%v\n", result.URL, result.Error)
			continue
		}
		fmt.Printf("result of crawling the url %v:%v\n", result.URL, result.Body[0:100])
	}
}
