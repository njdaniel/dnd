/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"strings"

	"github.com/njdaniel/dnd/services/commands/character"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new character",
	Long: `Create a new character

	Pass in arguments, weight to random choices. Save to file, push to server.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		nc := character.NewCharacter()
		var data []byte
		if GetFlagString(cmd, "output") == "yaml" {
			//fmt.Println(nc)
			d, err := yaml.Marshal(nc)
			if err != nil {
				log.Fatal(err)
			}
			data = d
			fmt.Println(string(data))
		} else {
			d, err := json.MarshalIndent(nc, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			data = d
			fmt.Println(string(data))
		}
		if GetFlagBool(cmd, "save") {
			if err := createCharacterFile(nc, data); err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	characterCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().Bool("save", false, "save to default location $HOME/.dnd/characters/{{name-profession}}.json")
	createCmd.Flags().StringP("output", "o", "json", "Output in json|simple|yaml")
}

//GetFlagBool returns the value of bool flag
func GetFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.Flags().GetBool(flag)
	if err != nil {
		log.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return b
}

//GetFlagString return value of string flag
func GetFlagString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return s
}

func createCharacterFile(nc character.Character, b []byte) error {
	fmt.Println("\n save flag is passed")
	home := os.Getenv("HOME")
	baseDir := fmt.Sprintf(home + "/.dnd/character/")
	//strip whitespace
	fullName := strings.Replace(nc.Name, " ", "", -1)
	firstProfession := strings.Replace(nc.Professions[0], " ", "", -1)
	filename := fmt.Sprintf("%s-%s.json", fullName, firstProfession)
	fmt.Printf("creating character file %s\n", filename)
	filepath := fmt.Sprintf(baseDir + filename)
	fmt.Println(filepath)
	//make dir
	//$HOME/.dnd/character/filename
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directories %s: %v", baseDir, err)
	}
	//check if file already exists
	if _, err := os.Stat(filepath); os.IsExist(err) {
		return fmt.Errorf("error file %s already exists: %v \n", filepath, err)
	} else if os.IsNotExist(err) {
		fmt.Printf("file %s does not exist \n", filepath)
		f, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("error creating new file %s: %v", filepath, err)
		}
		f.Write(b)
	} else {
		return fmt.Errorf("error could not determine if file %s exists or not: %v", filepath, err)
	}
	return nil
}
