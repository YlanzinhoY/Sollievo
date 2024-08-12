/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ylanzinhoy/sollievo/controller"
)

// toolsCmd represents the tools command
var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "General tools for Go",
	Run: func(cmd *cobra.Command, args []string) {
		controller.ToolingController()
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toolsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toolsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
