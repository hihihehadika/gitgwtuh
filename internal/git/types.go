package git

// Commit represents a single Git commit record with parsed metadata.
type Commit struct {
	Hash         string
	ShortHash    string
	Author       string
	Email        string
	Date         string
	RelativeDate string
	Subject      string
	Body         string
}

// FileChangeStatus represents the modification status of a file in a commit.
type FileChangeStatus string

const (
	StatusModified FileChangeStatus = "M"
	StatusAdded    FileChangeStatus = "A"
	StatusDeleted  FileChangeStatus = "D"
	StatusRenamed  FileChangeStatus = "R"
	StatusUnknown  FileChangeStatus = "?"
)

// FileChange represents a file affected by a specific commit.
type FileChange struct {
	Path      string
	Status    FileChangeStatus
	Additions int
	Deletions int
}

// LineType defines the semantic type of a line within a code diff.
type LineType int

const (
	LineContext LineType = iota
	LineAdded
	LineDeleted
	LineHeader
)

// DiffLine represents a single line inside a diff hunk.
type DiffLine struct {
	Type      LineType
	Content   string
	OldLineNo int
	NewLineNo int
}

// DiffHunk represents a unified diff chunk with a header and slice of lines.
type DiffHunk struct {
	Header string
	Lines  []DiffLine
}

// FileDiff represents the full parsed diff for a specific file.
type FileDiff struct {
	Path    string
	OldPath string
	Hunks   []DiffHunk
	Raw     string
}

// RepoSummary contains high-level telemetry data for the top header.
type RepoSummary struct {
	Name          string
	CurrentBranch string
	TotalCommits  int
	TotalAuthors  int
	IsClean       bool
}
