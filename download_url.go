/* package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// make a request
	response, err := http.Get("https://tektutsgroup.blogspot.com/")

	check(err)

	defer response.Body.Close()

	// create a output file
	outFile, err := os.Create("output.html")
	check(err)

	defer outFile.Close()

	// copy data from http response to file
	_, err = io.Copy(outFile, response.Body)
	check(err)
}
*/