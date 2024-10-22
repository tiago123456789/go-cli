package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type SearchGithubRepositoriesCommand struct {
}

func NewSearchGithubRepositoriesCommand() *SearchGithubRepositoriesCommand {
	return &SearchGithubRepositoriesCommand{}
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
}

func listRepositories(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repositories: %s", resp.Status)
	}

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func (cP *SearchGithubRepositoriesCommand) Get() *cobra.Command {
	var profile string

	var cmd = &cobra.Command{
		Use:   "search_github_repositories",
		Short: "Search repositories from public Github profile",
		Run: func(cmd *cobra.Command, args []string) {
			if profile == "" {
				fmt.Println("You must supply a Github profile.")
				return
			}
			fmt.Println("Search repositories...")

			repositories, err := listRepositories(profile)
			if err != nil {
				fmt.Println("Couldn't found the repositories from Github profile.")
				return
			}

			for _, repository := range repositories {
				fmt.Println("Repository name => ", repository.Name, " | Link => ", repository.HTMLURL)
			}
		},
	}

	cmd.Flags().StringVarP(&profile, "profile", "p", "", "The Github profile")

	return cmd
}
