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

		t := table.NewWriter()
		t.AppendHeader(table.Row{"Group ID", "Group Name", "Group full path", "Description"})

		t.SetColumnConfigs([]table.ColumnConfig{
			{Name: "Group ID", AlignHeader: text.AlignCenter},
			{Name: "Group Name", AlignHeader: text.AlignCenter},
			{Name: "Group full path", AlignHeader: text.AlignCenter},
			{Name: "Description", AlignHeader: text.AlignCenter},
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
