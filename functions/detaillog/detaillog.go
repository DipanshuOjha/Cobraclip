package detaillog

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/DipanshuOjha/cobraclip/functions/fork"
	"github.com/google/go-github/v62/github"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var repoMenu = []struct {
	key         string
	description string
	handler     func(*github.Repository, *github.Client)
}{
	{"1", "Language used", showLanguages},
	{"2", "Key files", showFiles},
	{"3", "Branches", showBranches},
	{"4", "Recent Activity", showRecentActivity},
	{"5", "Open Issues", showOpenIssues},
	{"6", "Closed Issues", showClosedIssues},
	{"7", "Pull Requests", showPullRequests},
	{"8", "Contributors", showContributors},
	{"9", "Fork", forkit},
	{"q", "Quit", nil},
}

// type ForkStage string

// type ForkProgress struct {
// 	Stage ForkStage
// 	Repo  *github.Repository
// 	Error error
// 	Tips  string // Helpful tips for long waits
// }

// const (
// 	StageStarting       ForkStage = "Starting fork..."
// 	StageCheckingExists ForkStage = "Checking if fork already exists..."
// 	StageCreating       ForkStage = "Creating fork (this may take a few minutes)..."
// 	StageVerifying      ForkStage = "Verifying fork content..."
// 	StageDone           ForkStage = "Fork ready! 🎉"
// 	StageFailed         ForkStage = "Fork failed ❌"
// )

func forkit(repo *github.Repository, client *github.Client) {
	updates := make(chan fork.ForkProgress, 10)

	go func() {
		_, err := fork.ForkRepo(repo, client, updates)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// Display updates in real-time
	for progress := range updates {
		clearScreen() // Optional: Clears console for cleaner output

		fmt.Printf("\n🔹 %s\n", progress.Stage)
		if progress.Tips != "" {
			fmt.Printf("   💡 %s\n", progress.Tips)
		}

		for progress := range updates {
			// Use the package's stage constants
			switch progress.Stage {
			case fork.StageStarting:
				fmt.Println("Starting fork operation...")
			case fork.StageCheckingExists:
				fmt.Println("Checking for existing fork...")
			case fork.StageCreating:
				fmt.Println("Creating fork...")
			case fork.StageVerifying:
				fmt.Println("Verifying content...")
			case fork.StageDone:
				fmt.Printf("✅ Fork ready: %s\n", progress.Repo.GetHTMLURL())
			case fork.StageFailed:
				fmt.Printf("❌ Error: %v\n", progress.Error)
			}
		}
	}

}

func displayrepodetails(repo *github.Repository, client *github.Client, index int, reader *bufio.Reader) {
	items := repoMenu[index]
	clearScreen()
	showBasicInfo(repo)
	fmt.Printf("\n=== %s ===\n", items.description)
	items.handler(repo, client)
	fmt.Print("\nPress Enter to continue...")
	reader.ReadString('\n')
}

func showBasicInfo(repo *github.Repository) {
	fmt.Printf("\n\033[1m%s/%s\033[0m\n", *repo.Owner.Login, *repo.Name)
	fmt.Printf("📝 %s\n", safeString(repo.Description))
	fmt.Printf("🌐 %s\n", *repo.HTMLURL)
	fmt.Printf("⭐ Stars: %d | 🍴 Forks: %d | 👀 Watchers: %d\n",
		*repo.StargazersCount, *repo.ForksCount, *repo.WatchersCount)
}

func ShowRepoDetail(repo *github.Repository, client *github.Client) {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		showBasicInfo(repo)

		fmt.Printf("What would you like to check\n")

		for _, item := range repoMenu {
			fmt.Printf("[%s] %s\n", item.key, item.description)
		}

		fmt.Println("\n Choose any : ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		if choice == "q" {
			break
		}

		//if choice != "1" || choice != "2" || choice != "3" || choice != "4" || choice != "5" || choice != "6" || choice != "7" || choice != "8" || choice!= "9"

		// for _, items := range repoMenu {
		// 	if items.key == choice {
		// 		clearScreen()
		// 		showBasicInfo(repo)
		// 		fmt.Printf("\n=== %s ===\n", items.description)
		// 		items.handler(repo, client)
		// 		fmt.Print("\nPress Enter to continue...")
		// 		reader.ReadString('\n')
		// 		break
		// 	}
		// }

		switch choice {
		case "1":
			displayrepodetails(repo, client, 0, reader)
		case "2":
			displayrepodetails(repo, client, 1, reader)
		case "3":
			displayrepodetails(repo, client, 2, reader)
		case "4":
			displayrepodetails(repo, client, 3, reader)
		case "5":
			displayrepodetails(repo, client, 4, reader)
		case "6":
			displayrepodetails(repo, client, 5, reader)
		case "7":
			displayrepodetails(repo, client, 6, reader)
		case "8":
			displayrepodetails(repo, client, 7, reader)
		case "9":
			displayrepodetails(repo, client, 8, reader)
		default:
			fmt.Println("Wrong input please retype correctly")
		}
	}
}

func showLanguages(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name

	ctx := context.Background()

	langs, _, err := client.Repositories.ListLanguages(ctx, owner, name)
	if err != nil {
		fmt.Println("failed to list languages:")
		return
	}
	fmt.Println("\n\033[1mLanguages:\033[0m")
	for lang := range langs {
		fmt.Printf("- %s\n", lang)
		time.Sleep(time.Millisecond * 640)
	}

}

func showFiles(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name

	ctx := context.Background()

	_, content, _, _ := client.Repositories.GetContents(ctx, owner, name, "", nil)
	fmt.Println("\n\033[1mKey Files:\033[0m")
	for _, file := range content {
		fmt.Printf("- %s\n", *file.Path)
		time.Sleep(time.Millisecond * 640)
	}
}

