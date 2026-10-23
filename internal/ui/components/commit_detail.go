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

	title := styles.PanelTitle.Render("COMMIT DETAIL & FILES")
	separator := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render(strings.Repeat("─", textW))

	if commit == nil {
		emptyMsg := lipgloss.NewStyle().Foreground(styles.ColorMuted).Render("Pilih commit di panel kiri untuk melihat detail.")
		content := lipgloss.JoinVertical(lipgloss.Left, title, separator, "", emptyMsg)
		lines := strings.Split(content, "\n")
		if len(lines) > boxH {
			lines = lines[:boxH]
		}
		return boxStyle.Width(boxW).Height(boxH).Render(strings.Join(lines, "\n"))
	}

	metaLabelStyle := lipgloss.NewStyle().Foreground(styles.ColorPrimary).Bold(true).Width(8)
	metaValStyle := lipgloss.NewStyle().Foreground(styles.ColorLight)

	hashVal := commit.Hash
	if len(hashVal) > textW-10 && textW > 14 {
		hashVal = hashVal[:textW-13] + "..."
	}
	hashRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("HASH   :"), metaValStyle.Render(hashVal))

	authorVal := fmt.Sprintf("%s <%s>", commit.Author, commit.Email)
	if len(authorVal) > textW-10 && textW > 14 {
		authorVal = authorVal[:textW-13] + "..."
	}
	authorRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("AUTHOR :"), metaValStyle.Render(authorVal))

	dateVal := fmt.Sprintf("%s (%s)", commit.Date, commit.RelativeDate)
	if len(dateVal) > textW-10 && textW > 14 {
		dateVal = dateVal[:textW-13] + "..."
	}
	dateRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("DATE   :"), metaValStyle.Render(dateVal))

	subjVal := commit.Subject
	if len(subjVal) > textW-10 && textW > 14 {
		subjVal = subjVal[:textW-13] + "..."
	}
	subjRow := fmt.Sprintf("%s %s", metaLabelStyle.Render("SUBJ   :"), metaValStyle.Render(subjVal))

	metaBlock := lipgloss.JoinVertical(lipgloss.Left, hashRow, authorRow, dateRow, subjRow)

	var contentBlocks = []string{title, separator, metaBlock}
	if boxH >= 8 {
		maxFiles := boxH - 7
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
				availPathLen := textW - fixedLen
				pathStr := f.Path
				if availPathLen > 4 && len(pathStr) > availPathLen {
					pathStr = "..." + pathStr[len(pathStr)-(availPathLen-3):]
				}

				line := fmt.Sprintf("%s %s %s", statusTag, pathStr, statNum)
				if i == selectedFileIndex && isActive {
					fileRows = append(fileRows, styles.CommitItemSelected.MaxWidth(textW).Render(line))
				} else {
					fileRows = append(fileRows, styles.CommitItemNormal.MaxWidth(textW).Render(line))
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
	if len(lines) > boxH {
		lines = lines[:boxH]
	}
	content = strings.Join(lines, "\n")

	return boxStyle.Width(boxW).Height(boxH).Render(content)
}
