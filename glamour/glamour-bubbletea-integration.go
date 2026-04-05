package glamour

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"os"
)

// This example shows how to integrate Glamour into a Bubbletea app
// Press 'h' to show help documentation rendered with Glamour

const helpDoc = `# TODO TUI Help

## Getting Started

Welcome to **TODO TUI**! This application helps you manage your tasks beautifully in the terminal.

## Keyboard Shortcuts

### Navigation
- **↑/↓** - Move up and down through your todos
- **j/k** - Vim-style navigation (same as ↑/↓)

### Actions
- **Space/Enter** - Toggle todo completion status
- **a** - Add a new todo
- **d** - Delete the selected todo
- **h** - Show this help screen
- **q** - Quit the application

## Features

1. **Beautiful Table View** - See all your todos in a clean table
2. **Real-time Stats** - Track completed vs pending tasks
3. **Persistent Storage** - Your todos are saved (coming soon!)
4. **Keyboard-First** - No mouse needed!

## Tips & Tricks

> Use the Vim keys (**j/k**) for lightning-fast navigation!

Press **ESC** to close this help screen.
`

type helpModel struct {
	viewport viewport.Model
	ready    bool
	content  string
}

func initialHelpModel() helpModel {
	// Render the help markdown with Glamour
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(70),
	)

	rendered, _ := r.Render(helpDoc)

	return helpModel{
		content: rendered,
	}
}

func (m helpModel) Init() tea.Cmd {
	return nil
}

func (m helpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := 3
		footerHeight := 2
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m helpModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Padding(1, 2)

	header := headerStyle.Render("📚 Help Documentation")

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Padding(1, 2)

	footer := footerStyle.Render("↑/↓ scroll • q/esc quit")

	return fmt.Sprintf("%s\n%s\n%s", header, m.viewport.View(), footer)
}

func mainNine() {
	p := tea.NewProgram(
		initialHelpModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

// This demonstrates how you could add a help mode to your TODO app:
// 1. Add a 'modeHelp' to your mode enum
// 2. When user presses 'h', switch to help mode and render help with Glamour
// 3. When user presses 'esc', return to list mode
