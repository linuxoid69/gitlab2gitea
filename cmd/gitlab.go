/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"github.com/linuxoid69/gitlab2gitea/internal/flags"
	"github.com/spf13/cobra"
)

// gitlabCmd represents the gitlab command
var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Gitlab commands",
	Long:  `Commands for work with Gitlab`,
	Run: func(cmd *cobra.Command, args []string) {
		flags.CheckArgs(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(gitlabCmd)
}
