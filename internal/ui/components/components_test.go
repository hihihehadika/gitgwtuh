package components

import (
	"strings"
	"testing"

	"gitgwtuh/internal/git"

	"github.com/charmbracelet/lipgloss"
)

func TestRenderHeader(t *testing.T) {
	summary := git.RepoSummary{
		CurrentBranch: "main",
		TotalCommits:  10,
		Name:          "test-repo",
	}

	headerAnim := RenderHeader(summary, 2, true, 100)
	if !strings.Contains(headerAnim, "gitgwtuh") {
		t.Errorf("expected header to contain brand name")
	}
	if !strings.Contains(headerAnim, "main") {
		t.Errorf("expected header to contain branch name")
	}

	headerIdle := RenderHeader(summary, 0, false, 100)
	if !strings.Contains(headerIdle, "IDLE") {
		t.Errorf("expected idle header to contain IDLE pill")
	}

	if lipgloss.Width(headerAnim) > 100 {
		t.Errorf("header width %d exceeded allocated width 100", lipgloss.Width(headerAnim))
	}
}

func TestRenderCommitList(t *testing.T) {
	commits := []git.Commit{
		{Hash: "abc1234567890", ShortHash: "abc1234", Subject: "feat: add cyber animations", Author: "dika", RelativeDate: "10 mins ago"},
		{Hash: "def1234567890", ShortHash: "def1234", Subject: "fix: prevent visual clipping", Author: "dika", RelativeDate: "1 hour ago"},
	}

	rendered := RenderCommitList(commits, 0, 0, 40, 20, 0, true, true)
	if len(rendered) == 0 {
		t.Errorf("expected non-empty commit list render")
	}
	if !strings.Contains(rendered, "COMMIT TIMELINE") {
		t.Errorf("expected panel title in render")
	}
}

func TestRenderCommitDetail(t *testing.T) {
	commit := &git.Commit{
		Hash:         "abc1234567890",
		ShortHash:    "abc1234",
		Author:       "dika",
		Email:        "dika@example.com",
		Date:         "2026-10-14",
		RelativeDate: "10 mins ago",
		Subject:      "feat: add cyber animations",
	}
	files := []git.FileChange{
		{Status: "M", Path: "internal/ui/app.go", Additions: 15, Deletions: 5},
	}

	rendered := RenderCommitDetail(commit, files, 0, 60, 8, 0, true, false)
	t.Logf("height of 8: %d", lipgloss.Height(rendered))
	for i, l := range strings.Split(rendered, "\n") {
		t.Logf("line %d: %q", i, l)
	}
	if len(rendered) == 0 {
		t.Errorf("expected non-empty commit detail render")
	}
	if !strings.Contains(rendered, "COMMIT DETAIL & FILES") {
		t.Errorf("expected panel title in render")
	}
}

func TestRenderDiffView(t *testing.T) {
	diffText := "diff --git a/a.txt b/a.txt\n--- a/a.txt\n+++ b/a.txt\n@@ -1 +1 @@\n-old\n+new\n"

	rendered := RenderDiffView(diffText, 0, 60, 10, 0, true, true)
	if len(rendered) == 0 {
		t.Errorf("expected non-empty diff view render")
	}
	if !strings.Contains(rendered, "SYNTAX DIFF VIEWER") {
		t.Errorf("expected panel title in render")
	}
}

func TestRenderFooter(t *testing.T) {
	footer := RenderFooter("LOGS", "", 100)
	if !strings.Contains(footer, "Animasi") {
		t.Errorf("expected footer to contain Animasi shortcut")
	}
	if lipgloss.Width(footer) > 100 {
		t.Errorf("footer width %d exceeded allocated width 100", lipgloss.Width(footer))
	}
}
