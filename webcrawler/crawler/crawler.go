package crawler

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func Crawl() {

}

type CrawlResult struct {
	URL   string
	Body  string
	Error error
}

func CrawlURL(url string, wg *sync.WaitGroup, results chan<- CrawlResult) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- CrawlResult{URL: url, Body: "", Error: err}
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- CrawlResult{URL: url, Error: err}
		return
	}

	fmt.Printf("returning result for %s\n", url)
	results <- CrawlResult{URL: url, Body: string(body), Error: nil}
	

}
