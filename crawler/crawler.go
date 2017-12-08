package crawler

import (
	"net/http"
	"net/url"

	"github.com/jackdanger/collectlinks"
	"golang.org/x/sync/syncmap"
)

//We need to check if url is alredy visited // Normal map do not work becouse it's throws concurent read/write
var visited = syncmap.Map{}

var baseURL = ""

// Crawler function that starsts crawling process
func Crawler(url string, crawledLinks chan string) {

	if baseURL == "" {
		baseURL = url
	}

	//Check if url is visited if not crawl it
	_, ok := visited.Load(url)
	if !ok && canCrawl(url) {
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

func canCrawl(site string) bool {
	base, bErr := url.Parse(baseURL)
	siteBase, sErr := url.Parse(site)

	if bErr != nil || sErr != nil {
		return false
	}

	return bool(base.Host == siteBase.Host)
}
