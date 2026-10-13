package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme colors adhering to the cyber cosmic dark palette.
var (
	ColorPrimary   = lipgloss.Color("#38bdf8") // Cyan neon
	ColorElectric  = lipgloss.Color("#3b82f6") // Electric blue
	ColorDarkBg    = lipgloss.Color("#090d16") // Deep void
	ColorPanelBg   = lipgloss.Color("#0f172a") // Slate dark panel
	ColorBorder    = lipgloss.Color("#334155") // Muted border
	ColorBorderAct = lipgloss.Color("#38bdf8") // Active border cyan
	ColorMuted     = lipgloss.Color("#64748b") // Gray text
	ColorLight     = lipgloss.Color("#f8fafc") // Bright text
	ColorAdded     = lipgloss.Color("#10b981") // Emerald neon (+)
	ColorDeleted   = lipgloss.Color("#ef4444") // Crimson rose (-)
	ColorHeader    = lipgloss.Color("#f59e0b") // Amber @@ headers
)

// UI Component styles.
var (
	// Brand and header pills
	BrandBadge = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorDarkBg).
			Background(ColorPrimary).
			Padding(0, 1).
			MarginRight(1)

	BranchPill = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			Background(lipgloss.Color("#1e293b")).
			Padding(0, 1).
			MarginRight(1)

	StatusPillClean = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAdded).
			Background(lipgloss.Color("#064e3b")).
			Padding(0, 1).
			MarginRight(1)

	StatusPillDirty = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorDeleted).
			Background(lipgloss.Color("#7f1d1d")).
			Padding(0, 1).
			MarginRight(1)

	StatPill = lipgloss.NewStyle().
			Foreground(ColorLight).
			Background(lipgloss.Color("#1e293b")).
			Padding(0, 1).
			MarginRight(1)

	// Panel boxes
	PanelBoxInactive = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorBorder).
				Background(ColorPanelBg).
				Padding(0, 1)

	PanelBoxActive = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorderAct).
			Background(ColorPanelBg).
			Padding(0, 1)

	PanelTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginBottom(1)

	// Commit list items
	CommitItemNormal = lipgloss.NewStyle().
				Foreground(ColorLight)

	CommitItemSelected = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorDarkBg).
				Background(ColorPrimary)

	CommitHashPill = lipgloss.NewStyle().
			Foreground(ColorElectric).
			Bold(true)

	CommitDateText = lipgloss.NewStyle().
			Foreground(ColorMuted)

	// Diff Viewer styles
	DiffLineContext = lipgloss.NewStyle().
			Foreground(ColorMuted)

	DiffLineAdded = lipgloss.NewStyle().
			Foreground(ColorAdded).
			Background(lipgloss.Color("#052e16"))

	DiffLineDeleted = lipgloss.NewStyle().
			Foreground(ColorDeleted).
			Background(lipgloss.Color("#450a0a"))

	DiffLineHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorHeader).
			Background(lipgloss.Color("#1e293b"))

	DiffLineNum = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#475569")).
			Width(5).
			Align(lipgloss.Right).
			MarginRight(1)

	// Footer status bar
	FooterKey = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorDarkBg).
			Background(ColorPrimary).
			Padding(0, 1)

	FooterDesc = lipgloss.NewStyle().
			Foreground(ColorLight).
			Background(lipgloss.Color("#1e293b")).
			Padding(0, 1).
			MarginRight(1)
)
