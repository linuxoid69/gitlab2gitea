/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/linuxoid69/gitlab2gitea/internal/config"
	"github.com/linuxoid69/gitlab2gitea/internal/flags"
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

		groupName := flags.CheckFlag(cmd, "group-name")

		glab := gitlab.NewClient(viper.GetString("gitlab.url"), viper.GetString("gitlab.token"))

		projects, err := glab.ListGroupProjects(groupName)
		if err != nil {
			log.Fatalf("Failed to list projects from GitLab group: %v", err)
		}

		t := table.NewWriter()
		tableProjectID := "Project ID"
		tableProjectName := "Project Name"
		tableProjectFullPath := "Project full path"
		tableProjectDesctiption := "Description"

		t.SetStyle(table.StyleBold)
		t.Style().Color.Header = text.Colors{text.FgHiCyan}
		t.Style().Color.Border = text.Colors{text.FgHiCyan}
		t.Style().Color.Row = text.Colors{text.FgHiCyan}
		t.Style().Color.RowAlternate = text.Colors{text.FgCyan}

		t.AppendHeader(table.Row{tableProjectID, tableProjectName, tableProjectFullPath, tableProjectDesctiption})

		t.SetColumnConfigs([]table.ColumnConfig{
			{Name: tableProjectID, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
			{Name: tableProjectName, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
			{Name: tableProjectFullPath, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
			{Name: tableProjectDesctiption, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
		})

		for _, p := range projects {
			t.AppendRow(table.Row{p.ID, p.Path, p.PathWithNamespace, p.Description})
		}

		fmt.Println(t.Render())

	},
}

func init() {
	gitlabCmd.AddCommand(listProjectsCmd)

	listProjectsCmd.Flags().StringP("group-name", "g", "", "GitLab group name")
}
