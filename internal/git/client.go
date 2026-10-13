package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Client handles command interactions with the local Git binary.
type Client struct {
	workDir string
}

// NewClient creates a new Git client instance targeted at workDir.
func NewClient(workDir string) *Client {
	return &Client{workDir: workDir}
}

// run executes a git command in the working directory and returns its output.
func (c *Client) run(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = c.workDir

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s failed: %w (%s)", strings.Join(args, " "), err, strings.TrimSpace(errOut.String()))
	}

	return out.String(), nil
}

// IsGitRepository verifies whether the working directory is inside a valid Git work tree.
func (c *Client) IsGitRepository() bool {
	out, err := c.run("rev-parse", "--is-inside-work-tree")
	if err != nil {
		return false
	}
	return strings.TrimSpace(out) == "true"
}

// GetRepoSummary inspects repository metadata for top bar telemetry.
func (c *Client) GetRepoSummary() (RepoSummary, error) {
	summary := RepoSummary{
		Name:          filepath.Base(c.workDir),
		CurrentBranch: "HEAD",
		IsClean:       true,
	}

	// Get current branch name
	if branchOut, err := c.run("rev-parse", "--abbrev-ref", "HEAD"); err == nil {
		summary.CurrentBranch = strings.TrimSpace(branchOut)
	}

	// Check if repository has uncommitted changes
	if statusOut, err := c.run("status", "--porcelain"); err == nil {
		if len(strings.TrimSpace(statusOut)) > 0 {
			summary.IsClean = false
		}
	}

	// Count total commits
	if countOut, err := c.run("rev-list", "--count", "HEAD"); err == nil {
		if count, err := strconv.Atoi(strings.TrimSpace(countOut)); err == nil {
			summary.TotalCommits = count
		}
	}

	// Count unique authors
	if authorOut, err := c.run("shortlog", "-sn", "HEAD"); err == nil {
		lines := strings.Split(strings.TrimSpace(authorOut), "\n")
		validAuthors := 0
		for _, l := range lines {
			if strings.TrimSpace(l) != "" {
				validAuthors++
			}
		}
		summary.TotalAuthors = validAuthors
	}

	return summary, nil
}

// GetCommits fetches commits with limit and offset paging.
func (c *Client) GetCommits(limit, skip int) ([]Commit, error) {
	format := "%x00%H%x00%h%x00%an%x00%ae%x00%ad%x00%ar%x00%s%x00%b"
	args := []string{
		"log",
		fmt.Sprintf("-n%d", limit),
		fmt.Sprintf("--skip=%d", skip),
		fmt.Sprintf("--format=%s", format),
		"--date=format:%Y-%m-%d %H:%M",
	}

	out, err := c.run(args...)
	if err != nil {
		return nil, err
	}

	return ParseLogOutput(out), nil
}

// GetCommitFiles fetches modified files and changed line statistics for a specific commit hash.
func (c *Client) GetCommitFiles(hash string) ([]FileChange, error) {
	out, err := c.run("show", "--numstat", "--format=", hash)
	if err != nil {
		return nil, err
	}

	return ParseNumstatOutput(out), nil
}

// GetCommitDiff fetches the raw unified diff for a commit, optionally filtered by file path.
func (c *Client) GetCommitDiff(hash, filePath string) (string, error) {
	args := []string{"show", "--color=never", "--unified=3", hash}
	if filePath != "" {
		args = append(args, "--", filePath)
	}

	out, err := c.run(args...)
	if err != nil {
		return "", err
	}

	return out, nil
}
