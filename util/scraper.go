package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	PullTable()
}

//PullTable grab data from table in html and create json file
func PullTable() {
	//open output.html
	dataInBytes, err := ioutil.ReadFile("output.html")
	if err != nil {

	}
	pageContent := string(dataInBytes)

	// get <table with children
	tableStartIndex := strings.Index(pageContent, "<table")
	if tableStartIndex == -1 {
		fmt.Println("No table")
		os.Exit(0)
	}
	tableEndIndex := strings.Index(pageContent, "</table>")
	if tableEndIndex == -1 {
		fmt.Println("No closing tag for table")
	}
	pageTable := []byte(pageContent[tableStartIndex:tableEndIndex])
	fmt.Println(pageTable)

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
