package update

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v62/github"
)

func UpdateDescriptionandTitle(repo *github.Repository, client *github.Client, newName string, newDecrip string) (*github.Repository, error) {
	fmt.Printf("Starting to update title of repo")
	time.Sleep(time.Millisecond * 340)

	owner := *repo.Owner.Login
	name := *repo.Name

	if owner == "" || name == "" {
		return nil, fmt.Errorf("error while reading owner and name fo repo check out :- ")
	}

	ctx := context.Background()

	if newName == "" {
		repo, resp, err := client.Repositories.Edit(ctx, owner, name, &github.Repository{
			Description: github.String(newDecrip),
		})
		if err != nil {
			return nil, fmt.Errorf("error file updating the info %v check out :- %w", resp, err)
		}

		return repo, nil
	}

	repo, resp, err := client.Repositories.Edit(ctx, owner, name, &github.Repository{
		Name:        github.String(newName),
		Description: github.String(newDecrip),
	})
	if err != nil {
		return nil, fmt.Errorf("error file updating the info %v check out :- %w", resp, err)
	}

	return repo, nil

}
