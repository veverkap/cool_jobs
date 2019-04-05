package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://boards.greenhouse.io/github/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//*[@id="main"]/section[4]/div[1]
	// Find the review items
	doc.Find(".opening").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		item := s.Find("a")
		title := item.Text()
		link, _ := item.Attr("href")
		department_id, _ := s.Attr("department_id")

		fmt.Printf("%d: Title: %s Link: %s department_id: %s\n", i, title, link, department_id)
	})
}

func main() {
	ExampleScrape()
}
