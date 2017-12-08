package crawler

import (
	"net/http"

	"github.com/jackdanger/collectlinks"
)

// Crawler function that starsts crawling process
func Crawler(url string, crawledLinks chan string) {

	resp, err := http.Get(url)

	if err == nil {
		linksChannel := make(chan string, 10000)

		links := collectlinks.All(resp.Body)

		crawledLinks <- url

		for _, link := range links {
			go Crawler(link, crawledLinks)
		}

		close(linksChannel)
		resp.Body.Close()
	}

}
