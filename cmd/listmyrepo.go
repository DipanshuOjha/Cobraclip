/*
Copyright © 2025 Dipanshu Ojha
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	optionsforuser "github.com/DipanshuOjha/cobraclip/functions/optionsForUser"
	"github.com/DipanshuOjha/cobraclip/internal/config"
	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

var (
	cntforrepo int
)

// listmyrepoCmd represents the listmyrepo command
var listmyrepoCmd = &cobra.Command{
	Use:   "listMyRepo",
	Short: "a command to list personal repos",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Not able to config the user token i guess you should login using cobraclip login \n %w\n", err)
			return
		}

		client, err := config.GetGithubClient(cfg)

		if err != nil {
			fmt.Printf("Not able to create github client check out  %v\n", err)
			return
		}

		ctx := context.Background()
		user, _, err := client.Users.Get(ctx, "")
		if err != nil {
			fmt.Printf("Error fetching user: %v\n", err)
			return
		}
		username := *user.Login
		opt := &github.RepositoryListByUserOptions{
			Type:        "updated",
			ListOptions: github.ListOptions{PerPage: 100},
		}

		var repos []*github.Repository

		for {
			repo, resp, err := client.Repositories.ListByUser(ctx, username, opt)
			if err != nil {
				fmt.Printf("Error listing repositories: %v (status: %s)\n", err, resp.Status)
				return
			}

			repos = append(repos, repo...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}

		fmt.Printf("Total Number of Repository :- %d\n", len(repos))
		fmt.Println()
		for i, repo := range repos {
			if i == cntforrepo {
				break
			}
			fmt.Printf("%d Repository: %s\n", i+1, *repo.FullName)
			desc := "No description"
			if repo.Description != nil {
				desc = *repo.Description
			}
			fmt.Printf("   Description: %s\n", desc)
			fmt.Printf("   Stars: %d\n", *repo.StargazersCount)
			fmt.Printf("   URL: %s\n", *repo.HTMLURL)
			fmt.Println()
			time.Sleep(time.Millisecond * 500)
		}

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("That all your repo looking to checkout a repo then enter its index here or enter q to free the terminal:-")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("What you said i didnt read try again")
				continue
			}

			num, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				break
			}

			if num > len(repos) || num <= 0 {
				fmt.Printf("Enter within the range")
				continue
			}

			optionsforuser.Options(repos[num-1], client)
		}

	},
}

func init() {
	repoCmd.AddCommand(listmyrepoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listmyrepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listmyrepoCmd.Flags().IntVarP(&cntforrepo, "count", "c", 10000000, "Help message for toggle")
}
