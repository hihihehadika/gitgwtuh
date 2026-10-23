package components

import (
	"fmt"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderHeader renders the top telemetry header with repository statistics, pill badges, and live wave.
func RenderHeader(summary git.RepoSummary, frame int, animEnabled bool, width int) string {
	var waveBadge string
	if animEnabled {
		waveBadge = styles.WavePill.Render(styles.GetTelemetryWave(frame))
	} else {
		waveBadge = styles.StatPill.Render("IDLE")
	}

	leftPills := lipgloss.JoinHorizontal(
		lipgloss.Left,
		styles.BrandBadge.Render("gitgwtuh"),
		waveBadge,
		styles.BranchPill.Render("* "+summary.CurrentBranch),
	)

	var statusPill string
	if summary.IsClean {
		statusPill = styles.StatusPillClean.Render("CLEAN")
	} else {
		statusPill = styles.StatusPillDirty.Render("DIRTY")
	}

	commitPill := styles.StatPill.Render(fmt.Sprintf("%d COMMITS", summary.TotalCommits))
	authorPill := styles.StatPill.Render(fmt.Sprintf("%d AUTHORS", summary.TotalAuthors))
	repoPill := styles.StatPill.Render("REPO: " + summary.Name)

	leftW := lipgloss.Width(leftPills)
	availForRight := width - leftW - 1

	var rightPillCandidates = [][]string{
		{statusPill, commitPill, authorPill, repoPill},
		{statusPill, commitPill, repoPill},
		{statusPill, commitPill},
		{statusPill},
	}

	var chosenRightPills []string
	for _, cand := range rightPillCandidates {
		joined := lipgloss.JoinHorizontal(lipgloss.Left, cand...)
		if lipgloss.Width(joined) <= availForRight {
			chosenRightPills = cand
			break
		}
	}

	var rightPills string
	if len(chosenRightPills) > 0 {
		rightPills = lipgloss.JoinHorizontal(lipgloss.Left, chosenRightPills...)
	}

	totalPillsWidth := leftW + lipgloss.Width(rightPills)
	spacerWidth := width - totalPillsWidth
	if spacerWidth < 0 {
		spacerWidth = 0
	}

	spacer := lipgloss.NewStyle().Width(spacerWidth).Render("")
	headerLine := lipgloss.JoinHorizontal(lipgloss.Left, leftPills, spacer, rightPills)

	return lipgloss.NewStyle().MaxWidth(width).Render(headerLine)
}
