package main

import (
	"flag"
	"fmt"
	"strings"

	"./crawler"
)

func main() {
	baseURL := flag.String("url", "", "Crawl domain URL")

	status := flag.String("status", "200,500", "Output urls with status")

	flag.Parse()

	statusList := strings.Split(*status, ",")
	if *baseURL != "" {
		fmt.Println(fmt.Sprintf("Crawling... searching links with status %s", *status))
		crawledLinks := make(chan crawler.SiteCrawl, 100)
		go crawler.Crawler(crawler.CrawlLink{*baseURL, *baseURL}, crawledLinks)

		for craw := range crawledLinks {

			if contains(statusList, fmt.Sprintf("%d", craw.StatusCode)) {
				fmt.Println(craw.CrawlLink.URL)
			}

		}

		close(crawledLinks)
	} else {
		fmt.Println("Please provide url -url=")
	}

}

func contains(strings []string, e string) bool {
	for _, a := range strings {
		if a == e {
			return true
		}
	}
	return false
}
