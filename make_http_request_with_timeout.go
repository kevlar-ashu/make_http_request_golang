/* package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//create http client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// make request
	response, err := client.Get("https://tektutsgroup.blogspot.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// copy data from response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}
*/