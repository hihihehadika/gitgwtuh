package components

import (
	"fmt"
	"strings"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderDiffView renders the bottom-right syntax-highlighted code diff viewer.
func RenderDiffView(diffText string, scrollOffset, outerWidth, outerHeight, frame int, animEnabled, isActive bool) string {
	boxStyle := styles.PanelBoxInactive
	if isActive {
		boxStyle = styles.GetActivePanelBox(frame, animEnabled)
	}

	boxW := outerWidth - 2
	if boxW < 8 {
		boxW = 8
	}
	textW := boxW - 2
	if textW < 6 {
		textW = 6
	}
	boxH := outerHeight - 2
	if boxH < 1 {
		boxH = 1
	}

	title := styles.PanelTitle.Render("SYNTAX DIFF VIEWER")
	separator := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render(strings.Repeat("─", textW))

	if strings.TrimSpace(diffText) == "" {
		emptyMsg := lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Tidak ada perubahan kode untuk ditampilkan.")
		content := lipgloss.JoinVertical(lipgloss.Left, title, separator, "", emptyMsg)
		lines := strings.Split(content, "\n")
		if len(lines) > boxH {
			lines = lines[:boxH]
		}
		return boxStyle.Width(boxW).Height(boxH).Render(strings.Join(lines, "\n"))
	}

	parsedLines := git.ParseDiffOutput(diffText)
	visibleCount := boxH - 2
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

		lineContentWidth := textW - lipgloss.Width(numGutter) - 1
		if lineContentWidth < 5 {
			lineContentWidth = 5
		}

		var styledContent string
		switch line.Type {
		case git.LineAdded:
			styledContent = styles.DiffLineAdded.MaxWidth(lineContentWidth).Render(line.Content)
		case git.LineDeleted:
			styledContent = styles.DiffLineDeleted.MaxWidth(lineContentWidth).Render(line.Content)
		case git.LineHeader:
			styledContent = styles.DiffLineHeader.MaxWidth(lineContentWidth).Render(line.Content)
		default:
			styledContent = styles.DiffLineContext.MaxWidth(lineContentWidth).Render(line.Content)
		}

		renderedRows = append(renderedRows, lipgloss.JoinHorizontal(lipgloss.Left, numGutter, styledContent))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, title, separator, strings.Join(renderedRows, "\n"))
	lines := strings.Split(content, "\n")
	if len(lines) > boxH {
		lines = lines[:boxH]
	}
	content = strings.Join(lines, "\n")

	return boxStyle.Width(boxW).Height(boxH).Render(content)
}
