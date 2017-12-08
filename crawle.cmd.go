package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./crawler"
)

func main() {
	baseURL := flag.String("url", "", "Crawl domain URL")
	logFile := flag.String("log", "", "Log urls to file")
	status := flag.String("status", "200,500", "Output urls with status")

	flag.Parse()

	statusList := strings.Split(*status, ",")
	if *baseURL != "" {
		fmt.Println(fmt.Sprintf("Crawling... searching links with status %s", *status))
		crawledLinks := make(chan crawler.SiteCrawl, 100)
		go crawler.Crawler(crawler.CrawlLink{*baseURL, *baseURL}, crawledLinks)

		for craw := range crawledLinks {

			if contains(statusList, fmt.Sprintf("%d", craw.StatusCode)) {

				if *logFile != "" {
					logToFile(*logFile, craw)
				}

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

func logToFile(filename string, crawled crawler.SiteCrawl) {

	if !exists("logs") {

		os.MkdirAll("logs", os.ModePerm)
	}

	folder := fmt.Sprintf("logs/%s", filename)
	if !exists(folder) {
		os.Create(folder)
	}

	f, _ := os.OpenFile(folder, os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString(fmt.Sprintf("FROM - %s | URL - %s | Status - %d \n", crawled.CrawlLink.ExtractedFrom, crawled.CrawlLink.URL, crawled.StatusCode))
	f.Close()
}

// Exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
