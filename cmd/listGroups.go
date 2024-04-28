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

// listGroupsCmd represents the listGroups command
var listGroupsCmd = &cobra.Command{
	Use:   "listGroups",
	Short: "Get list of groups from GitLab",
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
