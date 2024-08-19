/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/sollievo/controller"
)

// frameworksCmd represents the frameworks command
var frameworksCmd = &cobra.Command{
	Use:   "frameworks",
	Short: "Frameworks web for Go",
	Run: func(cmd *cobra.Command, args []string) {
		controller.FrameworkController()
	},
}

func init() {
	rootCmd.AddCommand(frameworksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// frameworksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// frameworksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
