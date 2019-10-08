/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long:  `A longer description of delete command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ls called")
		name, _ := cmd.Flags().GetString("name")
		age, _ := cmd.Flags().GetInt("age")
		if name == "" {
			name = "World"
		}
		if viper.GetString("name") != "" {
			name = viper.GetString("name")
		}
		if viper.GetString("age") != "" {
			age = viper.GetInt("age")
		}

		fmt.Printf("Name: %s, Age: %d\n", name, age)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().StringP("name", "n", "", "Give your name")
	lsCmd.Flags().Int("age", 30, "Give your age.")
}
