/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/linuxoid69/gitlab2gitea/internal/config"
	"github.com/linuxoid69/gitlab2gitea/internal/gitlab"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listGroupsCmd represents the listGroups command
var listGroupsCmd = &cobra.Command{
	Use:   "listGroups",
	Short: "Get list of groups from GitLab",
	Run: func(cmd *cobra.Command, args []string) {
		
		config.CheckConfigFileExists()
		
		glab := gitlab.NewClient(viper.GetString("gitlab.url"), viper.GetString("gitlab.token"))

		groups, err := glab.ListGroups()
		if err != nil {
			fmt.Println("Error getting groups: ", err)
			os.Exit(1)
		}

		for _, g := range groups {
			fmt.Println(g.FullName)
		}
	},
}

func init() {
	gitlabCmd.AddCommand(listGroupsCmd)
}
