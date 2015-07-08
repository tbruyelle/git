// Package git provides methods to execute raw git commands
// in the current folder.
package git

import (
	"github.com/tbruyelle/qexec"
	"strings"
)

// Branch returns the current branch.
func Branch() (string, error) {
	branch, err := qexec.Run("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(branch), nil
}

// RevParse executes the git rev-parse command.
func RevParse(arg string) (string, error) {
	ref, err := qexec.Run("git", "rev-parse", "-q", arg)
	if err != nil {
		return "", nil
	}
	return strings.TrimSpace(ref), nil
}

// Fetch executes the git fetch command.
func Fetch() error {
	_, err := qexec.Run("git", "fetch")
	return err
}

// HasLocalDiff returns true if repo has local modifications.
func HasLocalDiff() (bool, error) {
	_, err := qexec.Run("git", "diff", "--quiet", "HEAD")
	status, err := qexec.ExitStatus(err)
	if err != nil {
		return false, err
	}
	return status != 0, nil
}

// RefExists checks if the ref exists in the repository.
func RefExists(ref string) (bool, error) {
	_, err := qexec.Run("git", "rev-parse", "--quiet", "--verify", ref)
	if err != nil {
		status, err := qexec.ExitStatus(err)
		if err != nil || status != 1 {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

type Commit struct {
	Ref     string
	Message string
}

// Log returns the list of commits from the requested range.
func Log(start, end string) ([]Commit, error) {
	output, err := qexec.Run("git", "log", start+".."+end, "--oneline")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	commits := make([]Commit, 0)
	for _, line := range lines {
		tokens := strings.SplitN(line, " ", 2)
		if len(tokens) != 2 {
			continue
		}
		commits = append(commits, Commit{tokens[0], tokens[1]})
	}
	return commits, nil
}
