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

func TestHeaderSingleLineStrict(t *testing.T) {
	summary := git.RepoSummary{
		CurrentBranch: "feature/super-long-branch-name-for-testing",
		TotalCommits:  12345,
		TotalAuthors:  42,
		Name:          "super-long-repository-name",
		IsClean:       false,
	}

	widths := []int{75, 80, 85, 90, 100, 120, 160, 200}
	for _, w := range widths {
		header := RenderHeader(summary, 1, true, w)
		h := lipgloss.Height(header)
		actualW := lipgloss.Width(header)
		if h != 1 {
			t.Errorf("width %d: expected header height 1, got %d", w, h)
		}
		if actualW > w {
			t.Errorf("width %d: header width %d exceeded allocated width", w, actualW)
		}
	}
}

func TestBoxDimensionParity(t *testing.T) {
	commits := []git.Commit{
		{Hash: "abc1234567890", ShortHash: "abc1234", Subject: "feat: add animations", Author: "dika", RelativeDate: "10 mins ago"},
	}
	files := []git.FileChange{
		{Status: "M", Path: "internal/ui/app.go", Additions: 15, Deletions: 5},
	}
	diffText := "diff --git a/a.txt b/a.txt\n--- a/a.txt\n+++ b/a.txt\n@@ -1 +1 @@\n-old\n+new\n"

	testCases := []struct {
		outerW int
		outerH int
	}{
		{30, 12},
		{45, 18},
		{60, 24},
		{80, 30},
	}

	for _, tc := range testCases {
		resList := RenderCommitList(commits, 0, 0, tc.outerW, tc.outerH, 0, false, true)
		if w := lipgloss.Width(resList); w != tc.outerW {
			t.Errorf("RenderCommitList: expected width %d, got %d", tc.outerW, w)
		}
		if h := lipgloss.Height(resList); h != tc.outerH {
			t.Errorf("RenderCommitList: expected height %d, got %d", tc.outerH, h)
		}

		resDetail := RenderCommitDetail(&commits[0], files, 0, tc.outerW, tc.outerH, 0, false, true)
		if w := lipgloss.Width(resDetail); w != tc.outerW {
			t.Errorf("RenderCommitDetail: expected width %d, got %d", tc.outerW, w)
		}
		if h := lipgloss.Height(resDetail); h != tc.outerH {
			t.Errorf("RenderCommitDetail: expected height %d, got %d", tc.outerH, h)
		}

		resDiff := RenderDiffView(diffText, 0, tc.outerW, tc.outerH, 0, false, true)
		if w := lipgloss.Width(resDiff); w != tc.outerW {
			t.Errorf("RenderDiffView: expected width %d, got %d", tc.outerW, w)
		}
		if h := lipgloss.Height(resDiff); h != tc.outerH {
			t.Errorf("RenderDiffView: expected height %d, got %d", tc.outerH, h)
		}
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
