package scraper

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {

	// Iterate over token attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func Crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("ERROR: Failed to crawl:", url)
		return
	}
	fmt.Println(resp, "GET HTTP Response")

	// Notify that we're done after this function
	// defer func() {
	// 	chFinished <- true
	// }()

	// Close Body when the function completes
	b := resp.Body
	fmt.Println(b, "GET RESPONSE BODY")
	defer b.Close()
}
