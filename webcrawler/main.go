package main

import (
	"fmt"
	"sync"
	"webcrawler/crawler"
	"webcrawler/requestratelimit"
)

func main() {
	urls := []string{
		"www.google.com",
		"http://news.google.com",
	}
	var wg sync.WaitGroup

	results := make(chan crawler.CrawlResult, len(urls))

	for _, url := range urls {
		requestratelimit.HandleRateLimit(url)
		wg.Add(1)
		go crawler.CrawlURL(url, &wg, results)
	}

	//close the channel in the background
	go func() {
		wg.Wait()
		close(results)
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
