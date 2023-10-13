package statistics

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func Pushing(repoPath string) error {
	// Open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	// Get the branch reference
	ref, err := repo.Head()
	if err != nil {
		return err
	}

	// Get the remote
	remote, err := repo.Remote("origin")
	if err != nil {
		return err
	}

	// Push using the default authentication (whatever is set up in the local Git config)
	err = remote.Push(&git.PushOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
		RefSpecs:   []config.RefSpec{config.RefSpec(fmt.Sprintf("+%s:%s", ref.Name().String(), ref.Name().String()))},
		Auth:       nil, // No explicit auth provided here, it will use the default from the Git config
	})
	if err != nil {
		// Handle the "everything up-to-date" case
		if err == git.NoErrAlreadyUpToDate {
			return nil
		}
		return err
	}

	return nil
}
