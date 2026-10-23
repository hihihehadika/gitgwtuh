package components

import (
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderFooter renders the bottom keybindings status bar.
func RenderFooter(activePaneName, statusMsg string, width int) string {
	modeBadge := styles.BrandBadge.Render("PANE: " + activePaneName)
	if statusMsg != "" {
		modeBadge = styles.StatusPillClean.Render(statusMsg)
	}

	key1 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("j/k"), styles.FooterDesc.Render("Navigasi"))
	key2 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("Tab"), styles.FooterDesc.Render("Panel"))
	key3 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("/"), styles.FooterDesc.Render("Cari"))
	key4 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("y"), styles.FooterDesc.Render("Salin SHA"))
	key5 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("a"), styles.FooterDesc.Render("Animasi"))
	key6 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("q"), styles.FooterDesc.Render("Keluar"))

	keyGroup := lipgloss.JoinHorizontal(lipgloss.Left, key1, key2, key3, key4, key5, key6)

	totalPillsWidth := lipgloss.Width(modeBadge) + lipgloss.Width(keyGroup)
	spacerWidth := width - totalPillsWidth
	if spacerWidth < 0 {
		spacerWidth = 0
	}

	spacer := lipgloss.NewStyle().Width(spacerWidth).Render("")
	footerLine := lipgloss.JoinHorizontal(lipgloss.Left, modeBadge, spacer, keyGroup)

	return lipgloss.NewStyle().MaxWidth(width).Render(footerLine)
}
