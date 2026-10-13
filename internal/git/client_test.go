package git

import (
	"os"
	"testing"
)

func TestIsGitRepository(t *testing.T) {
	// Current repository must be a valid git repository
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}

	client := NewClient(currentDir)
	if !client.IsGitRepository() {
		t.Errorf("expected current workDir to be recognized as git repo")
	}

	// Temp empty dir must NOT be a git repository
	tempDir, err := os.MkdirTemp("", "gitgwtuh_nongit_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	nonGitClient := NewClient(tempDir)
	if nonGitClient.IsGitRepository() {
		t.Errorf("temp empty dir should not be recognized as git repo")
	}
}

func TestGetRepoSummaryAndCommits(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}

	client := NewClient(currentDir)
	summary, err := client.GetRepoSummary()
	if err != nil {
		t.Fatalf("GetRepoSummary returned unexpected error: %v", err)
	}

	if summary.Name != "gitgwtuh" && summary.Name != "git" {
		t.Logf("repo name: %s", summary.Name)
	}

	commits, err := client.GetCommits(10, 0)
	if err != nil {
		t.Fatalf("GetCommits returned error: %v", err)
	}

	if len(commits) == 0 {
		t.Errorf("expected at least 1 commit in current repository")
	}
}
