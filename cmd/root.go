/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "github.ylanzinhoy.tooling_golang",
	Long: `
___________           .__  .__                   ________       .__                         
\__    ___/___   ____ |  | |__| ____    ____    /  _____/  ____ |  | _____    ____    ____  
  |    | /  _ \ /  _ \|  | |  |/    \  / ___\  /   \  ___ /  _ \|  | \__  \  /    \  / ___\ 
  |    |(  <_> |  <_> )  |_|  |   |  \/ /_/  > \    \_\  (  <_> )  |__/ __ \|   |  \/ /_/  >
  |____| \____/ \____/|____/__|___|  /\___  /   \______  /\____/|____(____  /___|  /\___  / 
                                   \//_____/           \/                 \/     \//_____/  
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Version: "0.95.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github.ylanzinhoy.tooling_golang.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
