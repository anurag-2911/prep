package robots

import (
	"fmt"
	"github.com/temoto/robotstxt"
	"net/http"
	"net/url"
	"webcrawler/useragent"
)

var (
	customuserAgent string
)

func init() {
	customuserAgent = useragent.GetCustomUserAgent()
}
func fetchRobotstxt(targetURL string) (*robotstxt.RobotsData, error) {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}
	robotsTextURL := fmt.Sprintf("%s:%s/robots.txt", parsedURL.Scheme, parsedURL.Host)
	req, err := http.NewRequest("GET", robotsTextURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", customuserAgent)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("robots.txt not found")
	}
	robotsData, err := robotstxt.FromResponse(resp)
	if err != nil {
		return nil, err
	}
	return robotsData, nil
}

func CanCrawl(targetURL string) bool {
	robotsData, err := fetchRobotstxt(targetURL)
	if robotsData == nil || err != nil {
		fmt.Printf("robot data %v, err %v\n", robotsData, err)
		return true // no robots.txt or error in reading robots.txt so proceed
	}
	parsedURl, err := url.Parse(targetURL)
	if err != nil {
		return false
	}

	isallowed := robotsData.TestAgent(parsedURl.RequestURI(), customuserAgent)
	fmt.Println("can crawl? ", targetURL, parsedURl, isallowed)
	return isallowed

}
