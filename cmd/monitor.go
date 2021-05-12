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
	"time"

	"github.com/souvikhaldar/cw/pkg/market_ticker"
	"github.com/souvikhaldar/cw/pkg/typedef"
	"github.com/spf13/cobra"
)

var upper int
var lower int

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring: ", cryptoName)
		mt := market_ticker.NewWazir()

		c := typedef.ConvertLingo(cryptoName)

		for {
			price, err := market_ticker.GetPrice(
				mt,
				c,
			)
			if err != nil {
				//fmt.Println("Error in getting price: ", err)
				continue
			}
			time.Sleep(10 * time.Second)

			if lower == 0 && upper == 999999999 {
				fmt.Printf("Price of %s is:%f\n", cryptoName, price)
				continue
			}
			if price >= float64(upper) {
				fmt.Println("Exceeded target")
				fmt.Printf("Price of %s is:%f\n ", c, price)
				fmt.Println("Time: ", time.Now())
			} else if price <= float64(lower) {
				fmt.Println("Price below lower limit!")
				fmt.Printf("Price of %s is:%f\n ", c, price)
				fmt.Println("Time: ", time.Now())
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	rootCmd.PersistentFlags().IntVarP(&upper, "upper", "u", 999999999, "Upper limit or target")
	rootCmd.PersistentFlags().IntVarP(&lower, "lower", "l", 0, "Lower limit")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
