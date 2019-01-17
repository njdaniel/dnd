package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Item struct {
	Name         string
	Price        string
	Damage       string
	Bulk         string
	Hands        string
	Group        string
	WeaponTraits string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	fmt.Println("Scraping")

	items := make([]Item, 0, 500)

	// Find table
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		fmt.Println("found tbody?")
		table := e.Attr("tbody")
		fmt.Println(table)
		item := Item{}
		e.ForEach("tr", func(_ int, element *colly.HTMLElement) {

			switch element.ChildText("td:first-child") {
			case "Name":
				item.Name = element.ChildText("td:nth-child(2)")
			case "Price":
				item.Price = element.ChildText("td:nth-child(2)")
			case "Damage":
				item.Damage = element.ChildText("td:nth-child(2)")
			case "Bulk":
				item.Bulk = element.ChildText("td:nth-child(2)")

			}
		})
		items = append(items, item)
	})

	// Start scraping http://pf2playtest.opengamingnetwork.com/equipment/weapons/
	c.Visit("http://pf2playtest.opengamingnetwork.com/equipment/weapons/")

	fmt.Println(items)

}

//PullTable grab data from table in html and create json file
func MyPullTable() {
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
