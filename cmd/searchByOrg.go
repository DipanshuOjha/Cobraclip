/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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
	org string
	cnt int
)

// searchByOrgCmd represents the searchByOrg command
var searchByOrgCmd = &cobra.Command{
	Use:   "SearchByOrg",
	Short: "Search Github Repository by Organisation",
	Long: `
	   parameter that are to search repositories within an organisation
	                --org or -o name of organisation (Required)
					--cnt or -c Number of repos you want to print
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Not able to config the user token i guess you should login using cobraclip login \n %w\n", err)
			return
		}

		client, err := config.GetGithubClient(cfg)

		if err != nil {
			fmt.Println("Not able to create github client check out \n %w\n", err)
			return
		}

		ctx := context.Background()

		data := &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{PerPage: 100},
		}

		var AllRepos []*github.Repository

		for {
			repos, resp, err := client.Repositories.ListByOrg(ctx, org, data)
			if err != nil {
				fmt.Printf("Error listing repositories: %v\n", err)
				if resp != nil {
					fmt.Printf("GitHub API response: %s\n", resp.Status)
				}
				return
			}
			AllRepos = append(AllRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			data.Page = resp.NextPage
		}

		if len(AllRepos) == 0 {
			fmt.Printf("No repositories found for organization: %s\n", org)
			return
		}

		fmt.Printf("Found %d repositories for %s\n", len(AllRepos), org)
		for i, repo := range AllRepos {
			fmt.Printf("\n%d. %s\n", i+1, *repo.FullName)
			desc := "No description"
			if repo.Description != nil {
				desc = *repo.Description
			}
			fmt.Printf("   Description: %s\n", desc)
			fmt.Printf("   Stars: %d\n", *repo.StargazersCount)
			fmt.Printf("   URL: %s\n", *repo.HTMLURL)
			time.Sleep(time.Millisecond * 700)
			if i+1 == cnt {
				break
			}
		}

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

			if num > len(AllRepos) || num <= 0 {
				fmt.Printf("Enter within the range")
				continue
			}

			detaillog.ShowRepoDetail(AllRepos[num-1], client)
		}

	},
}

func init() {
	repoCmd.AddCommand(searchByOrgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchByOrgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchByOrgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchByOrgCmd.Flags().StringVarP(&org, "org", "o", "", "Name of Organisation(Required)")
	searchByOrgCmd.Flags().IntVarP(&cnt, "cnt", "c", 100000000, "Count of repos")
	searchByOrgCmd.MarkFlagRequired("org")
}
