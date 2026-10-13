package components

import (
	"fmt"
	"strings"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderCommitDetail renders the top-right card containing commit metadata and changed files list.
func RenderCommitDetail(commit *git.Commit, files []git.FileChange, selectedFileIndex, width, height int, isActive bool) string {
	boxStyle := styles.PanelBoxInactive
	if isActive {
		boxStyle = styles.PanelBoxActive
	}

	contentWidth := width - 4
	if contentWidth < 10 {
		contentWidth = 10
	}

	title := styles.PanelTitle.Render("COMMIT DETAIL & FILES")
	separator := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render(strings.Repeat("─", contentWidth))

	if commit == nil {
		emptyMsg := lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Pilih commit di panel kiri untuk melihat detail.")
		content := lipgloss.JoinVertical(lipgloss.Left, title, separator, "", emptyMsg)
		return boxStyle.Width(width).Height(height).Render(content)
	}

	metaLabelStyle := lipgloss.NewStyle().Foreground(styles.ColorPrimary).Bold(true).Width(8)
	metaValStyle := lipgloss.NewStyle().Foreground(styles.ColorLight)

	hashRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("HASH   :"), metaValStyle.Render(commit.Hash))
	authorRow := fmt.Sprintf("%s %s <%s>", metaLabelStyle.Render("AUTHOR :"), metaValStyle.Render(commit.Author), commit.Email)
	dateRow := fmt.Sprintf("%s %s (%s)", metaLabelStyle.Render("DATE   :"), metaValStyle.Render(commit.Date), commit.RelativeDate)
	subjRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("SUBJ   :"), metaValStyle.Render(commit.Subject))

	metaBlock := lipgloss.JoinVertical(lipgloss.Left, hashRow, authorRow, dateRow, subjRow)

	filesTitle := lipgloss.NewStyle().Foreground(styles.ColorPrimary).Bold(true).Render("CHANGED FILES:")
	var fileRows []string
	if len(files) == 0 {
		fileRows = append(fileRows, lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Tidak ada file perubahan."))
	} else {
		for i, f := range files {
			statusTag := fmt.Sprintf("[%s]", f.Status)
			statNum := fmt.Sprintf("(+%d -%d)", f.Additions, f.Deletions)

			line := fmt.Sprintf("%s %s %s", statusTag, f.Path, statNum)
			if i == selectedFileIndex && isActive {
				fileRows = append(fileRows, styles.CommitItemSelected.Width(contentWidth).Render(line))
			} else {
				fileRows = append(fileRows, styles.CommitItemNormal.Width(contentWidth).Render(line))
			}
		}
	}

	filesBlock := lipgloss.JoinVertical(lipgloss.Left, filesTitle, strings.Join(fileRows, "\n"))
	content := lipgloss.JoinVertical(lipgloss.Left, title, separator, metaBlock, "", filesBlock)

	return boxStyle.Width(width).Height(height).Render(content)
}
