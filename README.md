# gitgwtuh

A high-aesthetic, keyboard-driven Git TUI explorer and interactive diff viewer built with Go.

---

## Features

- **Terminal-Native Tiling Layout:** Clean 3-panel display inspired by modern tiling window managers, showing your commit timeline, metadata inspector, and syntax-highlighted code diff simultaneously.
- **Deep Cosmic Cyber Theme:** Crafted with 24-bit TrueColor ANSI palettes featuring deep slate backgrounds, vibrant cyan accents, and clear status pill badges.
- **Pure Go & Zero CGo:** Runs smoothly on Windows Terminal, PowerShell, and Unix shells with zero native library headaches.
- **Vim-Style Keybindings:** Navigate effortlessly using standard keyboard shortcuts (`j`, `k`, `Tab`, `Enter`, `/`, `q`).
- **Dynamic Resize Engine:** Automatically recalculates panel bounds on window resize events to keep borders and text crisp.
- **Safe Read-Only Operation:** Inspects commits, branches, and diffs without modifying or risking your repository state.

---

## Keybindings

| Key | Action |
|-----|--------|
| `j` / `Down` | Move down to next commit / item |
| `k` / `Up` | Move up to previous commit / item |
| `Tab` | Switch active panel focus |
| `Enter` | Expand commit details or view file diff |
| `/` | Filter and search commits by message or author |
| `q` / `Esc` | Quit and restore terminal screen cleanly |

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

---

## License

MIT License (c) 2026 hihihehadika
