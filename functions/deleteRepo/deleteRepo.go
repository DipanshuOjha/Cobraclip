package deleterepo

import (
	"context"
	"fmt"

	"github.com/google/go-github/v62/github"
)

func DeleteRepo(repo *github.Repository, client *github.Client) error {

	if repo.Owner == nil || repo.Owner.Login == nil {
		return fmt.Errorf("repository owner is nil")
	}
	if repo.Name == nil {
		return fmt.Errorf("repository name is nil")
	}

	owner := *repo.Owner.Login
	reponame := *repo.Name

	ctx := context.Background()

	resp, err := client.Repositories.Delete(ctx, owner, reponame)

	if err != nil {
		return fmt.Errorf("error while deleting the repo check out the response %v and error %v", resp, err)
	}

	return nil
}
