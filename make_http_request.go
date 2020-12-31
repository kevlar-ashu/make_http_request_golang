/* package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//make http get request
	response, err := http.Get("https://tektutsgroup.blogspot.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// copy data from response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT", n)

}
*/