func showBranches(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name

	ctx := context.Background()

	branches, _, _ := client.Repositories.ListBranches(ctx, owner, name, nil)
	fmt.Println("\n\033[1mBranches:\033[0m")
	for _, branch := range branches {
		fmt.Printf("- %s\n", *branch.Name)
		time.Sleep(time.Millisecond * 640)
	}
}

func showRecentActivity(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name

	ctx := context.Background()

	commits, _, _ := client.Repositories.ListCommits(ctx, owner, name, &github.CommitsListOptions{ListOptions: github.ListOptions{PerPage: 10}})
	for _, commit := range commits {
		fmt.Printf("- %s: %.50s\n",
			(*commit.SHA)[:7],
			strings.ReplaceAll(*commit.Commit.Message, "\n", " "))
		time.Sleep(time.Millisecond * 640)
	}

}

func showOpenIssues(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name

	ctx := context.Background()

	issues, _, _ := client.Issues.ListByRepo(ctx, owner, name, &github.IssueListByRepoOptions{State: "open"})
	fmt.Printf("\n\033[1mOpen Issues (%d):\033[0m\n", len(issues))
	for _, issue := range issues {
		fmt.Printf("#%d: %s\n", *issue.Number, *issue.Title)
		fmt.Printf("   Created: %s | Comments: %d\n",
			issue.CreatedAt.Format("2006-01-02"),
			*issue.Comments)
		fmt.Println("   URL:", *issue.HTMLURL)
		time.Sleep(time.Millisecond * 640)
	}
}

func showClosedIssues(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name

	ctx := context.Background()

	issues, _, _ := client.Issues.ListByRepo(ctx, owner, name, &github.IssueListByRepoOptions{State: "closed"})

	fmt.Printf("\n\033[1mOpen Issues (%d):\033[0m\n", len(issues))
	for _, issue := range issues {
		fmt.Printf("#%d: %s\n", *issue.Number, *issue.Title)
		fmt.Printf("   Created: %s | Comments: %d\n",
			issue.CreatedAt.Format("2006-01-02"),
			*issue.Comments)
		fmt.Println("   URL:", *issue.HTMLURL)
		time.Sleep(time.Millisecond * 640)
	}
}

func showPullRequests(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name
	ctx := context.Background()

	states := []string{"open", "closed", "merged"}
	titlecaser := cases.Title(language.English)

	for _, state := range states {
		prs, _, err := client.PullRequests.List(ctx, owner, name,
			&github.PullRequestListOptions{
				State:       state,
				ListOptions: github.ListOptions{PerPage: 5},
			})

		if err != nil {
			fmt.Printf("Error fetching %s PRs: %v\n", state, err)
			continue
		}

		fmt.Printf("\n\033[1m%s PRs (%d):\033[0m\n", titlecaser.String(state), len(prs))
		for _, pr := range prs {
			printPRSummary(pr)
		}
		time.Sleep(time.Millisecond * 640)
	}
}

func showContributors(repo *github.Repository, client *github.Client) {
	owner := *repo.Owner.Login
	name := *repo.Name
	ctx := context.Background()
	contributors, _, err := client.Repositories.ListContributors(ctx, owner, name,
		&github.ListContributorsOptions{
			ListOptions: github.ListOptions{PerPage: 20},
		})

	if err != nil {
		fmt.Printf("Error fetching contributors: %v\n", err)
		return
	}

	fmt.Printf("\n\033[1;36mTop Contributors (%d):\033[0m\n", len(contributors))

	// Sort by contributions (descending)
	sort.Slice(contributors, func(i, j int) bool {
		return *contributors[i].Contributions > *contributors[j].Contributions
	})

	for i, contributor := range contributors {
		if i >= 10 { // Show top 10
			break
		}

		user, _, _ := client.Users.Get(ctx, *contributor.Login)

		fmt.Printf("\n%d. \033[1m%s\033[0m\n", i+1, *contributor.Login)
		fmt.Printf("├─ Contributions: %d\n", *contributor.Contributions)

		if user.Name != nil {
			fmt.Printf("├─ Name: %s\n", *user.Name)
		}

		if user.Bio != nil {
			fmt.Printf("├─ Bio: %s\n", truncate(*user.Bio, 60))
		}

		fmt.Printf("└─ Profile: %s\n", *user.HTMLURL)
		time.Sleep(time.Millisecond * 640)
	}

}

func printPRSummary(pr *github.PullRequest) {
	if pr == nil || pr.Number == nil || pr.Title == nil || pr.User == nil {
		fmt.Println("Invalid pull request data")
		return
	}

	fmt.Printf("\n#%d: \033[1;34m%s\033[0m\n", *pr.Number, *pr.Title)
	fmt.Printf("├─ By: %s\n", *pr.User.Login)

	if pr.State != nil {
		fmt.Printf("├─ State: %s", *pr.State)
		if pr.MergedAt != nil {
			fmt.Printf(" (merged)")
		}
		fmt.Println()
	}

	if pr.CreatedAt != nil {
		fmt.Printf("├─ Created: %s\n", pr.CreatedAt.Format("2006-01-02"))
	}

	if pr.ChangedFiles != nil && pr.Commits != nil {
		fmt.Printf("├─ Changed Files: %d | Commits: %d\n", *pr.ChangedFiles, *pr.Commits)
	}

	if pr.HTMLURL != nil {
		fmt.Printf("└─ URL: %s\n", *pr.HTMLURL)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func safeString(s *string) string {
	if s == nil {
		return "N/A"
	}
	return *s
}

func truncate(text string, length int) string {
	if len(text) <= length {
		return text
	}
	return text[:length] + "..."
}
