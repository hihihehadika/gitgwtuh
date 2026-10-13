package components

import (
	"fmt"
	"strings"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderCommitList renders the interactive scrollable commit history list.
func RenderCommitList(commits []git.Commit, selectedIndex, scrollOffset, width, height int, isActive bool) string {
	boxStyle := styles.PanelBoxInactive
	if isActive {
		boxStyle = styles.PanelBoxActive
	}

	contentWidth := width - 4
	if contentWidth < 10 {
		contentWidth = 10
	}

	title := styles.PanelTitle.Render("COMMIT TIMELINE")
	separator := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render(strings.Repeat("─", contentWidth))

	visibleCount := height - 4
	if visibleCount < 1 {
		visibleCount = 1
	}

	if len(commits) == 0 {
		emptyMsg := lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Belum ada riwayat commit pada repositori ini.")
		content := lipgloss.JoinVertical(lipgloss.Left, title, separator, "", emptyMsg)
		return boxStyle.Width(width).Height(height).Render(content)
	}

	endIndex := scrollOffset + visibleCount
	if endIndex > len(commits) {
		endIndex = len(commits)
	}

	var rows []string
	for i := scrollOffset; i < endIndex; i++ {
		c := commits[i]
		isSelected := i == selectedIndex

		marker := "○"
		if isSelected {
			marker = "●"
		}

		hashText := fmt.Sprintf("[%s]", c.ShortHash)
		timeText := fmt.Sprintf("(%s)", c.RelativeDate)

		// Calculate available width for subject line
		fixedWidth := lipgloss.Width(marker) + 1 + lipgloss.Width(hashText) + 1 + lipgloss.Width(timeText) + 2
		availSubjectWidth := contentWidth - fixedWidth
		subject := c.Subject
		if availSubjectWidth > 5 && len(subject) > availSubjectWidth {
			subject = subject[:availSubjectWidth-3] + "..."
		} else if availSubjectWidth <= 5 {
			subject = ""
		}

		line := fmt.Sprintf("%s %s %s %s", marker, hashText, subject, timeText)

		if isSelected {
			styledLine := styles.CommitItemSelected.Width(contentWidth).Render(line)
			rows = append(rows, styledLine)
		} else {
			styledLine := styles.CommitItemNormal.Width(contentWidth).Render(line)
			rows = append(rows, styledLine)
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Left, title, separator, strings.Join(rows, "\n"))
	return boxStyle.Width(width).Height(height).Render(content)
}
