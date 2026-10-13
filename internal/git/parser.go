package git

import (
	"bufio"
	"strconv"
	"strings"
)

// ParseLogOutput parses raw git log output separated by null delimiters.
// The format expected is:
// \x00<Hash>\x00<ShortHash>\x00<Author>\x00<Email>\x00<Date>\x00<RelativeDate>\x00<Subject>\x00<Body>
func ParseLogOutput(raw string) []Commit {
	var commits []Commit
	if len(raw) == 0 {
		return commits
	}

	records := strings.Split(raw, "\x00")
	// Each commit requires 8 fields preceded by an initial delimiter
	for i := 1; i+7 < len(records); i += 8 {
		c := Commit{
			Hash:         strings.TrimSpace(records[i]),
			ShortHash:    strings.TrimSpace(records[i+1]),
			Author:       strings.TrimSpace(records[i+2]),
			Email:        strings.TrimSpace(records[i+3]),
			Date:         strings.TrimSpace(records[i+4]),
			RelativeDate: strings.TrimSpace(records[i+5]),
			Subject:      strings.TrimSpace(records[i+6]),
			Body:         strings.TrimSpace(records[i+7]),
		}
		if c.Hash != "" {
			commits = append(commits, c)
		}
	}

	return commits
}

// ParseNumstatOutput parses git show --numstat output into FileChange structs.
func ParseNumstatOutput(raw string) []FileChange {
	var changes []FileChange
	scanner := bufio.NewScanner(strings.NewReader(raw))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}

		adds, _ := strconv.Atoi(parts[0])
		dels, _ := strconv.Atoi(parts[1])
		filePath := parts[2]

		status := StatusModified
		if adds > 0 && dels == 0 {
			status = StatusAdded
		} else if adds == 0 && dels > 0 {
			status = StatusDeleted
		}

		changes = append(changes, FileChange{
			Path:      filePath,
			Status:    status,
			Additions: adds,
			Deletions: dels,
		})
	}

	return changes
}

// ParseDiffOutput parses unified diff text into structured lines with line numbering.
func ParseDiffOutput(raw string) []DiffLine {
	var diffLines []DiffLine
	scanner := bufio.NewScanner(strings.NewReader(raw))

	oldLine := 0
	newLine := 0

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "@@") {
			diffLines = append(diffLines, DiffLine{
				Type:    LineHeader,
				Content: text,
			})
			continue
		}

		if strings.HasPrefix(text, "+") && !strings.HasPrefix(text, "+++") {
			newLine++
			diffLines = append(diffLines, DiffLine{
				Type:      LineAdded,
				Content:   text,
				NewLineNo: newLine,
			})
		} else if strings.HasPrefix(text, "-") && !strings.HasPrefix(text, "---") {
			oldLine++
			diffLines = append(diffLines, DiffLine{
				Type:      LineDeleted,
				Content:   text,
				OldLineNo: oldLine,
			})
		} else {
			oldLine++
			newLine++
			diffLines = append(diffLines, DiffLine{
				Type:      LineContext,
				Content:   text,
				OldLineNo: oldLine,
				NewLineNo: newLine,
			})
		}
	}

	return diffLines
}
