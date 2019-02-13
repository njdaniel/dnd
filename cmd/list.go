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
	"context"
	"fmt"
	"github.com/njdaniel/dnd/util/list"
	"github.com/spf13/cobra"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("list called")
		List(args, context.Background())
	},
}

var client list.FilesClient

func init() {
	dataCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// List prints out all resources
func List(args []string, ctx context.Context) error {
	path := ""
	if len(args) != 0 {
		//fmt.Println("This is the args" + args[0])
		path = "/" + args[0]
	}
	l, err := client.List(ctx, &list.Path{})
	if err != nil {
		return fmt.Errorf("could not fetch data %v", err)
	}
	return nil

}
