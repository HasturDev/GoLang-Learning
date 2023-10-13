package statistics

import (
	"os"

	git "github.com/go-git/go-git/v5"
)

// CloneRepo clones a repository based on the provided HTTPS URL.
func CloneRepo(repoURL string, directory string) error {
	print(directory)
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
		Auth:     nil, // Assuming that if authentication is needed, it's set up in the Git configuration
	})
	return err
}
