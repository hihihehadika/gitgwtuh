package components

import (
	"fmt"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderHeader renders the top telemetry header with repository statistics and pill badges.
func RenderHeader(summary git.RepoSummary, width int) string {
	leftPills := lipgloss.JoinHorizontal(
		lipgloss.Center,
		styles.BrandBadge.Render("gitgwtuh"),
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

	rightPills := lipgloss.JoinHorizontal(
		lipgloss.Center,
		statusPill,
		commitPill,
		authorPill,
		repoPill,
	)

	// Calculate space between left badges and right badges
	totalPillsWidth := lipgloss.Width(leftPills) + lipgloss.Width(rightPills)
	spacerWidth := width - totalPillsWidth
	if spacerWidth < 1 {
		spacerWidth = 1
	}

	spacer := lipgloss.NewStyle().Width(spacerWidth).Render("")

	return lipgloss.JoinHorizontal(lipgloss.Center, leftPills, spacer, rightPills)
}
