# gitgwtuh

[![CI](https://img.shields.io/github/actions/workflow/status/hihihehadika/gitgwtuh/ci.yml?branch=main&style=for-the-badge&logo=github&logoColor=white&label=CI&labelColor=0f172a&color=38bdf8)](https://github.com/hihihehadika/gitgwtuh/actions)
[![Go](https://img.shields.io/badge/Go-1.22+-38bdf8?style=for-the-badge&logo=go&logoColor=white&labelColor=0f172a)](https://golang.org)
[![TUI](https://img.shields.io/badge/TUI-Bubbletea-3b82f6?style=for-the-badge&logo=terminal&logoColor=white&labelColor=0f172a)](https://github.com/charmbracelet/bubbletea)

A high-aesthetic, keyboard-driven Git TUI explorer and interactive diff viewer built with Go.

```text
+-- [ gitgwtuh ] --------------------- [ * main ] -- [ 18 COMMITS ] -- [ REPO: gitgwtuh ] -+
|                                  |                                                         |
|  COMMIT TIMELINE                 |  +-- COMMIT DETAIL & FILES ---------------------------+ |
|  ─────────────────────────────── |  | HASH   : d110060938a1bc4e910248...                 | |
|  ● [d110060] ci: add github acti |  | AUTHOR : hihihehadika <hihihehadika@gmail.com>     | |
|  ○ [7e7097a] test(perf): add ben |  | DATE   : 2026-10-13 17:04 (2 minutes ago)          | |
|  ○ [bc91c0c] feat(ui): add commi |  | SUBJ   : ci: add github actions workflow           | |
|  ○ [cff4981] test(search): add u |  +----------------------------------------------------+ |
|  ○ [4bc079b] feat(search): imple |  +-- CHANGED FILES -----------------------------------+ |
|  ○ [6e5c910] build: add run.bat  |  | [A] .github/workflows/ci.yml   (+58  -0)           | |
|  ○ [c332046] feat(cli): create a |  +----------------------------------------------------+ |
|  ○ [ea0dc95] feat(ui): implement |  +-- SYNTAX DIFF VIEWER ------------------------------+ |
|  ○ [4938927] feat(ui): implement |  | @@ -0,0 +1,58 @@                                   | |
|  ○ [5749e9a] feat(ui): implement |  | +   name: CI                                       | |
|  ○ [25e9415] feat(ui): implement |  | +   on:                                            | |
|  ○ [7559d07] feat(ui): implement |  | +     push:                                        | |
|  ○ [4c596fe] feat(ui): implement |  | +       branches: [ main ]                         | |
|  ○ [6463da2] feat(ui): design cy |  +----------------------------------------------------+ |
+--------------------------------------------------------------------------------------------+
| [PANE: LOGS]   j/k: Navigasi | Tab: Panel | /: Cari | y: Salin SHA | q: Keluar             |
+--------------------------------------------------------------------------------------------+
```

---

## Features

- **Terminal-Native Tiling Layout:** Clean 3-panel display inspired by modern tiling window managers, showing your commit timeline, metadata inspector, and syntax-highlighted code diff simultaneously.
- **Deep Cosmic Cyber Theme:** Crafted with 24-bit TrueColor ANSI palettes featuring deep slate backgrounds, vibrant cyan accents, and clear status pill badges.
- **Interactive Search Filter:** Instant commit searching by message, author, or hash with `/` shortcut.
- **Clipboard SHA Yanking:** Copy full commit hashes directly to system clipboard using `y`.
- **Pure Go & Zero CGo:** Runs smoothly on Windows Terminal, PowerShell, and Unix shells with zero native library headaches.
- **Vim-Style Keybindings:** Navigate effortlessly using standard keyboard shortcuts (`j`, `k`, `Tab`, `Enter`, `/`, `y`, `q`).
- **Dynamic Resize Engine:** Automatically recalculates panel bounds on window resize events to keep borders and text crisp.
- **Safe Read-Only Operation:** Inspects commits, branches, and diffs without modifying or risking your repository state.

---

## Keybindings

| Key | Action |
|-----|--------|
| `j` / `Down` | Move down to next commit / file / diff line |
| `k` / `Up` | Move up to previous commit / file / diff line |
| `Tab` | Switch active panel focus (`LOGS` <-> `FILES` <-> `DIFF`) |
| `Enter` | Expand commit details or view file-specific diff |
| `/` | Open live search bar to filter commits |
| `y` | Copy (yank) selected commit hash to clipboard |
| `Esc` | Clear search filter or dismiss active prompt |
| `q` / `Ctrl+C` | Quit and restore terminal screen cleanly |

---

## Quick Start

### Run from Source

```bash
git clone https://github.com/hihihehadika/gitgwtuh.git
cd gitgwtuh
go run ./cmd/gitgwtuh
```

### Build Standalone Executable

```bash
go build -ldflags="-s -w" -o gitgwtuh.exe ./cmd/gitgwtuh
./gitgwtuh.exe
```

---

## Architecture

```
gitgwtuh/
├── cmd/gitgwtuh/         CLI entrypoint and alternate screen lifecycle
├── internal/
│   ├── git/              Git command driver, log parser, and unified diff engine
│   └── ui/
│       ├── styles/       Cosmic theme tokens, colors, borders, and badge styles
│       ├── components/   Header telemetry, commit list, detail card, and diff viewport
│       └── app.go        Bubbletea Model-Update-View application controller
└── README.md
```
