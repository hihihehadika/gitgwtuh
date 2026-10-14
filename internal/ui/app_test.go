package ui

import (
	"os"
	"strings"
	"testing"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/components"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func TestAnimationToggleAndTick(t *testing.T) {
	client := git.NewClient(".")
	model := NewAppModel(client)

	if !model.animEnabled {
		t.Errorf("expected animation to be enabled by default")
	}

	// Test tick increments frame
	newModel, _ := model.Update(tickMsg{})
	m := newModel.(*AppModel)
	if m.animFrame != 1 {
		t.Errorf("expected animFrame to increment to 1, got %d", m.animFrame)
	}

	// Press 'a' to toggle animation off
	newModel, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})
	m = newModel.(*AppModel)
	if m.animEnabled {
		t.Errorf("expected animEnabled to be false after toggle")
	}

	// Tick when disabled should not increment frame
	newModel, _ = m.Update(tickMsg{})
	m = newModel.(*AppModel)
	if m.animFrame != 1 {
		t.Errorf("expected animFrame to remain 1 when disabled, got %d", m.animFrame)
	}
}

func TestViewLineBudgetStrictClamping(t *testing.T) {
	client := git.NewClient(".")
	model := NewAppModel(client)

	model.allCommits = []git.Commit{
		{Hash: "111", ShortHash: "111", Author: "alice", Subject: "feat: add cyber animations", RelativeDate: "10 mins ago"},
		{Hash: "222", ShortHash: "222", Author: "bob", Subject: "fix: prevent visual clipping", RelativeDate: "1 hour ago"},
	}
	model.commits = model.allCommits

	dimensions := [][2]int{
		{80, 24},
		{100, 30},
		{120, 40},
		{160, 50},
	}

	for _, dim := range dimensions {
		model.width = dim[0]
		model.height = dim[1]

		bodyHeight := model.height - 3
		leftWidth := (model.width * 40) / 100
		rightWidth := model.width - leftWidth
		detailHeight := (bodyHeight * 36) / 100
		if detailHeight < 8 {
			detailHeight = 8
		}
		diffHeight := bodyHeight - detailHeight

		hH := lipgloss.Height(components.RenderHeader(model.summary, 0, false, model.width))
		hL := lipgloss.Height(components.RenderCommitList(model.commits, 0, 0, leftWidth, bodyHeight, 0, false, true))
		hD := lipgloss.Height(components.RenderCommitDetail(&model.commits[0], model.files, 0, rightWidth, detailHeight, 0, false, false))
		hDiff := lipgloss.Height(components.RenderDiffView(model.diff, 0, rightWidth, diffHeight, 0, false, false))
		hF := lipgloss.Height(components.RenderFooter("LOGS", "", model.width))

		view := model.View()
		lines := strings.Split(view, "\n")
		t.Logf("size %dx%d: header=%d, left=%d(want %d), detail=%d(want %d), diff=%d(want %d), footer=%d -> total=%d",
			dim[0], dim[1], hH, hL, bodyHeight, hD, detailHeight, hDiff, diffHeight, hF, len(lines))
		if len(lines) > model.height {
			t.Errorf("terminal size %dx%d: view lines (%d) exceeded height (%d)", model.width, model.height, len(lines), model.height)
		}
	}
}
