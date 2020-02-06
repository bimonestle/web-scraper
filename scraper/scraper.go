package scraper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchURL(u string) {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("HTML:\n\n", string(bytes))
	resp.Body.Close()
}

func getAnchor() {
	fmt.Println("Get Anchor")
}
