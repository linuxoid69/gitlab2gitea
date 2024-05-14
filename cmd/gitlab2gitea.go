/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/linuxoid69/gitlab2gitea/internal/config"
	"github.com/linuxoid69/gitlab2gitea/internal/flags"
	"github.com/linuxoid69/gitlab2gitea/internal/gitea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gitlab2giteaCmd represents the gitlab2gitea command
var gitlab2giteaCmd = &cobra.Command{
	Use:   "gitlab2gitea",
	Short: "Migrate GitLab repo to Gitea",
	Run: func(cmd *cobra.Command, args []string) {
		config.CheckConfigFileExists()

		gitlabProject := flags.CheckFlag(cmd, "gitlab-project")
		giteaOrg := flags.CheckFlag(cmd, "gitea-org")
		giteaProject := flags.CheckFlag(cmd, "gitea-project")

		giteaClient := gitea.NewClient(viper.GetString("gitea.url"), viper.GetString("gitea.token"))
		
		projectExists, err := giteaClient.IsExistsProjectInOrg(&gitea.IsExistsProjectInOrgOpt{
			OrgName:     giteaOrg,
			ProjectName: giteaProject,
		})
		if err != nil {
			fmt.Println("Error migrating repo: ", err)
			os.Exit(1)
		}

		if projectExists {
			fmt.Printf("Project `%s` already exists in `%s` - SKIP\n", giteaProject, giteaOrg)
			os.Exit(1)
		} else {
			fmt.Printf("Run migration from GitLab project `%s` to Gitea project `%s/%s`\n", gitlabProject, giteaOrg, giteaProject)
		}

		if err := giteaClient.MigrateRepo(&gitea.MigrateRepoOpt{
			OrgName:     giteaOrg,
			ProjectName: giteaProject,
			GitlabToken: viper.GetString("gitlab.token"),
			CloneAddr:   viper.GetString("gitlab.url") + "/" + gitlabProject,

		}); err != nil {
			fmt.Println("Error migrating repo: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	migrateCmd.AddCommand(gitlab2giteaCmd)
	gitlab2giteaCmd.Flags().StringP("gitlab-project", "p", "", "Gitlab project name")
	gitlab2giteaCmd.Flags().StringP("gitea-project", "P", "", "Gitea project name")
	gitlab2giteaCmd.Flags().StringP("gitea-org", "o", "", "Gitea organization name")
}
