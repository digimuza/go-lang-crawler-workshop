package main

import (
	"fmt"

	"./crawler"
)

func main() {

	crawledLinks := make(chan string, 100)
	go crawler.Crawler("https://www.knygos.lt/", crawledLinks)

	for craw := range crawledLinks {
		fmt.Println(craw)
	}

	close(crawledLinks)
}
