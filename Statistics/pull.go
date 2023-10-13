package statistics

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
)

func RepoExists(dir, owner, repoName string) bool {
	repoPath := filepath.Join(dir, repoName)
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func PullRepo(dir, owner, repoName string) error {
	url := fmt.Sprintf("https://github.com/%s/%s.git", owner, repoName)
	_, err := git.PlainClone(filepath.Join(dir, repoName), false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// If you need to authenticate:
		// Auth: &http.BasicAuth{
		// 	Username: "your_username", // can be anything, GitHub ignores this
		// 	Password: "your_personal_access_token",
		// },
	})
	return err
}

func Pulling() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: pull <directory> <owner> <repo_name>")
		return
	}

	dir := os.Args[1]
	owner := os.Args[2]
	repoName := os.Args[3]

	if RepoExists(dir, owner, repoName) {
		fmt.Printf("Repository %s/%s already exists in %s\n", owner, repoName, dir)
		return
	}

	fmt.Printf("Pulling repository %s/%s into %s...\n", owner, repoName, dir)
	if err := PullRepo(dir, owner, repoName); err != nil {
		log.Fatalf("Error pulling repository: %v", err)
	}
	fmt.Println("Repository pulled successfully!")
}
