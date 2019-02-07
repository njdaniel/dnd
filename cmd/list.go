// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// listCmd represents the list command
//TODO: Use this command with many resources, ex: dnd spells ls, dnd items ls, dnd creatures ls
//TODO: Add labels
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the resources",
	Long:    `Command to list a certain resource and pass filters`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		List()
	},
}

func init() {
	itemsCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//TODO: list by weapons, armor, gear
	//TODO: filter by fields
	//ex:
	//	"name": "Staff",
	//	"price": "0 sp",
	//	"damage": "1d4 B",
	//	"range": "melee",
	//	"reload": "-",
	//	"bulk": "1",
	//	"hands": "1",
	//	"group": "Club",
	//	"weapon_traits": "Two-hand d8"
	//TODO: list by json vs text name
}

// Weapon struct from scraped weapon page
// example:
//	"name": "Staff",
//	"price": "0 sp",
//	"damage": "1d4 B",
//	"range": "melee",
//	"reload": "-",
//	"bulk": "1",
//	"hands": "1",
//	"group": "Club",
//	"weapon_traits": "Two-hand d8"
type Weapon struct {
	Name         string `json:"name"`
	Price        string `json:"price"`
	Damage       string `json:"damage"`
	WeaponRange  string `json:"range"`
	Reload       string `json:"reload"`
	Bulk         string `json:"bulk"`
	Hands        string `json:"hands"`
	Group        string `json:"group"`
	WeaponTraits string `json:"weapon_traits"`
}

// List prints out all resources
func List() []string {
	// return list of directories/files directly under
	//fmt.Println("Func list called")
	listed := make([]string, 0)
	//TODO: insert parameter into data path
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		fmt.Errorf("%v", err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		listed = append(listed, f.Name())
	}
	return listed

	//items in json
	//loop through jsons in /items directory

	//open json file
	//jsonFile, err := os.Open("../data/pf2playtest/weapons.json")
	//if err != nil {
	//	//TODO: return err? use RunE instead of Run?
	//	fmt.Println(err)
	//}
	//defer jsonFile.Close()

	//parse into struct

	//print out each name

}
