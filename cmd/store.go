/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/njdaniel/dnd/services/store"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store is a business establishment",
	Long: `Store is a resource of a business establishment. Allowing for generating a new store, updating the and viewing inventory, searching and closing. For example:

    $ dnd store create
	Store Created!
	Name: Prancy Pony
	StoreType: Tavern
	Owner: Overhill
	Inventory:
	  - name: "pint of ale"
	    price: "5cp"
	  - name: "rabbit stew"
	    price: "8cp"
	Money:
	  CP: 200
	  SS: 50
	  GC: 8
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("store called")
		//name := ""
		//owner := ""
		//location := ""
		//storeType := ""
		//cp := 0
		//ss := 0
		//gc := 0
		//money := store.NewMoney(cp, ss, gc)
		//items := []store.Item{}
		//name = GetFlagString(cmd, "name")
		//owner = GetFlagString(cmd, "owner")
		//location = GetFlagString(cmd, "location")
		//storeType = GetFlagString(cmd, "storeType")
		//cp = GetFlagInt(cmd, "cp")
		//ss = GetFlagInt(cmd, "ss")
		//gc = GetFlagInt(cmd, "gc")
		//ns := store.NewStore("fletcher")
		builder := store.GetStoreBuilder("Fletcher")
		ns := builder.GetStore()
		//var data []byte
		data, err := json.MarshalIndent(ns, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))

	},
}

func init() {
	rootCmd.AddCommand(storeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//createCmd.Flags().StringP("output", "o", "json", "Output in json|simple|yaml")
	//createCmd.Flags().String("name", "", "give custom name")
	//createCmd.Flags().String("owner", "", "give name of owner")
	//createCmd.Flags().String("storeType", "", "give store type")
	//createCmd.Flags().String("location", "", "give location of store")
	//createCmd.Flags().Int("cp", 0, "Amount of Copper Pieces")
	//createCmd.Flags().Int("ss", 0, "Amount of Silver Shillings")
	//createCmd.Flags().Int("gc", 0, "Amount of Gold Crowns")
}
