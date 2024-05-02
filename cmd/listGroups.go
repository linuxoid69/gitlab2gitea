/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

		// TODO: move to separate function

		t := table.NewWriter()
		tableGroupID := "Group ID"
		tableGroupName := "Group Name"
		tableGroupFullPath := "Group full path"
		tableGroupDesctiption := "Description"

		t.SetStyle(table.StyleBold)
		t.Style().Color.Header = text.Colors{text.FgHiCyan}
		t.Style().Color.Border = text.Colors{text.FgHiCyan}
		t.Style().Color.Row = text.Colors{text.FgHiCyan}
		t.Style().Color.RowAlternate = text.Colors{text.FgCyan}

		t.AppendHeader(table.Row{tableGroupID, tableGroupName, tableGroupFullPath, tableGroupDesctiption})

		t.SetColumnConfigs([]table.ColumnConfig{
			{Name: tableGroupID, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
			{Name: tableGroupName, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
			{Name: tableGroupFullPath, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
			{Name: tableGroupDesctiption, AlignHeader: text.AlignCenter, Colors: text.Colors{text.FgHiCyan}},
		})

		for _, g := range groups {
			t.AppendRow(table.Row{g.ID, g.Name, g.FullPath, g.Description})
		}

		fmt.Println(t.Render())
	},
}

func init() {
	gitlabCmd.AddCommand(listGroupsCmd)
}
