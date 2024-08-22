/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backendCmd represents the backend command
var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Creating backend Application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backend called")
	},
}

func init() {
	rootCmd.AddCommand(backendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
