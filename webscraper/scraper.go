package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.cricbuzz.com/live-cricket-scores/87829/afg-vs-aus-48th-match-super-8-group-1-icc-mens-t20-world-cup-2024"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error: Status code %d\n", res.StatusCode)
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find and print all paragraph texts
	doc.Find("p").Each(func(index int, element *goquery.Selection) {
		paragraph := element.Text()
		fmt.Println(paragraph)
	})

}
