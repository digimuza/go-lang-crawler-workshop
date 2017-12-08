package main

import (
	"flag"
	"fmt"

	"./crawler"
)

func main() {
	baseURL := flag.String("url", "", "Crawl domain URL")

	flag.Parse()

	if *baseURL != "" {
		crawledLinks := make(chan crawler.SiteCrawl, 100)
		go crawler.Crawler(crawler.CrawlLink{*baseURL, *baseURL}, crawledLinks)

		for craw := range crawledLinks {
			fmt.Println(craw.CrawlLink.URL)
		}

		close(crawledLinks)
	} else {
		fmt.Println("Please provide url -url=")
	}

}
