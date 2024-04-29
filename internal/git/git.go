package git

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const (
	GIT_MIGRATE_TEMP_DIR        = "/tmp/gitlab2gitea"
	GIT_GITLAB_DEFAULT_USERNAME = "gitlab-ci-token"
)

type RepoOpt struct {
	URL      string
	Username string
	Token    string
	Project  string
	TemDir   string
}

func CloneByHTTPS(opt *RepoOpt) error {
	repoPath := fmt.Sprintf("%s/%s", opt.URL, opt.Project)
	_, err := git.PlainClone(fmt.Sprintf("%s/%s", opt.TemDir, opt.Project), true, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: opt.Username,
			Password: opt.Token,
		},
		URL:      repoPath,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}
