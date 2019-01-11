package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

}

//PullTable grab data from table in html and create json file
func PullTable() {
	//open output.html

	// get <table with children

}

//PullHtml pulls down html
func PullHtml() {
	//http://pf2playtest.opengamingnetwork.com/equipment/weapons/

	// make request
	resp, err := http.Get("http://pf2playtest.opengamingnetwork.com/equipment/weapons/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//create output file
	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	//copy data from resp to outfile
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
