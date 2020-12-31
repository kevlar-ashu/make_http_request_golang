/* package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// create http client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// create and modify http request before sending
	request, err := http.NewRequest("GET", "https://tektutsgroup.blogspot.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Not Firefox")

	// make request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
*/