package update

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v62/github"
)

func UpdateTitle(repo *github.Repository, client *github.Client, newName string) (*github.Repository, error) {
	fmt.Printf("Starting to update title of repo")
	time.Sleep(time.Millisecond * 340)

	owner := *repo.Owner.Login
	name := *repo.Name

	if owner == "" || name == "" {
		return nil, fmt.Errorf("error while reading owner and name fo repo check out :- ")
	}

	ctx := context.Background()

	repo, resp, err := client.Repositories.Edit(ctx, owner, name, &github.Repository{
		Name: github.String(newName),
	})

	if err != nil {
		return nil, fmt.Errorf("error file updating the info %v check out :- %w", resp, err)
	}

	return repo, nil

}
