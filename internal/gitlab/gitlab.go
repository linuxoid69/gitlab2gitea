package gitlab

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

type Client struct {
	Token string
	URL   string
}

// NewClient returns a new Client
func NewClient(url, token string) *Client {
	return &Client{
		Token: token,
		URL:   url,
	}
}

func (c *Client) login() (*gitlab.Client, error) {
	gitlabClient, err := gitlab.NewClient(c.Token, gitlab.WithBaseURL(fmt.Sprintf("%s/api/v4", c.URL)))
	if err != nil {
		return nil, err
	}

	return gitlabClient, err
}

// ListGroups lists all groups from GitLab
func (c *Client) ListGroups() ([]*gitlab.Group, error) {
	git, err := c.login()
	if err != nil {
		return nil, err
	}

	groups, _, err := git.Groups.ListGroups(&gitlab.ListGroupsOptions{})
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (c *Client) ListGroupProjects(group string) ([]*gitlab.Project, error) {
	git, err := c.login()
	if err != nil {
		return nil, err
	}

	groupId, err := c.getGroupID(group)
	if err != nil {
		return nil, err
	}

	projects, _, err := git.Groups.ListGroupProjects(groupId, &gitlab.ListGroupProjectsOptions{})
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (c *Client) getGroupID(group string) (int, error) {
	groups, err := c.ListGroups()
	if err != nil {
		return 0, err
	}

	for _, g := range groups {
		if g.FullPath == group {
			return g.ID, nil
		}
	}

	return 0, nil
}
