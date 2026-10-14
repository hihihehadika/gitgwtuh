package components

import (
	"fmt"
	"strings"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderCommitList renders the interactive scrollable commit history list.
func RenderCommitList(commits []git.Commit, selectedIndex, scrollOffset, outerWidth, outerHeight, frame int, animEnabled, isActive bool) string {
	boxStyle := styles.PanelBoxInactive
	if isActive {
		boxStyle = styles.GetActivePanelBox(frame, animEnabled)
	}

	innerW := outerWidth - 4
	if innerW < 10 {
		innerW = 10
	}
	innerH := outerHeight - 2
	if innerH < 1 {
		innerH = 1
	}

	title := styles.PanelTitle.Render("COMMIT TIMELINE")
	separator := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render(strings.Repeat("─", innerW))

	visibleCount := innerH - 2
	if visibleCount < 1 {
		visibleCount = 1
	}

	if len(commits) == 0 {
		emptyMsg := lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Belum ada riwayat commit pada repositori ini.")
		content := lipgloss.JoinVertical(lipgloss.Left, title, separator, emptyMsg)
		return boxStyle.Width(innerW).Height(innerH).Render(content)
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
		fixedWidth := lipgloss.Width(marker) + 1 + lipgloss.Width(hashText) + 1 + lipgloss.Width(timeText) + 1
		availSubjectWidth := innerW - fixedWidth
		subject := c.Subject
		if availSubjectWidth > 4 && len(subject) > availSubjectWidth {
			subject = subject[:availSubjectWidth-3] + "..."
		} else if availSubjectWidth <= 4 {
			subject = ""
		}

		line := fmt.Sprintf("%s %s %s %s", marker, hashText, subject, timeText)

		if isSelected {
			styledLine := styles.CommitItemSelected.MaxWidth(innerW).Render(line)
			rows = append(rows, styledLine)
		} else {
			styledLine := styles.CommitItemNormal.MaxWidth(innerW).Render(line)
			rows = append(rows, styledLine)
		}
	}

	content := lipgloss.JoinVertical(lipgloss.Left, title, separator, strings.Join(rows, "\n"))
	lines := strings.Split(content, "\n")
	if len(lines) > innerH {
		lines = lines[:innerH]
	}
	content = strings.Join(lines, "\n")

	return boxStyle.Width(innerW).Height(innerH).Render(content)
}
