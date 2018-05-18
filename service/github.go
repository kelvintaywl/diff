package service

import (
	"context"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/kelvintaywl/diff/domain"
	"golang.org/x/oauth2"
)

type (
	// GitHubCompareResp is a response for git compare between 2 refs
	GitHubCompareResp struct {
		AheadBy  int
		BehindBy int
		DiffURL  string
		HTMLURL  string
	}

	// GitHubRepo ...
	GitHubRepo struct {
		Owner string
		Repo  string
	}

	// GitHubClient is an interface that allows us to interact with some of the Github API.
	GitHubClient interface {
		Repos(ctx context.Context) []GitHubRepo
		LatestTag(ctx context.Context, owner, repo string) (string, error)
		Compare(ctx context.Context, owner, repo, base, head string) (*GitHubCompareResp, error)
	}

	proxy struct {
		client *github.Client
	}
)

// NewGitHubClient returns a GithubClient.
func NewGitHubClient(ctx context.Context) GitHubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(domain.GitHubTokenEnv)},
	)
	tc := oauth2.NewClient(ctx, ts)
	c := github.NewClient(tc)
	return proxy{
		client: c,
	}
}

func (c proxy) Compare(ctx context.Context, owner, repo, base, head string) (*GitHubCompareResp, error) {
	cmp, _, err := c.client.Repositories.CompareCommits(ctx, owner, repo, base, head)
	if err != nil {
		return nil, err
	}
	return &GitHubCompareResp{
		AheadBy:  *cmp.AheadBy,
		BehindBy: *cmp.BehindBy,
		DiffURL:  *cmp.DiffURL,
		HTMLURL:  *cmp.HTMLURL,
	}, nil
}

func (c proxy) LatestTag(ctx context.Context, owner, repo string) (string, error) {
	rel, _, err := c.client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return "", err
	}
	return *rel.TagName, nil
}

func (c proxy) Repos(ctx context.Context) []GitHubRepo {
	owner := os.Getenv(domain.GitHubOwner)
	rs := os.Getenv(domain.GitHubRepos)
	repoNames := strings.Split(rs, ",")
	repos := make([]GitHubRepo, 0, len(repoNames))
	for _, name := range repoNames {
		repos = append(repos, GitHubRepo{
			Owner: owner,
			Repo:  name,
		})
	}
	return repos
}
