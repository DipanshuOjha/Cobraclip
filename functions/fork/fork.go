package fork

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v62/github"
)

type ForkStage string

const (
	StageStarting       ForkStage = "Starting fork..."
	StageCheckingExists ForkStage = "Checking if fork already exists..."
	StageCreating       ForkStage = "Creating fork..."
	StageVerifying      ForkStage = "Verifying content..."
	StageDone           ForkStage = "Done!"
	StageFailed         ForkStage = "Failed!"
)

// ForkProgress tracks the progress of a fork operation
type ForkProgress struct {
	Stage ForkStage
	Repo  *github.Repository
	Error error
	Tips  string
}

func ForkRepo(repo *github.Repository, client *github.Client, updates chan<- ForkProgress) (*github.Repository, error) {
	defer close(updates)

	// Validate inputs
	if repo == nil || client == nil {
		return nil, fmt.Errorf("nil repository or client")
	}

	owner := repo.GetOwner().GetLogin()
	repoName := repo.GetName()
	if owner == "" || repoName == "" {
		return nil, fmt.Errorf("invalid repository details")
	}

	ctx := context.Background()
	currentUser, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}
	yourUsername := currentUser.GetLogin()

	// STAGE 1: Check for existing fork
	updates <- ForkProgress{Stage: StageCheckingExists}
	existing, err := getExistingFork(ctx, client, yourUsername, repoName)
	if err != nil {
		return nil, fmt.Errorf("error checking for existing fork: %w", err)
	}

	// If fork exists, return it immediately
	if existing != nil {
		updates <- ForkProgress{
			Stage: StageDone,
			Repo:  existing,
			Tips:  "Using existing fork",
		}
		return existing, nil
	}

	// STAGE 2: Create new fork
	updates <- ForkProgress{Stage: StageCreating}
	fork, resp, err := client.Repositories.CreateFork(ctx, owner, repoName, nil)

	// Handle GitHub's async processing
	if resp != nil && resp.StatusCode == http.StatusAccepted {
		return handleAsyncFork(ctx, client, yourUsername, repoName, updates)
	}

	if err != nil {
		// Final check in case fork was created but we got an error
		if existing, err := getExistingFork(ctx, client, yourUsername, repoName); err == nil && existing != nil {
			updates <- ForkProgress{Stage: StageDone, Repo: existing}
			return existing, nil
		}
		return nil, fmt.Errorf("failed to create fork: %w", err)
	}

	updates <- ForkProgress{Stage: StageDone, Repo: fork}
	return fork, nil

}

func handleAsyncFork(ctx context.Context, client *github.Client, owner, repo string, updates chan<- ForkProgress) (*github.Repository, error) {
	// GitHub is processing the fork asynchronously
	updates <- ForkProgress{
		Stage: StageCreating,
		Tips:  "GitHub is processing your fork in the background. This may take a few minutes...",
	}

	// Check every 5 seconds for up to 10 minutes
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	timeout := time.After(10 * time.Minute)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-timeout:
			return nil, fmt.Errorf("fork verification timeout")
		case <-ticker.C:
			fork, err := getExistingFork(ctx, client, owner, repo)
			if err != nil {
				continue
			}
			if fork != nil {
				// Verify the fork is fully ready
				if _, _, err := client.Repositories.GetBranch(ctx, owner, repo, fork.GetDefaultBranch(), 0); err == nil {
					updates <- ForkProgress{Stage: StageDone, Repo: fork, Tips: "GitHub has finished creating your fork"}
					return fork, nil
				}
			}
		}
	}
}

func getExistingFork(ctx context.Context, client *github.Client, owner, repo string) (*github.Repository, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	fork, resp, err := client.Repositories.Get(ctx, owner, repo)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	// Only return if it's actually a fork
	if fork != nil && fork.GetFork() {
		return fork, nil
	}
	return nil, nil

	// Only return if it's actually a fork
}
func createRepositoryFork(ctx context.Context, client *github.Client, owner, repo string) (*github.Repository, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	fork, resp, err := client.Repositories.CreateFork(ctx, owner, repo, nil)
	if err != nil {
		// If fork is being processed asynchronously (status 202)
		if resp != nil && resp.StatusCode == 202 {
			return waitForForkCreation(ctx, client, owner, repo)
		}
		return nil, err
	}
	return fork, nil
}

func verifyRepositoryFork(ctx context.Context, client *github.Client, fork *github.Repository) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	// Check default branch exists
	_, _, err := client.Repositories.GetBranch(ctx, fork.GetOwner().GetLogin(), fork.GetName(), fork.GetDefaultBranch(), 3)
	return err
}

func waitForForkCreation(ctx context.Context, client *github.Client, owner, repo string) (*github.Repository, error) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			fork, err := getExistingFork(ctx, client, owner, repo)
			if err != nil {
				return nil, err
			}
			if fork != nil {
				return fork, nil
			}
		}
	}
}
