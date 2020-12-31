/* package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	//  parse a complex url
	complexUrl := "https://www.example.com/path/to/?query=123&this=that#fragment"
	parseUrl, err := url.Parse(complexUrl)
	if err != nil {
		log.Fatal(err)
	}

	// print our url pieces
	fmt.Println("Scheme: " + parseUrl.Scheme)
	fmt.Println("Host: " + parseUrl.Host)
	fmt.Println("Path: " + parseUrl.Path)
	fmt.Println("Query String: " + parseUrl.RawQuery)
	fmt.Println("Fragment: " + parseUrl.Fragment)

	// Get the query key/values as a map
	fmt.Println("\nQuery values:")
	queryMap := parseUrl.Query()
	fmt.Println(queryMap)

	// Craft a new URL from scratch
	var customURL url.URL
	customURL.Scheme = "https"
	customURL.Host = "google.com"
	newQueryValues := customURL.Query()
	newQueryValues.Set("key1", "value1")
	newQueryValues.Set("key2", "value2")
	customURL.Fragment = "bookmarkLink"
	customURL.RawQuery = newQueryValues.Encode()

	fmt.Println("\nCustom URL:")
	fmt.Println(customURL.String())

}
*/