/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/DipanshuOjha/cobraclip/internal/config"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login using github personal access token",

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("DEBUG 1: Command started") // This will confirm if command executes
		// defer fmt.Println("DEBUG 2: Command completed")
		// _, err := config.Loadconfig()
		// if err != nil {
		// 	fmt.Printf("Error loading config: %v\n", err)
		// 	return
		// }

		// fmt.Printf("Loaded the config\n")
		fmt.Println("Enter your Github personal access token: ")
		tokenbyte, err := term.ReadPassword(int(syscall.Stdin))

		//fmt.Println("you are here")

		if err != nil {
			fmt.Printf("Error reading token try again \n%v", err)
			return
		}
		token := strings.TrimSpace(string((tokenbyte)))
		fmt.Println()

		if token == "" {
			fmt.Printf("token should not be empty \n%v", err)
			return
		}

		err = config.SaveToken(token)
		if err != nil {
			fmt.Printf("❌ Failed to save token securely: %v\n", err)
			return
		}

		if config.HasToken() {
			fmt.Println("✅ Token is saved securely in your system keyring.")
		} else {
			fmt.Println("❌ No token found. Please run 'cobraclip login'.")
		}

		// Verify by loading config
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error verifying config: %v\n", err)
			return
		}
		fmt.Printf("Logged in successfully with token: %s\n", cfg.GitToken)
		fmt.Printf("Just restart the powershell to enjoy all the features....have a good day")

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
