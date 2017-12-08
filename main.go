package main

import (
	"fmt"

	"./crawler"
)

func main() {

	crawledLinks := make(chan crawler.SiteCrawl, 100)
	go crawler.Crawler(crawler.CrawlLink{"https://rekvizitai.vz.lt/", "https://rekvizitai.vz.lt/"}, crawledLinks)

	for craw := range crawledLinks {
		fmt.Println(craw.CrawlLink.URL)
	}

	close(crawledLinks)
}
