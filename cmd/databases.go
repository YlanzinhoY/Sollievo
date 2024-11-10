/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ylanzinhoy/sollievo/controller"

	"github.com/spf13/cobra"
)

// databasesCmd represents the databasesCmd command
var databasesCmd = &cobra.Command{
	Use:   "databases",
	Short: "drivers of databases",
	Run: func(cmd *cobra.Command, args []string) {
		controller.SqlDriversController()
	},
}

func init() {
	rootCmd.AddCommand(databasesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sqlDriversCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sqlDriversCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
