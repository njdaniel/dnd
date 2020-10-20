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
	"fmt"
	"github.com/njdaniel/dnd/util/dice"

	"github.com/spf13/cobra"
)

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "Rolls dice",
	Long: `Dice Calculator. Can do basic rolls and more complex calculations.

	Examples:
	ex1.
	$ dnd roll d20
	$ 13

	ex2.
	$ dnd roll 2d6
	$ [4,1]

	ex3.
	$ 4kh3d6
	$ [6,5,3]

	ex4.
	$ 4kl3d6
	$ [1,1,4]

	ex5.
	$ 6!
	$ 13
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("roll called")
		d := dice.ParseRollString(args[0])
		fmt.Println(d.RollDice())
	},
}

func init() {
	rootCmd.AddCommand(rollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
