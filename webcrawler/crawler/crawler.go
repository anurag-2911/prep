package crawler

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"webcrawler/robots"
	"webcrawler/useragent"
)

func Crawl() {

}

type CrawlResult struct {
	URL   string
	Body  string
	Error error
}

func CrawlURL(url string, wg *sync.WaitGroup, results chan<- CrawlResult) {
	if robots.CanCrawl(url) {
		defer wg.Done()

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			results <- CrawlResult{URL: url, Error: err}
			return
		}
		req.Header.Add("User-Agent", useragent.GetCustomUserAgent())
		client := &http.Client{}
		resp, err := client.Do(req)
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
	} else {
		fmt.Println("cannot crawl on this url ", url)
	}

}
