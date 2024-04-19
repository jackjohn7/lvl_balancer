/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init will create a basic configuration file",
	Long: `To get started with LVL, use the <init> command 
to create a basic configuration file that you can customize.

To do this, run the following command:
cli init`,
	Run: func(cmd *cobra.Command, args []string) {
		// create new file
		fmt.Println("Initializing")
		_, err := os.Create("lvl.yml")
		if err != nil {
			panic("Could not create lvl.yml file")
		}
		// write default configuration to initFile
		// initFile.WriteString()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
