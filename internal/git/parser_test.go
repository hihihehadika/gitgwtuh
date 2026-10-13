package git

import (
	"testing"
)

func TestParseLogOutput(t *testing.T) {
	raw := "\x00dcbdb0f1234567890abcdef1234567890abcdef1\x00dcbdb0f\x00hihihehadika\x00dika@example.com\x002026-10-13 16:42\x002 minutes ago\x00chore(init): initialize gitgwtuh\x00Initial commit body"

	commits := ParseLogOutput(raw)
	if len(commits) != 1 {
		t.Fatalf("expected 1 commit, got %d", len(commits))
	}

	c := commits[0]
	if c.ShortHash != "dcbdb0f" {
		t.Errorf("expected short hash 'dcbdb0f', got '%s'", c.ShortHash)
	}
	if c.Author != "hihihehadika" {
		t.Errorf("expected author 'hihihehadika', got '%s'", c.Author)
	}
	if c.Subject != "chore(init): initialize gitgwtuh" {
		t.Errorf("expected subject 'chore(init): initialize gitgwtuh', got '%s'", c.Subject)
	}
	if c.Body != "Initial commit body" {
		t.Errorf("expected body 'Initial commit body', got '%s'", c.Body)
	}
}

func TestParseNumstatOutput(t *testing.T) {
	raw := "10\t2\tinternal/git/client.go\n5\t0\tREADME.md\n0\t4\tdeleted.txt\n"

	changes := ParseNumstatOutput(raw)
	if len(changes) != 3 {
		t.Fatalf("expected 3 file changes, got %d", len(changes))
	}

	if changes[0].Path != "internal/git/client.go" || changes[0].Additions != 10 || changes[0].Deletions != 2 {
		t.Errorf("unexpected change record: %+v", changes[0])
	}

	if changes[1].Status != StatusAdded {
		t.Errorf("expected StatusAdded for README.md, got %v", changes[1].Status)
	}

	if changes[2].Status != StatusDeleted {
		t.Errorf("expected StatusDeleted for deleted.txt, got %v", changes[2].Status)
	}
}

func TestParseDiffOutput(t *testing.T) {
	raw := `diff --git a/main.go b/main.go
index 1234567..89abcdef 100644
--- a/main.go
+++ b/main.go
@@ -1,3 +1,4 @@
 package main
+import "fmt"
-func old() {}
+func new() {}
`

	lines := ParseDiffOutput(raw)
	if len(lines) == 0 {
		t.Fatal("expected parsed diff lines, got none")
	}

	var hasHeader, hasAdded, hasDeleted bool
	for _, l := range lines {
		if l.Type == LineHeader {
			hasHeader = true
		}
		if l.Type == LineAdded {
			hasAdded = true
		}
		if l.Type == LineDeleted {
			hasDeleted = true
		}
	}

	if !hasHeader || !hasAdded || !hasDeleted {
		t.Errorf("diff parser missing elements: header=%v, added=%v, deleted=%v", hasHeader, hasAdded, hasDeleted)
	}
}
