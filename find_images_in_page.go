package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func processElement(index int, element *goquery.Selection) {
	imgSrc, exists := element.Attr("src")
	if exists {
		fmt.Println(imgSrc)
	}
}

func main() {
	// make http request
	response, err := http.Get("https://tektutsgroup.blogspot.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// create a goquery document from http  response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//find and print image urls
	document.Find("img").Each(processElement)
}
