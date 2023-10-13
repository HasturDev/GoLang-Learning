package main

import (
	"fmt"
	"log"
	"os"

	statistics "xithulhu.com/Statistics" // the import path
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <command> [args...]")
		fmt.Println("Commands: pull, push, coverage, cleanliness")
		return
	}

	command := os.Args[1]
	repoPath := "./" // Default repository path. Adjust as needed or get it as an argument.
	switch command {
	case "pull":
		// Call the pull function (assuming it's named PullRepo)
		// You might need owner and repo name arguments.
		if len(os.Args) < 4 {
			fmt.Println("Usage: program pull <owner> <repo_name>")
			return
		}
		owner := os.Args[2]
		repoName := os.Args[3]
		dir := os.Args[4]
		err := statistics.PullRepo(dir, owner, repoName)
		if err != nil {
			log.Fatalf("Error pulling repository: %v", err)
		}

	//case "push":
	//	// Call the push function (assuming it's named PushRepo)
	//	if err := statistics.Pushing("."); err != nil {
	//		log.Fatalf("Failed to push: %s", err)
	//	}

	case "coverage":
		// Call the coverage function
		coverage, err := statistics.CoveragePercentage(repoPath)
		if err != nil {
			log.Fatalf("Error calculating coverage: %v", err)
		}
		fmt.Printf("Code coverage: %.2f%%\n", coverage)

	case "cleanliness":
		// Call the cleanliness function
		cleanliness, err := statistics.CleanlinessPercentage(repoPath)
		if err != nil {
			log.Fatalf("Error estimating code cleanliness: %v", err)
		}
		fmt.Printf("Code cleanliness: %.2f%%\n", cleanliness)

	default:
		fmt.Println("Unknown command:", command)

		// func (r *Repository) Log(o *LogOptions) (object.CommitIter, error)
		fmt.Println("Welcome to main() function")
		myCode := statistics.Code{FileSomething: "File1", CodeSomething: "Some Code"}

		result := statistics.Search(myCode)
		fmt.Println(result)
	}
}
