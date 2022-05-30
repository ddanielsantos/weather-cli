/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/ddanielsantos/weather-cli/app"
	"github.com/spf13/cobra"
)

var city string

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "now",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.PrintWeather(city)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	weatherCmd.Flags().StringVarP(&city, "city", "c", "", "city name")

}
