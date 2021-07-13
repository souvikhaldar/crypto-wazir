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
	"fmt"

	"github.com/souvikhaldar/cw/pkg/market_ticker"
	"github.com/souvikhaldar/cw/pkg/typedef"
	"github.com/spf13/cobra"
)

var fetchURL string = "/api/v2/tickers"

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches all the details of the cypto at any time",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching details of:", cryptoName)
		c := typedef.ConvertLingo(cryptoName)

		mt := market_ticker.NewWazir()
		details, err := market_ticker.GetAllDetails(
			mt,
			c,
		)
		if err != nil || details == nil {
			fmt.Println("Error in getting price: ", err)
			return
		}
		fmt.Printf("Details of %s : %+v\n", cryptoName, *details)
		return

	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
