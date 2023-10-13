package statistics

import (
	"errors"
	"fmt"
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// 	myError := errors.New("WRONG MESSAGE")
func Pushing(repoPath string) error {
	if len(os.Args) != 2 {
		fmt.Println("Usage: push <repository_path>")
		return errors.New("An error occured in line 18")
	}

	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Cannot open repository: %v", err)
	}

	// Fetch the latest updates from the remote repository
	err = r.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{config.RefSpec("refs/heads/*:refs/heads/*")},
		Auth:     getAuth(),
	})

	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatalf("Failed to fetch: %v", err)
	}

	// Check for any changes that haven't been pushed
	headRef, err := r.Head()
	if err != nil {
		log.Fatalf("Cannot get HEAD: %v", err)
	}

	remoteRef, err := r.Reference(plumbing.NewRemoteReferenceName("origin", headRef.Name().Short()), true)
	if err != nil && err != plumbing.ErrReferenceNotFound {
		log.Fatalf("Cannot get remote ref: %v", err)
	}

	if err == plumbing.ErrReferenceNotFound || headRef.Hash() != remoteRef.Hash() {
		w, err := r.Worktree()
		if err != nil {
			log.Fatalf("Getting worktree failed: %v", err)
		}

		// Pull the latest changes
		err = w.Pull(&git.PullOptions{RemoteName: "origin", Auth: getAuth()})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			log.Fatalf("Failed to pull: %v", err)
		}

		// Push the changes
		err = r.Push(&git.PushOptions{Auth: getAuth()})
		if err != nil {
			log.Fatalf("Failed to push: %v", err)
		}
		fmt.Println("Pushed to repository successfully!")
	} else {
		fmt.Println("No changes to push.")
	}
	return nil
}

func getAuth() *http.BasicAuth {
	// Use this function to retrieve authentication credentials.
	// Replace placeholders with actual username and password/token.
	// It's advisable to use environment variables or a config file rather than hardcoding credentials.
	return &http.BasicAuth{
		Username: "your_username_or_token", // can be just the token for GitHub
		Password: "your_password_or_token",
	}
}
