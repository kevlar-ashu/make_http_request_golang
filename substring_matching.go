/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//make http get request
	response, err := http.Get("https://tektutsgroup.blogspot.com/")
	check(err)

	defer response.Body.Close()

	// get the response body as string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	check(err)
	pageContent := string(dataInBytes)

	// find a substr
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}

	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	titleStartIndex += 7

	// find the index of closing tab
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}

	// (Optional)
	// Copy the substring in to a separate variable so the
	// variables with the full document data can be garbage collected
	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	// Print out the result
	fmt.Printf("Page title: %s\n", pageTitle)

}
*/