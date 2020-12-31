/* package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exits := element.Attr("href")
	if exits {
		fmt.Println(href)
	}
}

func main() {
	//make http request
	response, err := http.Get("https://tektutsgroup.blogspot.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loding http response body", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)
}
*/