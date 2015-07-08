// Package git provides methods to execute raw git commands
// in the current folder.
package git

import (
	"fmt"
	"github.com/tbruyelle/qexec"
	"strings"
)

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
		tokens := strings.SplitN(line, " ", 1)
		fmt.Printf("%+v %d\n", tokens, len(tokens))
		if len(tokens) != 2 {
			continue
		}
		commits = append(commits, Commit{tokens[0], tokens[1]})
	}
	return commits, nil
}
