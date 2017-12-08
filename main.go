package main

import (
	"fmt"
	"net/http"

	"github.com/jackdanger/collectlinks"
)

func main() {
	resp, err := http.Get("https://www.wikipedia.org/")

	if err == nil {
		links := collectlinks.All(resp.Body)
		fmt.Println(links)
	}
}
