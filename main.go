package main

import (
	"fmt"
	"net/http"

	"github.com/jackdanger/collectlinks"
)

func main() {

	crawledLinks := make(chan string, 100)
	go Crawler("https://www.google.lt/search?dcr=0&source=hp&ei=kksqWuWQCuTU6ASqtbPwBQ&q=alus&oq=alus&gs_l=psy-ab.3..0l10.1938.2470.0.2698.5.4.0.0.0.0.104.301.1j2.3.0....0...1c.1.64.psy-ab..2.3.300.0..35i39k1j0i131k1.0.CkUwPC4GWZo", crawledLinks)

	for craw := range crawledLinks {
		fmt.Println(craw)
	}

	close(crawledLinks)
}

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
