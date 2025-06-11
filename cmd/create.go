/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/DipanshuOjha/cobraclip/internal/config"
	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

var (
	name        string
	description string
	private     bool
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a github repository",
	Long: `
	Their is many way to create a repo 

	1. you can pass a flag directly using --name or -n for repo name 
	                                  and --desc or -d for decription 
								      and --private or -p for making private repo
	2. just call cobraclip create and they will do the rest for you

	dont forget to add repo before you run any command related to repository 
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

		repo := &github.Repository{
			Name:        &name,
			Description: &description,
			Private:     &private,
		}

		createRepo, _, err := client.Repositories.Create(ctx, "", repo)

		if err != nil {
			fmt.Println("Not able to create repository check out:- \n %w\n", err)
			return
		}

		fmt.Printf("Github Repository Created succesfully :- %s\n", *createRepo.HTMLURL)

	},
}

func init() {
	repoCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Repository name (Required)")
	createCmd.Flags().StringVarP(&description, "desc", "d", "", "Repository Description")
	createCmd.Flags().BoolVarP(&private, "private", "p", false, "Make repository private(optional)")
	createCmd.MarkFlagRequired("repo")
}
