package optionsforuser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DipanshuOjha/cobraclip/functions/clone"
	deleterepo "github.com/DipanshuOjha/cobraclip/functions/deleteRepo"
	"github.com/DipanshuOjha/cobraclip/functions/detaillog"
	"github.com/DipanshuOjha/cobraclip/functions/update"
	"github.com/google/go-github/v62/github"
)

var optionsMenu = []struct {
	key         string
	description string
	handler     func(*github.Repository, *github.Client)
}{
	{"1", "Clone", makeclone},
	{"2", "Delete", deleteit},
	{"3", "Update Title", updatetitle},
	{"4", "Update Description", updatedescription},
	{"q", "Quit", nil},
}

func updatetitle(repo *github.Repository, client *github.Client) {
	reader := bufio.NewReader(os.Stdin)

	for {
		println("Enter New Title:- ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error while reading try again")
			continue
		}

		title := strings.TrimSpace(input)

		repo, err := update.UpdateTitle(repo, client, title)

		if err != nil {
			fmt.Println("error while updating title of the check out :- ", err)
			break
		}

		fmt.Printf("Repository title is being updated to %q. Check it out at: %s\n", *repo.Name, *repo.HTMLURL)
		break
	}

}

func updatedescription(repo *github.Repository, client *github.Client) {
	reader := bufio.NewReader(os.Stdin)

	for {
		println("Enter New Description:- ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error while reading try again")
			continue
		}

		description := strings.TrimSpace(input)

		repo, err := update.UpdateDescriptionandTitle(repo, client, "", description)

		if err != nil {
			fmt.Println("error while updating Description of the check out :- ", err)
			break
		}

		fmt.Printf("Repository description is being updated to %q. Check it out at: %s\n", *repo.Description, *repo.HTMLURL)
		break
	}
}

func makeclone(repo *github.Repository, client *github.Client) {
	reader := bufio.NewReader(os.Stdin)
	// consider adding absolute

	fmt.Println("Enter the path you want to enter :- ")
	fmt.Println("format should be like C:\\Users\\Desktop\\projects")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("What you said i didnt read try again try again")
		return
	}

	repoPath := strings.TrimSpace(input)

	detaillog.ClearScreen()
	fmt.Println("Starting to clone the repo in the specified path")

	err = clone.CloneTheRepo(repo, repoPath)

	if err != nil {
		fmt.Printf("error cloning the repo\n check out the error:- %v", err)
		return
	}

	fmt.Println("Cloned successfully in your system check out directory you gave")
}

func deleteit(repo *github.Repository, client *github.Client) {
	err := deleterepo.DeleteRepo(repo, client)
	if err != nil {
		fmt.Printf("error deleting the target repo\n check out:- %v", err)
		return
	}

	fmt.Println("Deleted successfully from your system")

}

func Options(repo *github.Repository, client *github.Client) {
	reader := bufio.NewReader(os.Stdin)

OuterLoop:
	for {
		//detaillog.ClearScreen()
		detaillog.ShowBasicInfo(repo)

		fmt.Printf("What operation you want to perform\n")

		for _, item := range optionsMenu {
			fmt.Printf("[%s] %s\n", item.key, item.description)
		}

		fmt.Printf("\n Choose any : \n")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		if choice == "q" {
			break
		}

		switch choice {
		case "1":
			makeclone(repo, client)
		case "2":
			deleteit(repo, client)
			break OuterLoop
		case "3":
			updatetitle(repo, client)
		case "4":
			updatedescription(repo, client)
		default:
			fmt.Println("Wrong input please retype correctly")
		}

	}
}
