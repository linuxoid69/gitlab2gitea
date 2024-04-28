/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Gitlab commands",
	Long:  `Commands for work with Gitlab`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(gitlabCmd)

	gitlabCmd.PersistentFlags().StringP("gitlab-url", "u", "", "GitLab URL")
	gitlabCmd.PersistentFlags().StringP("gitlab-token", "t", "", "Gitlab token")
}
