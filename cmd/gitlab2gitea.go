/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/linuxoid69/gitlab2gitea/internal/config"
	"github.com/linuxoid69/gitlab2gitea/internal/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gitlab2giteaCmd represents the gitlab2gitea command
var gitlab2giteaCmd = &cobra.Command{
	Use:   "gitlab2gitea",
	Short: "Migrate GitLab repo to Gitea",
	Run: func(cmd *cobra.Command, args []string) {
		config.CheckConfigFileExists()

		project, err := cmd.Flags().GetString("gitlab-project")
		if err != nil || project == "" {
			fmt.Println("Failed to get flag `gitlab-project`")
			os.Exit(1)
		}

		if err := git.CloneByHTTPS(&git.RepoOpt{
			URL:      viper.GetString("gitlab.url"),
			Username: git.GIT_GITLAB_DEFAULT_USERNAME,
			Token:    viper.GetString("gitlab.token"),
			TemDir:   git.GIT_MIGRATE_TEMP_DIR,
			Project:  project,
		}); err != nil {
			fmt.Println("Error cloning repo: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	migrateCmd.AddCommand(gitlab2giteaCmd)
	gitlab2giteaCmd.Flags().StringP("gitlab-project", "p", "", "Gitlab project")
}
