package domain

const (
	// GitHubTokenEnv ...
	GitHubTokenEnv string = "GITHUB_ACCESS_TOKEN"
	// GitHubOwner ...
	GitHubOwner string = "GITHUB_OWNER"
	// GitHubRepos ...
	GitHubRepos string = "GITHUB_REPO"
	// StagingCheckpoint ...
	StagingCheckpoint string = "master"
)

type (
	// RepoPayload ...
	RepoPayload struct {
		Name                string `json:"name"`
		IsProductionUpdated bool   `json:"isProductionUpdated"`
		URL                 string `json:"url"`
	}

	// OrganizationPayload ...
	OrganizationPayload struct {
		Owner string        `json:"owner"`
		Repos []RepoPayload `json:"repos"`
	}
)
