package main

import (
	"os"
	"testing"

	"gitgwtuh/internal/git"
)

func TestMainRepoGuard(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}

	client := git.NewClient(currentDir)
	if !client.IsGitRepository() {
		t.Errorf("current directory must be a valid git repository")
	}

	tempDir, err := os.MkdirTemp("", "gitgwtuh_temp_*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempClient := git.NewClient(tempDir)
	if tempClient.IsGitRepository() {
		t.Errorf("empty temp directory should not be recognized as git repo")
	}
}
