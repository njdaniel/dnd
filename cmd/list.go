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
//TODO: Use this command with many resources, ex: dnd data ls [args]
//args = relative directory paths
//TODO: Add labels
//TODO: recursive, all data_objects(references), --all
//TODO: tree, to show all directories and files, --tree
//TODO: filter data
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the resources",
	Long: `Command to list a certain resource and pass filters
example:
	dnd data ls					# lists all files(directories) under root data
	dnd data ls items/weapons	#lists all files under weapons
`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		List(args)
	},
}

func init() {
	dataCmd.AddCommand(listCmd)

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

// List prints out all resources
func List(args []string) []string {
	// return list of directories/files directly under
	//fmt.Println("Func list called")
	listed := make([]string, 0)
	path := ""
	if len(args) != 0 {
		//fmt.Println("This is the args" + args[0])
		path = "/" + args[0]
	}
	dirName := fmt.Sprintf("./data/default-house-example%v", path)
	files, err := ioutil.ReadDir(dirName)
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
