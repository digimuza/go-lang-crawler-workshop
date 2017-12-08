package crawler

import (
	"net/http"

	"github.com/jackdanger/collectlinks"
	"golang.org/x/sync/syncmap"
)

//We need to check if url is alredy visited // Normal map do not work becouse it's throws concurent read/write
var visited = syncmap.Map{}

// Crawler function that starsts crawling process
func Crawler(url string, crawledLinks chan string) {

	_, ok := visited.Load(url)
	if !ok {
		resp, err := http.Get(url)

		if err == nil {
			visited.Store(url, true)
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

}
