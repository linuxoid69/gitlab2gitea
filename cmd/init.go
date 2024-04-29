/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/linuxoid69/gitlab2gitea/internal/config"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Run init for creating config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.NewConfig())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
