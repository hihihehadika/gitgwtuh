package ui

import (
	"strings"

	"gitgwtuh/internal/git"
	"gitgwtuh/internal/ui/components"
	"gitgwtuh/internal/ui/styles"

	"github.com/charmbracelet/bubbles/textinput"
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

	summary     git.RepoSummary
	allCommits  []git.Commit
	commits     []git.Commit
	commitIdx   int
	commitOff   int

	files   []git.FileChange
	fileIdx int

	diff    string
	diffOff int

	activePane  Pane
	width       int
	height      int
	err         error

	isSearching bool
	searchInput textinput.Model
	statusMsg   string
}

// NewAppModel initializes the AppModel with a Git client.
func NewAppModel(client *git.Client) *AppModel {
	ti := textinput.New()
	ti.Placeholder = "Ketik untuk mencari commit (Esc untuk batal)..."
	ti.Prompt = "CARI: "
	ti.PromptStyle = styles.BrandBadge
	ti.TextStyle = styles.CommitItemNormal

	return &AppModel{
		gitClient:   client,
		activePane:  PaneCommits,
		width:       100,
		height:      30,
		searchInput: ti,
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
		m.allCommits = msg.commits
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
		if m.isSearching {
			switch msg.String() {
			case "esc":
				m.isSearching = false
				m.searchInput.Blur()
				m.searchInput.Reset()
				m.commits = m.allCommits
				m.commitIdx = 0
				m.commitOff = 0
				return m, nil
			case "enter":
				m.isSearching = false
				m.searchInput.Blur()
				return m, nil
			default:
				var cmd tea.Cmd
				m.searchInput, cmd = m.searchInput.Update(msg)
				m.filterCommits(m.searchInput.Value())
				return m, cmd
			}
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab":
			m.activePane = (m.activePane + 1) % 3
			return m, nil

		case "/":
			m.isSearching = true
			m.searchInput.Focus()
			return m, textinput.Blink

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

func (m *AppModel) filterCommits(query string) {
	if strings.TrimSpace(query) == "" {
		m.commits = m.allCommits
		m.commitIdx = 0
		m.commitOff = 0
		return
	}

	lowerQuery := strings.ToLower(query)
	var filtered []git.Commit
	for _, c := range m.allCommits {
		if strings.Contains(strings.ToLower(c.Subject), lowerQuery) ||
			strings.Contains(strings.ToLower(c.Author), lowerQuery) ||
			strings.Contains(strings.ToLower(c.ShortHash), lowerQuery) {
			filtered = append(filtered, c)
		}
	}

	m.commits = filtered
	m.commitIdx = 0
	m.commitOff = 0

	if len(m.commits) > 0 {
		hash := m.commits[0].Hash
		m.files, _ = m.gitClient.GetCommitFiles(hash)
		m.diff, _ = m.gitClient.GetCommitDiff(hash, "")
	} else {
		m.files = nil
		m.diff = ""
	}
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

	var footer string
	if m.isSearching {
		footer = m.searchInput.View()
	} else {
		paneName := "LOGS"
		if m.activePane == PaneFiles {
			paneName = "FILES"
		} else if m.activePane == PaneDiff {
			paneName = "DIFF"
		}
		footer = components.RenderFooter(paneName, m.width)
	}

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
