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
	"net/http"

	"github.com/spf13/cobra"
)

var fetchURL string = "/api/v2/tickers"

type ticker struct {
	At        int64       `json:"at"`
	BaseUnit  string      `json:"base_unit"`
	Buy       string      `json:"buy"`
	High      string      `json:"high"`
	Last      string      `json:"last"`
	Low       string      `json:"low"`
	Name      string      `json:"name"`
	Open      interface{} `json:"open"`
	QuoteUnit string      `json:"quote_unit"`
	Sell      string      `json:"sell"`
	Type      string      `json:"type"`
	Volume    string      `json:"volume"`
}

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching info of ", cryptoName)

		client := &http.Client{}
		req, err := http.NewRequest(
			"GET",
			baseURL+fetchURL,
			nil,
		)
		if err != nil {
			fmt.Println("Unable to create the request: ", err)
			return
		}
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		defer resp.Body.Close()

		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	fmt.Println("Error in reading the resp body: ", err)
		//	return
		//}
		var m map[string]ticker
		if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
			fmt.Println("Error in decoding the JSON: ", err)
			return
		}

		fmt.Printf("Current price of %s : %s\n", cryptoName, m[cryptoName].Last)
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
