package components

import (
	"fmt"
	"strings"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderDiffView renders the bottom-right syntax-highlighted code diff viewer.
func RenderDiffView(diffText string, scrollOffset, width, height int, isActive bool) string {
	boxStyle := styles.PanelBoxInactive
	if isActive {
		boxStyle = styles.PanelBoxActive
	}

	contentWidth := width - 4
	if contentWidth < 10 {
		contentWidth = 10
	}

	title := styles.PanelTitle.Render("SYNTAX DIFF VIEWER")
	separator := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render(strings.Repeat("─", contentWidth))

	if strings.TrimSpace(diffText) == "" {
		emptyMsg := lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Tidak ada perubahan kode untuk ditampilkan.")
		content := lipgloss.JoinVertical(lipgloss.Left, title, separator, "", emptyMsg)
		return boxStyle.Width(width).Height(height).Render(content)
	}

	parsedLines := git.ParseDiffOutput(diffText)
	visibleCount := height - 4
	if visibleCount < 1 {
		visibleCount = 1
	}

	if scrollOffset > len(parsedLines)-visibleCount {
		scrollOffset = len(parsedLines) - visibleCount
	}
	if scrollOffset < 0 {
		scrollOffset = 0
	}

	endIndex := scrollOffset + visibleCount
	if endIndex > len(parsedLines) {
		endIndex = len(parsedLines)
	}

	var renderedRows []string
	for i := scrollOffset; i < endIndex; i++ {
		line := parsedLines[i]

		var numGutter string
		if line.NewLineNo > 0 {
			numGutter = styles.DiffLineNum.Render(fmt.Sprintf("%4d", line.NewLineNo))
		} else if line.OldLineNo > 0 {
			numGutter = styles.DiffLineNum.Render(fmt.Sprintf("%4d", line.OldLineNo))
		} else {
			numGutter = styles.DiffLineNum.Render("    ")
		}

		lineContentWidth := contentWidth - lipgloss.Width(numGutter) - 1
		if lineContentWidth < 5 {
			lineContentWidth = 5
		}

		var styledContent string
		switch line.Type {
		case git.LineAdded:
			styledContent = styles.DiffLineAdded.Width(lineContentWidth).Render(line.Content)
		case git.LineDeleted:
			styledContent = styles.DiffLineDeleted.Width(lineContentWidth).Render(line.Content)
		case git.LineHeader:
			styledContent = styles.DiffLineHeader.Width(lineContentWidth).Render(line.Content)
		default:
			styledContent = styles.DiffLineContext.Width(lineContentWidth).Render(line.Content)
		}

		renderedRows = append(renderedRows, lipgloss.JoinHorizontal(lipgloss.Left, numGutter, styledContent))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, title, separator, strings.Join(renderedRows, "\n"))
	return boxStyle.Width(width).Height(height).Render(content)
}
