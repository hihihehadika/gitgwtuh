# gitgwtuh

[![CI](https://img.shields.io/github/actions/workflow/status/hihihehadika/gitgwtuh/ci.yml?branch=main&style=for-the-badge&logo=github&logoColor=white&label=CI&labelColor=0f172a&color=38bdf8)](https://github.com/hihihehadika/gitgwtuh/actions)
[![Go](https://img.shields.io/badge/Go-1.22+-38bdf8?style=for-the-badge&logo=go&logoColor=white&labelColor=0f172a)](https://golang.org)
[![TUI](https://img.shields.io/badge/TUI-Bubbletea-3b82f6?style=for-the-badge&logo=terminal&logoColor=white&labelColor=0f172a)](https://github.com/charmbracelet/bubbletea)

A high-aesthetic, keyboard-driven Git TUI explorer and interactive diff viewer built with Go.

```text
+-- [ gitgwtuh ] [ ▄▆█▆▄ ] [ * main ] ------------- [ 28 COMMITS ] -- [ REPO: gitgwtuh ] -+
|                                  |                                                         |
|  COMMIT TIMELINE                 |  +-- COMMIT DETAIL & FILES ---------------------------+ |
|  ─────────────────────────────── |  | HASH   : 063e5ff38a1bc4e910248...                  | |
|  ● [063e5ff] test: add animatio  |  | AUTHOR : hihihehadika <hihihehadika@gmail.com>      | |
|  ○ [a5126a0] feat(app): add ani  |  | DATE   : 2026-10-14 18:43 (just now)               | |
|  ○ [f96b2a0] feat(footer): add   |  | SUBJ   : test: add animation unit tests            | |
|  ○ [26103be] fix(diff-view): fi  |  +----------------------------------------------------+ |
|  ○ [063d1a6] fix(commit-detail)  |  +-- CHANGED FILES -----------------------------------+ |
|  ○ [576a16d] fix(commit-list):   |  | [M] internal/ui/app_test.go         (+58  -0)      | |
|  ○ [c659d5c] feat(header): add   |  +----------------------------------------------------+ |
|  ○ [3d96fa6] feat(theme): add B  |  +-- SYNTAX DIFF VIEWER ------------------------------+ |
|  ○ [84848c7] docs: remove relea  |  | @@ -0,0 +1,58 @@                                  | |
|  ○ [df2431f] style(docs): moder  |  | +   func TestAnimationToggle(t *testing.T) {       | |
|                                  |  | +     model := NewAppModel(client)                  | |
|                                  |  | +     if !model.animEnabled {                       | |
|                                  |  | +       t.Errorf("expected anim on by default")     | |
|                                  |  +----------------------------------------------------+ |
+--------------------------------------------------------------------------------------------+
| [PANE: LOGS]  j/k: Navigasi | Tab: Panel | /: Cari | y: Salin SHA | a: Animasi | q: Keluar |
+--------------------------------------------------------------------------------------------+
```

---

## Features

- **Terminal-Native Tiling Layout:** Clean 3-panel display inspired by modern tiling window managers, showing your commit timeline, metadata inspector, and syntax-highlighted code diff simultaneously.
- **Cyber Breathing Glow Animation:** Active panel border pulses with a smooth neon gradient cycle (Cyan to Indigo to Violet) at 10 FPS, zero CPU overhead.
- **Live Telemetry Wave:** Top header displays a real-time block-character pulse (`▂▄▆█▇▅▃`) to signal active repository monitoring.
- **Deep Cosmic Cyber Theme:** Crafted with 24-bit TrueColor ANSI palettes featuring deep slate backgrounds, vibrant cyan accents, and clear status pill badges.
- **Interactive Search Filter:** Instant commit searching by message, author, or hash with `/` shortcut.
- **Clipboard SHA Yanking:** Copy full commit hashes directly to system clipboard using `y`.
- **Pure Go & Zero CGo:** Runs smoothly on Windows Terminal, PowerShell, and Unix shells with zero native library headaches.
- **Vim-Style Keybindings:** Navigate effortlessly using standard keyboard shortcuts (`j`, `k`, `Tab`, `Enter`, `/`, `y`, `a`, `q`).
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
| `a` | Toggle animation on / off (battery saver mode) |
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
