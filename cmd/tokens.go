/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tokensCmd represents the tokens command
var tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getTokens()
	},
}

func init() {
	rootCmd.AddCommand(tokensCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokensCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokensCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getTokens() {
	viper.SetConfigName("tokens") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	topLevelKeys := viper.AllSettings()

	// Loop through the keys and print them
	color.Cyan("%-10s %-10s %-10s %-10s %-10s \n", "account", "type", "pass", "plus", "shared")

	for key := range topLevelKeys {
		var token = viper.Get(key + ".token")
		if token == nil {
			continue
		}
		var types = "access"
		var pass = viper.Get(key+".password") == nil
		var plus = viper.Get(key+".plus") == nil
		var shared = viper.Get(key+".shared") == nil
		fmt.Printf("%-10s %-10s %-10t %-10t %-10t \n", key, types, pass, plus, shared)
	}
}