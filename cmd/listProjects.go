/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/linuxoid69/gitlab2gitea/internal/config"
	"github.com/linuxoid69/gitlab2gitea/internal/gitlab"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listProjectsCmd represents the listProjects command
var listProjectsCmd = &cobra.Command{
	Use:   "listProjects",
	Short: "Get list of projects from GitLab group",
	Run: func(cmd *cobra.Command, args []string) {
		config.CheckConfigFileExists()

		glab := gitlab.NewClient(viper.GetString("gitlab.url"), viper.GetString("gitlab.token"))

		groupName, err := cmd.Flags().GetString("group-name")
		if err != nil {
			log.Fatalf("Failed to get GitLab URL flag: %v", err)
		}

		projects, err := glab.ListGroupProjects(groupName)
		if err != nil {
			log.Fatalf("Failed to list projects from GitLab group: %v", err)
		}

		for _, p := range projects {
			fmt.Println(p.Name)
		}
	},
}

func init() {
	gitlabCmd.AddCommand(listProjectsCmd)

	listProjectsCmd.Flags().StringP("group-name", "g", "", "GitLab group name")
}
