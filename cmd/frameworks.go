/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.ylanzinhoy.tooling_golang/controller"
)

// frameworksCmd represents the frameworks command
var frameworksCmd = &cobra.Command{
	Use:   "frameworks",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.ToolingControllerUpper{}
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
