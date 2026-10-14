package components

import (
	"fmt"
	"strings"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderCommitDetail renders the top-right card containing commit metadata and changed files list.
func RenderCommitDetail(commit *git.Commit, files []git.FileChange, selectedFileIndex, outerWidth, outerHeight, frame int, animEnabled, isActive bool) string {
	boxStyle := styles.PanelBoxInactive
	if isActive {
		boxStyle = styles.GetActivePanelBox(frame, animEnabled)
	}

	boxWidth := outerWidth - 2
	if boxWidth < 10 {
		boxWidth = 10
	}
	textWidth := boxWidth - 2
	if textWidth < 8 {
		textWidth = 8
	}

	boxHeight := outerHeight - 2
	if boxHeight < 1 {
		boxHeight = 1
	}

	title := styles.PanelTitle.Render("COMMIT DETAIL & FILES")
	separator := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render(strings.Repeat("─", textWidth))

	if commit == nil {
		emptyMsg := lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Pilih commit di panel kiri untuk melihat detail.")
		content := lipgloss.JoinVertical(lipgloss.Left, title, separator, "", emptyMsg)
		lines := strings.Split(content, "\n")
		if len(lines) > boxHeight {
			lines = lines[:boxHeight]
		}
		return boxStyle.Width(boxWidth).Height(boxHeight).Render(strings.Join(lines, "\n"))
	}

	metaLabelStyle := lipgloss.NewStyle().Foreground(styles.ColorPrimary).Bold(true).Width(8)
	metaValStyle := lipgloss.NewStyle().Foreground(styles.ColorLight)

	hashVal := commit.Hash
	if len(hashVal) > textWidth-10 && textWidth > 14 {
		hashVal = hashVal[:textWidth-13] + "..."
	}
	hashRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("HASH   :"), metaValStyle.Render(hashVal))

	authorVal := fmt.Sprintf("%s <%s>", commit.Author, commit.Email)
	if len(authorVal) > textWidth-10 && textWidth > 14 {
		authorVal = authorVal[:textWidth-13] + "..."
	}
	authorRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("AUTHOR :"), metaValStyle.Render(authorVal))

	dateVal := fmt.Sprintf("%s (%s)", commit.Date, commit.RelativeDate)
	if len(dateVal) > textWidth-10 && textWidth > 14 {
		dateVal = dateVal[:textWidth-13] + "..."
	}
	dateRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("DATE   :"), metaValStyle.Render(dateVal))

	subjVal := commit.Subject
	if len(subjVal) > textWidth-10 && textWidth > 14 {
		subjVal = subjVal[:textWidth-13] + "..."
	}
	subjRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("SUBJ   :"), metaValStyle.Render(subjVal))

	metaBlock := lipgloss.JoinVertical(lipgloss.Left, hashRow, authorRow, dateRow, subjRow)

	var contentBlocks = []string{title, separator, metaBlock}
	if boxHeight >= 8 {
		maxFiles := boxHeight - 7
		var fileRows []string
		if len(files) == 0 {
			fileRows = append(fileRows, lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Tidak ada file perubahan."))
		} else {
			limit := len(files)
			if limit > maxFiles {
				limit = maxFiles
			}
			for i := 0; i < limit; i++ {
				f := files[i]
				statusTag := fmt.Sprintf("[%s]", f.Status)
				statNum := fmt.Sprintf("(+%d -%d)", f.Additions, f.Deletions)

				fixedLen := len(statusTag) + 1 + len(statNum) + 1
				availPathLen := textWidth - fixedLen
				pathStr := f.Path
				if availPathLen > 4 && len(pathStr) > availPathLen {
					pathStr = "..." + pathStr[len(pathStr)-(availPathLen-3):]
				}

				line := fmt.Sprintf("%s %s %s", statusTag, pathStr, statNum)
				if i == selectedFileIndex && isActive {
					fileRows = append(fileRows, styles.CommitItemSelected.MaxWidth(textWidth).Render(line))
				} else {
					fileRows = append(fileRows, styles.CommitItemNormal.MaxWidth(textWidth).Render(line))
				}
			}
			if len(files) > maxFiles && maxFiles > 0 {
				moreCount := len(files) - maxFiles
				fileRows[len(fileRows)-1] = lipgloss.NewStyle().Foreground(styles.ColorMuted).Render(fmt.Sprintf("... dan %d file lainnya", moreCount+1))
			}
		}
		filesTitle := lipgloss.NewStyle().Foreground(styles.ColorPrimary).Bold(true).Render("CHANGED FILES:")
		filesBlock := lipgloss.JoinVertical(lipgloss.Left, filesTitle, strings.Join(fileRows, "\n"))
		contentBlocks = append(contentBlocks, "", filesBlock)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, contentBlocks...)
	lines := strings.Split(content, "\n")
	if len(lines) > boxHeight {
		lines = lines[:boxHeight]
	}
	content = strings.Join(lines, "\n")

	return boxStyle.Width(boxWidth).Height(boxHeight).Render(content)
}
