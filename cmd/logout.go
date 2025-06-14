/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/DipanshuOjha/cobraclip/internal/config"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout your github account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Processing Log out request.........\n")
		_, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Not able to config the user token(That means it is not present) \n %w\n", err)
			return
		}

		token, err := config.SendToken()
		if err != nil {
			fmt.Println("Error while fecthing current token check out:- ", err)
			return
		}
		err = config.RemoveToken(token)

		if err != nil {
			fmt.Println("Error while removing token check out :- ", err)
			return
		}
		fmt.Printf("✅ Token removed successfully.......\n")
		fmt.Printf("🙏 Thanks for using cobraclip\n")
		fmt.Println("❤️ Love Dipanshu Ojha")

	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
