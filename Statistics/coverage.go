package statistics

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// coveragePercentage calculates the code coverage percentage of a Go project.
func CoveragePercentage(repoPath string) (float64, error) {
	cmd := exec.Command("go", "list", "./...")
	cmd.Dir = repoPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("error listing packages: %v", err)
	}

	packages := strings.Split(string(out), "\n")

	totalCoverage := 0.0
	coveredPackages := 0

	// Regex to extract coverage percentage
	re := regexp.MustCompile(`coverage: ([\d.]+)%`)

	for _, pkg := range packages {
		if strings.TrimSpace(pkg) == "" {
			continue
		}

		cmd := exec.Command("go", "test", "-cover", pkg)
		cmd.Dir = repoPath
		out, err := cmd.CombinedOutput()
		if err != nil {
			return 0, fmt.Errorf("error running tests for package %s: %v", pkg, err)
		}

		// Extract coverage from the output
		matches := re.FindStringSubmatch(string(out))
		if len(matches) > 1 {
			coverageStr := matches[1]
			coverage, err := strconv.ParseFloat(coverageStr, 64)
			if err != nil {
				return 0, fmt.Errorf("error parsing coverage for package %s: %v", pkg, err)
			}
			totalCoverage += coverage
			coveredPackages++
		}
	}

	// Calculate average coverage
	averageCoverage := totalCoverage / float64(coveredPackages)
	return averageCoverage, nil
}
