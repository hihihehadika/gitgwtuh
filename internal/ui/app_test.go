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
