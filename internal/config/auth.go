package config

import (
	"context"
	"fmt"

	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"
)

func GetGithubClient(cfg *Config) (*github.Client, error) {
	if cfg.GitToken == "" {
		return nil, fmt.Errorf("Github token is not set, set github token using cobraclip login command ")
	}

	ctx := context.Background()
	// Create a token source that holds the GitHub token for authentication
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, nil
}
