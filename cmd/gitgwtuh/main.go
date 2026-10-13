package main

import (
	"fmt"
	"os"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui"
	"gitgwtuh/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FATAL] Gagal mendeteksi direktori aktif: %v\n", err)
		os.Exit(1)
	}

	gitClient := git.NewClient(workDir)
	if !gitClient.IsGitRepository() {
		errorCard := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(styles.ColorDeleted).
			Padding(1, 2).
			Align(lipgloss.Center).
			Render("gitgwtuh\n\nDirektori ini bukan repositori Git yang valid.\nJalankan di dalam folder proyek Git Anda.")

		fmt.Println(errorCard)
		os.Exit(1)
	}

	model := ui.NewAppModel(gitClient)
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Terjadi kegagalan saat menjalankan TUI: %v\n", err)
		os.Exit(1)
	}
}
