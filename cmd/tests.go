/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/sollievo/controller"
)

// testsCmd represents the tests command
var testsCmd = &cobra.Command{
	Use:   "tests",
	Short: "lib of tests",
	Run: func(cmd *cobra.Command, args []string) {
		controller.TestsController()
	},
}

func init() {
	rootCmd.AddCommand(testsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
