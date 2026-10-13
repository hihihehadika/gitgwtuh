package components

import (
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/lipgloss"
)

// RenderFooter renders the bottom keybindings status bar.
func RenderFooter(activePaneName string, width int) string {
	modeBadge := styles.BrandBadge.Render("PANE: " + activePaneName)

	key1 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("j/k"), styles.FooterDesc.Render("Navigasi"))
	key2 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("Tab"), styles.FooterDesc.Render("Pindah Panel"))
	key3 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("Enter"), styles.FooterDesc.Render("Buka Diff"))
	key4 := lipgloss.JoinHorizontal(lipgloss.Left, styles.FooterKey.Render("q/Esc"), styles.FooterDesc.Render("Keluar"))

	keyGroup := lipgloss.JoinHorizontal(lipgloss.Center, key1, key2, key3, key4)

	totalPillsWidth := lipgloss.Width(modeBadge) + lipgloss.Width(keyGroup)
	spacerWidth := width - totalPillsWidth
	if spacerWidth < 1 {
		spacerWidth = 1
	}

	spacer := lipgloss.NewStyle().Width(spacerWidth).Render("")

	return lipgloss.JoinHorizontal(lipgloss.Center, modeBadge, spacer, keyGroup)
}
