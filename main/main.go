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
	case "clone":
		if len(os.Args) < 3 {
			fmt.Println("Usage: clone <repository_https_url> <destination_directory>")
			return
		}

		repoURL := os.Args[2]
		destDir := os.Args[3]

		if err := statistics.CloneRepo(repoURL, destDir); err != nil {
			log.Fatalf("Failed to clone: %s", err)
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
