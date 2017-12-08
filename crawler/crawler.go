package crawler

import (
	"net/http"
	"net/url"

	"github.com/jackdanger/collectlinks"
	"golang.org/x/sync/syncmap"
)

// CrawlLink location type
type CrawlLink struct {
	ExtractedFrom string
	URL           string
}

// SiteCrawl Definition for crawl
type SiteCrawl struct {
	CrawlLink  CrawlLink
	StatusCode int
	LinksFound int
	Links      []string
}

//We need to check if url is alredy visited // Normal map do not work becouse it's throws concurent read/write
var visited = syncmap.Map{}

var baseURL = ""

// Crawler function that starsts crawling process
func Crawler(crawl CrawlLink, crawledLinks chan SiteCrawl) {

	if baseURL == "" {
		baseURL = crawl.URL
	}

	//Check if url is visited if not crawl it
	_, ok := visited.Load(crawl.URL)
	if !ok && canCrawl(crawl.URL) {
		visited.Store(crawl.URL, true)
		resp, err := http.Get(crawl.URL)

		if err == nil {

			linksChannel := make(chan CrawlLink, 10000)

			links := collectlinks.All(resp.Body)

			crawledLinks <- SiteCrawl{
				crawl,
				resp.StatusCode,
				len(links),
				links,
			}

			for _, link := range links {
				go Crawler(CrawlLink{crawl.URL, link}, crawledLinks)
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
