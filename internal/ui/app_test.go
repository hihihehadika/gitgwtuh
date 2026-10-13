package ui

import (
	"os"
	"testing"

	"gitgwtuh/internal/git"

	tea "github.com/charmbracelet/bubbletea"
)

func TestAppModelLifecycle(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}

	client := git.NewClient(currentDir)
	model := NewAppModel(client)

	if model.activePane != PaneCommits {
		t.Errorf("expected initial active pane to be PaneCommits")
	}

	// Test Tab key cycles panes
	newModel, _ := model.Update(tea.KeyMsg{Type: tea.KeyTab})
	appModel := newModel.(*AppModel)
	if appModel.activePane != PaneFiles {
		t.Errorf("expected active pane to switch to PaneFiles, got %v", appModel.activePane)
	}

	newModel, _ = appModel.Update(tea.KeyMsg{Type: tea.KeyTab})
	appModel = newModel.(*AppModel)
	if appModel.activePane != PaneDiff {
		t.Errorf("expected active pane to switch to PaneDiff, got %v", appModel.activePane)
	}

	newModel, _ = appModel.Update(tea.KeyMsg{Type: tea.KeyTab})
	appModel = newModel.(*AppModel)
	if appModel.activePane != PaneCommits {
		t.Errorf("expected active pane to cycle back to PaneCommits, got %v", appModel.activePane)
	}
}

func TestAppModelViewMinSizeGuard(t *testing.T) {
	client := git.NewClient(".")
	model := NewAppModel(client)
	model.width = 50
	model.height = 10

	view := model.View()
	if len(view) == 0 {
		t.Errorf("expected view output, got empty string")
	}
}

func TestCommitSearchFilter(t *testing.T) {
	client := git.NewClient(".")
	model := NewAppModel(client)

	model.allCommits = []git.Commit{
		{Hash: "111", ShortHash: "111", Author: "alice", Subject: "feat(auth): add login"},
		{Hash: "222", ShortHash: "222", Author: "bob", Subject: "fix(ui): adjust padding"},
		{Hash: "333", ShortHash: "333", Author: "alice", Subject: "docs: update readme"},
	}
	model.commits = model.allCommits

	// Search for 'auth'
	model.filterCommits("auth")
	if len(model.commits) != 1 || model.commits[0].ShortHash != "111" {
		t.Fatalf("expected 1 commit matching 'auth', got %d", len(model.commits))
	}

	// Search for author 'alice'
	model.filterCommits("alice")
	if len(model.commits) != 2 {
		t.Fatalf("expected 2 commits matching author 'alice', got %d", len(model.commits))
	}

	// Clear query resets to allCommits
	model.filterCommits("")
	if len(model.commits) != 3 {
		t.Fatalf("expected 3 commits after reset, got %d", len(model.commits))
	}
}
