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

	"github.com/DipanshuOjha/cobraclip/functions/detaillog"
	"github.com/DipanshuOjha/cobraclip/internal/config"
	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

var (
	repoToSearch     string
	cntforrepoSearch int
	nameofUser       string
)

// searchrepoCmd represents the searchrepo command
var searchrepoCmd = &cobra.Command{
	Use:   "searchrepo",
	Short: "A command to search github repository globally",
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

		opts := &github.SearchOptions{
			TextMatch:   true,
			ListOptions: github.ListOptions{PerPage: 100},
		}
		query := repoToSearch
		if nameofUser != "" {
			query += fmt.Sprintf(" user:%s", nameofUser)
		}
		repos, resp, err := client.Search.Repositories(ctx, query, opts)

		if err != nil {
			fmt.Printf("Error listing repositories: %v (status: %s)\n", err, resp.Status)
			return
		}

		fmt.Printf("Total Number of Repository we found :- %d\n", *repos.Total)
		fmt.Println()
		for i, repo := range repos.Repositories {
			if i == cntforrepoSearch {
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

		fmt.Println("If you didn't found your target repo try removing the count flag")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("That all Repos any perticular repository you want to look? just tell me the index number you see:- ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("What you said i didnt read try again")
				continue

			}

			num, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				break
			}

			detaillog.ShowRepoDetail(repos.Repositories[num-1], client)
		}

	},
}

func init() {
	repoCmd.AddCommand(searchrepoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchrepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchrepoCmd.Flags().StringVarP(&repoToSearch, "repo", "r", "", "Repo name(required)")
	searchrepoCmd.Flags().IntVarP(&cntforrepoSearch, "count", "c", 10000000, "count of result you want to display")
	searchrepoCmd.Flags().StringVarP(&nameofUser, "name", "n", "", "Name of user")
	searchByOrgCmd.MarkFlagRequired("repo")
}
