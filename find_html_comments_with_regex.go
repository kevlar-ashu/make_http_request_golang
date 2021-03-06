/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// make http request
	response, err := http.Get("https://tektutsgroup.blogspot.com/")
	check(err)

	defer response.Body.Close()

	// read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}
	// Create a regular expression to find comments
	re := regexp.MustCompile("<!--(.|\n)*?-->")
	comments := re.FindAllString(string(body), -1)
	if comments == nil {
		fmt.Println("No matches.")
	} else {
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}
}
*/