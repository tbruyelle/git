// Package git provides methods to execute raw git commands
// in the current folder.
package git

import (
	"errors"
	"fmt"
	"github.com/tbruyelle/qexec"
	"regexp"
	"strings"
)

// Remote returns the requested remote url
func Remote(name string) (string, error) {
	remote, err := qexec.Run("git", "config", "--get", fmt.Sprintf("remote.%s.url", name))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(remote), err
}

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

// RepoInfo returns structured data about the current repository.
func Repository() (*Repo, error) {
	remote, err := Remote("origin")
	if err != nil {
		return nil, err
	}
	return parseRepo(remote)
}

type Repo struct {
	Owner, Name string
}

var remoteGitUrlSsh = regexp.MustCompile("git@\\S+:(\\w+)/(\\w+)(\\.git)?")
var remoteGitUrlHttp = regexp.MustCompile("https?://\\S+/(\\w+)/(\\w+)(\\.git)?")

func parseRepo(remote string) (*Repo, error) {
	if strings.Index(remote, "http") != -1 {
		res := remoteGitUrlHttp.FindAllStringSubmatch(remote, -1)
		if len(res) == 0 || len(res[0]) < 3 {
			return nil, errors.New("Unable to parse remote " + remote)
		}
		repo := &Repo{
			Owner: res[0][1],
			Name:  res[0][2],
		}
		return repo, nil
	} else if strings.Index(remote, "git@") != -1 {
		res := remoteGitUrlSsh.FindAllStringSubmatch(remote, -1)
		if len(res) == 0 || len(res[0]) < 3 {
			return nil, errors.New("Unable to parse remote " + remote)
		}
		repo := &Repo{
			Owner: res[0][1],
			Name:  res[0][2],
		}
		return repo, nil
	}
	return nil, errors.New("Unhandled remote " + remote)
}
