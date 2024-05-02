/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"github.com/linuxoid69/gitlab2gitea/internal/flags"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate repositories",
	Long: `Migrate repositories from Gitlab to Gitea`,
	Run: func(cmd *cobra.Command, args []string) {
		flags.CheckArgs(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
