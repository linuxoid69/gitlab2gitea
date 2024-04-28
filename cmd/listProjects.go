/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/linuxoid69/gitlab2gitea/internal/gitlab"
	"github.com/spf13/cobra"
)

// listProjectsCmd represents the listProjects command
var listProjectsCmd = &cobra.Command{
	Use:   "listProjects",
	Short: "Get list of projects from GitLab group",
	Run: func(cmd *cobra.Command, args []string) {
		gitlabToken, err := cmd.Flags().GetString("gitlab-token")
		if err != nil {
			log.Fatalf("Failed to get GitLab token flag: %v", err)
		}
	
		gitlabURL, err := cmd.Flags().GetString("gitlab-url")
		if err != nil {
			log.Fatalf("Failed to get GitLab URL flag: %v", err)
		}
	
		if gitlabToken == "" || gitlabURL == "" {
			cmd.Help()
			os.Exit(0)
		}

		glab := gitlab.NewClient(gitlabToken, gitlabURL)

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
