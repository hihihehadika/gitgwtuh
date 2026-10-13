package ui

import (
	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/components"
	"gitgwtuh/internal/ui/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Pane int

const (
	PaneCommits Pane = iota
	PaneFiles
	PaneDiff
)

// AppModel represents the entire state of the gitgwtuh TUI application.
type AppModel struct {
	gitClient *git.Client

	summary   git.RepoSummary
	commits   []git.Commit
	commitIdx int
	commitOff int

	files   []git.FileChange
	fileIdx int

	diff    string
	diffOff int

	activePane Pane
	width      int
	height     int
	err        error
}

// NewAppModel initializes the AppModel with a Git client.
func NewAppModel(client *git.Client) *AppModel {
	return &AppModel{
		gitClient:  client,
		activePane: PaneCommits,
		width:      100,
		height:     30,
	}
}

// Init triggers initial data loading when Bubbletea begins.
func (m *AppModel) Init() tea.Cmd {
	return m.loadInitialData
}

// loadInitialData fetches summary, commit log, and initial commit detail.
func (m *AppModel) loadInitialData() tea.Msg {
	summary, err := m.gitClient.GetRepoSummary()
	if err != nil {
		return errMsg{err}
	}

	commits, err := m.gitClient.GetCommits(100, 0)
	if err != nil {
		return errMsg{err}
	}

	var files []git.FileChange
	var diff string
	if len(commits) > 0 {
		files, _ = m.gitClient.GetCommitFiles(commits[0].Hash)
		diff, _ = m.gitClient.GetCommitDiff(commits[0].Hash, "")
	}

	return initialDataMsg{
		summary: summary,
		commits: commits,
		files:   files,
		diff:    diff,
	}
}

type initialDataMsg struct {
	summary git.RepoSummary
	commits []git.Commit
	files   []git.FileChange
	diff    string
}

type commitLoadedMsg struct {
	files []git.FileChange
	diff  string
}

type fileDiffLoadedMsg struct {
	diff string
}

type errMsg struct{ err error }

// Update processes Bubbletea events and state mutations.
func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case initialDataMsg:
		m.summary = msg.summary
		m.commits = msg.commits
		m.files = msg.files
		m.diff = msg.diff
		m.commitIdx = 0
		m.fileIdx = 0
		return m, nil

	case commitLoadedMsg:
		m.files = msg.files
		m.diff = msg.diff
		m.fileIdx = 0
		m.diffOff = 0
		return m, nil

	case fileDiffLoadedMsg:
		m.diff = msg.diff
		m.diffOff = 0
		return m, nil

	case errMsg:
		m.err = msg.err
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "tab":
			m.activePane = (m.activePane + 1) % 3
			return m, nil

		case "j", "down":
			return m.handleNavDown()

		case "k", "up":
			return m.handleNavUp()

		case "enter":
			if m.activePane == PaneFiles && len(m.files) > 0 {
				hash := m.commits[m.commitIdx].Hash
				path := m.files[m.fileIdx].Path
				return m, func() tea.Msg {
					diff, _ := m.gitClient.GetCommitDiff(hash, path)
					return fileDiffLoadedMsg{diff: diff}
				}
			}
		}
	}

	return m, nil
}

func (m *AppModel) handleNavDown() (tea.Model, tea.Cmd) {
	switch m.activePane {
	case PaneCommits:
		if m.commitIdx < len(m.commits)-1 {
			m.commitIdx++
			// Auto scroll
			visibleHeight := m.height - 6
			if m.commitIdx >= m.commitOff+visibleHeight {
				m.commitOff = m.commitIdx - visibleHeight + 1
			}
			hash := m.commits[m.commitIdx].Hash
			return m, func() tea.Msg {
				files, _ := m.gitClient.GetCommitFiles(hash)
				diff, _ := m.gitClient.GetCommitDiff(hash, "")
				return commitLoadedMsg{files: files, diff: diff}
			}
		}
	case PaneFiles:
		if m.fileIdx < len(m.files)-1 {
			m.fileIdx++
		}
	case PaneDiff:
		m.diffOff += 2
	}
	return m, nil
}

func (m *AppModel) handleNavUp() (tea.Model, tea.Cmd) {
	switch m.activePane {
	case PaneCommits:
		if m.commitIdx > 0 {
			m.commitIdx--
			// Auto scroll
			if m.commitIdx < m.commitOff {
				m.commitOff = m.commitIdx
			}
			hash := m.commits[m.commitIdx].Hash
			return m, func() tea.Msg {
				files, _ := m.gitClient.GetCommitFiles(hash)
				diff, _ := m.gitClient.GetCommitDiff(hash, "")
				return commitLoadedMsg{files: files, diff: diff}
			}
		}
	case PaneFiles:
		if m.fileIdx > 0 {
			m.fileIdx--
		}
	case PaneDiff:
		if m.diffOff > 0 {
			m.diffOff -= 2
			if m.diffOff < 0 {
				m.diffOff = 0
			}
		}
	}
	return m, nil
}

// View renders the terminal application UI with responsive layout calculations.
func (m *AppModel) View() string {
	// Guard against tiny terminal windows
	if m.width < 75 || m.height < 18 {
		warningBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(styles.ColorPrimary).
			Padding(1, 2).
			Align(lipgloss.Center).
			Render("gitgwtuh\n\nPerlebar jendela terminal Anda\nuntuk tampilan optimal (min 75x18).")
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, warningBox)
	}

	header := components.RenderHeader(m.summary, m.width)

	paneName := "LOGS"
	if m.activePane == PaneFiles {
		paneName = "FILES"
	} else if m.activePane == PaneDiff {
		paneName = "DIFF"
	}
	footer := components.RenderFooter(paneName, m.width)

	// Available height for body panels
	bodyHeight := m.height - 4
	if bodyHeight < 10 {
		bodyHeight = 10
	}

	// Columns split: Left (40%), Right (60%)
	leftWidth := (m.width * 4) / 10
	rightWidth := m.width - leftWidth - 1

	var activeCommit *git.Commit
	if len(m.commits) > 0 && m.commitIdx < len(m.commits) {
		activeCommit = &m.commits[m.commitIdx]
	}

	leftPanel := components.RenderCommitList(
		m.commits,
		m.commitIdx,
		m.commitOff,
		leftWidth,
		bodyHeight,
		m.activePane == PaneCommits,
	)

	// Right column vertical split: Detail (38%), Diff (62%)
	detailHeight := (bodyHeight * 38) / 100
	diffHeight := bodyHeight - detailHeight - 1

	rightDetail := components.RenderCommitDetail(
		activeCommit,
		m.files,
		m.fileIdx,
		rightWidth,
		detailHeight,
		m.activePane == PaneFiles,
	)

	rightDiff := components.RenderDiffView(
		m.diff,
		m.diffOff,
		rightWidth,
		diffHeight,
		m.activePane == PaneDiff,
	)

	rightColumn := lipgloss.JoinVertical(lipgloss.Left, rightDetail, rightDiff)
	body := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightColumn)

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}
