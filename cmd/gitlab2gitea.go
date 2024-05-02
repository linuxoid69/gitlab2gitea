/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/linuxoid69/gitlab2gitea/internal/config"
	"github.com/linuxoid69/gitlab2gitea/internal/flags"
	"github.com/linuxoid69/gitlab2gitea/internal/git"
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
		_ = flags.CheckFlag(cmd, "gitea-project")

		if err := git.CloneByHTTPS(&git.RepoOpt{
			URL:      viper.GetString("gitlab.url"),
			Username: git.GIT_GITLAB_DEFAULT_USERNAME,
			Token:    viper.GetString("gitlab.token"),
			TemDir:   git.GIT_MIGRATE_TEMP_DIR,
			Project:  gitlabProject,
		}); err != nil {
			fmt.Println("Error cloning repo: ", err)
			os.Exit(1)
		}

		giteaClient := gitea.NewClient(viper.GetString("gitea.url"), viper.GetString("gitea.token"))
		fmt.Println(giteaClient.GetCurrentUser())
		// Check if repo already exists on Gitea
		// if exists, print skipt message and continue
		// if not then create new repo on Gitea
	},
}

func init() {
	migrateCmd.AddCommand(gitlab2giteaCmd)
	gitlab2giteaCmd.Flags().StringP("gitlab-project", "p", "", "Gitlab project")
	gitlab2giteaCmd.Flags().StringP("gitea-project", "P", "", "Gitea project")
}
