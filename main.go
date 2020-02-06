package main

import (
	"fmt"

	"github.com/bimonestle/web-scraper/scraper"
)

func main() {
	fmt.Println("Hello")
	scraper.Crawl("https://golang.org")
}
