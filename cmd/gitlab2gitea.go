/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitlab2giteaCmd represents the gitlab2gitea command
var gitlab2giteaCmd = &cobra.Command{
	Use:   "gitlab2gitea",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitlab2gitea called")
	},
}

func init() {
	migrateCmd.AddCommand(gitlab2giteaCmd)
}
