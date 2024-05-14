package gitea

import (
	"fmt"

	git "code.gitea.io/sdk/gitea"
)

const (
	GITEA_FAIL_LOGIN = "Can't login to Gitea"
)

type Client struct {
	Host  string
	Token string
}

// NewClient - new Gilab client.
func NewClient(host, token string) *Client {
	return &Client{
		Host:  host,
		Token: token,
	}
}

// login - function login for Gitlab.
func (c *Client) login() (*git.Client, error) {
	client, err := git.NewClient(c.Host, git.SetToken(c.Token))
	if err != nil {
		return nil, fmt.Errorf("Error login fail: %v", err)
	}

	return client, nil
}

type IsExistsProjectInOrgOpt struct {
	OrgName     string
	ProjectName string
}

// IsExistsProjectInOrg checks if a project exists in an org on Gitea
func (c *Client) IsExistsProjectInOrg(opt *IsExistsProjectInOrgOpt) (bool, error) {
	client, err := c.login()
	if err != nil {
		return false, fmt.Errorf("%s %w", GITEA_FAIL_LOGIN, err)
	}

	repos, _, err := client.ListOrgRepos(opt.OrgName, git.ListOrgReposOptions{
		ListOptions: git.ListOptions{
			Page:     -1,
			PageSize: 0,
		},
	})
	if err != nil {
		return false, fmt.Errorf("Can't check exists repo or not %w", err)
	}

	for _, repo := range repos {
		if repo.Name == opt.ProjectName {
			return true, nil
		}
	}

	return false, nil
}

type MigrateRepoOpt struct {
	OrgName     string
	ProjectName string
	CloneAddr   string
	GitlabToken string
}

func (c *Client) MigrateRepo(opt *MigrateRepoOpt) error {
	client, err := c.login()
	if err != nil {
		return fmt.Errorf("%s %w", GITEA_FAIL_LOGIN, err)
	}
	_, _, err = client.MigrateRepo(git.MigrateRepoOption{
		RepoName:     opt.ProjectName,
		RepoOwner:    opt.OrgName,
		CloneAddr:    opt.CloneAddr,
		Service:      git.GitServiceGitlab,
		AuthToken:    opt.GitlabToken,
		Private:      true,
		Issues:       true,
		PullRequests: true,
		Milestones:   true,
	})
	if err != nil {
		return fmt.Errorf("Can't migrate repo %s/%s", opt.OrgName, opt.ProjectName)
	}
	
	return nil
}
