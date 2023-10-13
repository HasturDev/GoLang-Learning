// cleanliness.go inside the "statistics" package
package statistics

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// CleanlinessPercentage calculates a simplistic cleanliness percentage for a Go project based on linting issues.
func CleanlinessPercentage(repoPath string) (float64, error) {
	// Run golangci-lint on the repository
	cmd := exec.Command("golangci-lint", "run", "--out-format", "json")
	cmd.Dir = repoPath
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	_, ok := err.(*exec.ExitError)
	if err != nil && !ok { // golangci-lint returns an exit error when issues are found, which is expected
		return 0, fmt.Errorf("error running golangci-lint: %v", err)
	}

	// Parse the output to count issues
	issuesCount := strings.Count(out.String(), `"text":`)

	// Here, we make an assumption:
	// Suppose every file in a large project might have on average 1 linting issue. If there are 100 files,
	// then 100 potential linting issues. If our lint tool finds 50 issues, then the cleanliness is 50%.
	// This is, of course, a big assumption. For better results, you might want to have better metrics
	// or perhaps compare against other projects of similar size.

	// Get total number of go files in the repository
	totalFiles, err := totalGoFiles(repoPath)
	if err != nil {
		return 0, fmt.Errorf("error counting Go files: %v", err)
	}

	cleanliness := (1 - float64(issuesCount)/float64(totalFiles)) * 100
	return cleanliness, nil
}

// totalGoFiles counts the total number of Go files in the repository
func totalGoFiles(repoPath string) (int, error) {
	cmd := exec.Command("find", ".", "-name", "*.go", "|", "wc", "-l")
	cmd.Dir = repoPath
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 0, err
	}

	return count, nil
}
