package main

import (
	"encoding/json"
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
	Name         string `json:"name"`
	Price        string `json:"price"`
	Damage       string `json:"damage"`
	Range        string `json:"range"`
	Reload       string `json:"reload"`
	Bulk         string `json:"bulk"`
	Hands        string `json:"hands"`
	Group        string `json:"group"`
	WeaponTraits string `json:"weapon_traits"`
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
			//fmt.Println(element.ChildText("td:nth-child(1)"))
			//fmt.Println(element.ChildText("td:nth-child(2)"))
			//fmt.Println(element.ChildText("td:nth-child(3)"))
			//fmt.Println(element.ChildText("td:nth-child(4)"))
			//fmt.Println(element.ChildText("td:nth-child(5)"))
			//fmt.Println(element.ChildText("td:nth-child(6)"))
			//fmt.Println(element.ChildText("td:nth-child(7)"))
			//fmt.Println(element.ChildText("td:nth-child(8)"))
			//fmt.Println("")
			item.Name = element.ChildText("td:nth-child(1)")
			item.Price = element.ChildText("td:nth-child(2)")
			item.Damage = element.ChildText("td:nth-child(3)")
			//TODO: Check if 'ft.' is any non range weapons
			if strings.Contains(element.ChildText("td:Contains(ft)"), "ft.") {
				fmt.Println("RANGED WEAPON")
				item.Range = element.ChildText("td:nth-child(4)")
				item.Reload = element.ChildText("td:nth-child(5)")
				item.Bulk = element.ChildText("td:nth-child(6)")
				item.Hands = element.ChildText("td:nth-child(7)")
				item.Group = element.ChildText("td:nth-child(8)")
				item.WeaponTraits = element.ChildText("td:nth-child(9)")
			} else {
				item.Range = "melee"
				item.Reload = "-"
				item.Bulk = element.ChildText("td:nth-child(4)")
				item.Hands = element.ChildText("td:nth-child(5)")
				item.Group = element.ChildText("td:nth-child(6)")
				item.WeaponTraits = element.ChildText("td:nth-child(7)")
			}
		})
		items = append(items, item)
	})

	// Start scraping http://pf2playtest.opengamingnetwork.com/equipment/weapons/
	c.Visit("http://pf2playtest.opengamingnetwork.com/equipment/weapons/")

	//fmt.Println(items)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(items)

	//TODO: Dump into json file
	itemsJson, _ := json.Marshal(items)
	err := ioutil.WriteFile("weapons.json", itemsJson, 0644)
	if err != nil {
		fmt.Errorf("error creating json file")
	}
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
