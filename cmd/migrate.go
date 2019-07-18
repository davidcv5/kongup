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
	"log"

	"github.com/davidcv5/kongup/kongdeck"

	"github.com/davidcv5/kongup/kongfig"
	"github.com/spf13/cobra"
)

var kongfigIn string
var deckOut string
var tags []string

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		k, err := kongfig.GetKongfigFromFile(kongfigIn)
		if err != nil {
			log.Fatal(err)
		}
		err = kongdeck.KongfigToDeck(k, tags, deckOut)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().StringVarP(&kongfigIn,
		"kongfig", "i", "kongfig.yaml", "file containing Kong's configuration "+
			"in kongfig format. "+
			"Use '-' to read from stdin.")

	migrateCmd.Flags().StringVarP(&deckOut, "output-file", "o",
		"deck.yaml", "write Kong configuration to FILE in deck format."+
			"Use '-' to write to stdout.")

	migrateCmd.Flags().StringSliceVar(&tags,
		"tag", []string{},
		"adds a tag to core resources.")
}
