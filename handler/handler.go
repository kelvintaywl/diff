package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kelvintaywl/diff/domain"
	"github.com/kelvintaywl/diff/service"
)

// DiffListHandler ...
func DiffListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	gc := service.NewGitHubClient(ctx)
	var repos []domain.RepoPayload
	org := domain.OrganizationPayload{}

	for _, repo := range gc.Repos(ctx) {
		if org.Owner == "" {
			org.Owner = repo.Owner
		}

		tag, err := gc.LatestTag(ctx, repo.Owner, repo.Repo)
		if err != nil {
			continue
		}

		cmp, err := gc.Compare(ctx, repo.Owner, repo.Repo, tag, domain.StagingCheckpoint)
		if err != nil {
			continue
		}
		repo := domain.RepoPayload{
			IsProductionUpdated: cmp.AheadBy <= 0,
			URL:                 cmp.HTMLURL,
		}
		repos = append(repos, repo)
	}
	org.Repos = repos

	body, err := json.Marshal(org)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
