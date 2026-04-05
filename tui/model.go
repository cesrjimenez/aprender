package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Mode represents the current app mode
type mode int

const (
	modeSplash mode = iota
	modeList
	modeAdd
)

// Logo ASCII art
const logo = `
 ████████╗ ██████╗ ██████╗  ██████╗     ████████╗██╗   ██╗██╗
 ╚══██╔══╝██╔═══██╗██╔══██╗██╔═══██╗    ╚══██╔══╝██║   ██║██║
    ██║   ██║   ██║██║  ██║██║   ██║       ██║   ██║   ██║██║
    ██║   ██║   ██║██║  ██║██║   ██║       ██║   ██║   ██║██║
    ██║   ╚██████╔╝██████╔╝╚██████╔╝       ██║   ╚██████╔╝██║
    ╚═╝    ╚═════╝ ╚═════╝  ╚═════╝        ╚═╝    ╚═════╝ ╚═╝
`

// model represents the state of the TUI
type model struct {
	service   *TodoService
	todos     []*Todo
	table     table.Model
	mode      mode
	input     textinput.Model
	err       error
	statusMsg string
}

// splashTimeoutMsg is sent when the splash screen should end
type splashTimeoutMsg struct{}

// splashTimeout returns a command that sends a message after a delay
func splashTimeout() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return splashTimeoutMsg{}
	})
}

// initialModel creates the initial model
func initialModel(service *TodoService) model {
	ti := textinput.New()
	ti.Placeholder = "Enter todo title..."
	ti.CharLimit = 100
	ti.Width = 50

	todos, _ := service.ListTodos()

	// Setup table
	columns := []table.Column{
		{Title: "Status", Width: 8},
		{Title: "Task", Width: 50},
		{Title: "Created", Width: 20},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	// Style the table
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#7D56F4")).
		Bold(false)
	t.SetStyles(s)

	m := model{
		service: service,
		todos:   todos,
		table:   t,
		mode:    modeSplash,
		input:   ti,
	}

	m.updateTableRows()
	return m
}

// updateTableRows updates the table with current todos
func (m *model) updateTableRows() {
	var rows []table.Row

	for _, todo := range m.todos {
		status := "☐"
		taskStyle := ""
		if todo.Completed {
			status = "☑"
			taskStyle = "✓ "
		}

		created := todo.CreatedAt.Format("Jan 02, 15:04")

		rows = append(rows, table.Row{
			status,
			taskStyle + todo.Title,
			created,
		})
	}

	m.table.SetRows(rows)
}

// Init initializes the model
func (m model) Init() tea.Cmd {
	return splashTimeout()
}

// Update handles messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case splashTimeoutMsg:
		// Transition from splash to list mode
		m.mode = modeList
		return m, nil

	case tea.KeyMsg:
		switch m.mode {
		case modeSplash:
			// Any key press skips the splash screen
			m.mode = modeList
			return m, nil
		case modeList:
			return m.updateListMode(msg)
		case modeAdd:
			return m.updateAddMode(msg)
		}
	}

	return m, nil
}

// updateListMode handles key presses in list mode
func (m *model) updateListMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit

	case "a", "n":
		// Switch to add mode
		m.mode = modeAdd
		m.input.Focus()
		m.input.SetValue("")
		m.statusMsg = ""
		return m, textinput.Blink

	case "enter", " ":
		// Toggle todo completion
		if len(m.todos) > 0 {
			selectedRow := m.table.Cursor()
			if selectedRow < len(m.todos) {
				todo := m.todos[selectedRow]
				_, err := m.service.ToggleTodo(todo.ID)
				if err != nil {
					m.statusMsg = fmt.Sprintf("Error: %v", err)
				} else {
					m.statusMsg = "Todo updated!"
				}
				// Refresh todos and table
				m.todos, _ = m.service.ListTodos()
				m.updateTableRows()
			}
		}

	case "d", "x":
		// Delete todo
		if len(m.todos) > 0 {
			selectedRow := m.table.Cursor()
			if selectedRow < len(m.todos) {
				todo := m.todos[selectedRow]
				err := m.service.DeleteTodo(todo.ID)
				if err != nil {
					m.statusMsg = fmt.Sprintf("Error: %v", err)
				} else {
					m.statusMsg = "Todo deleted!"
				}
				// Refresh todos and table
				m.todos, _ = m.service.ListTodos()
				m.updateTableRows()
			}
		}

	default:
		// Pass other keys to the table (for navigation)
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}

	return m, nil
}

// updateAddMode handles key presses in add mode
func (m *model) updateAddMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit

	case "esc":
		// Cancel and return to list mode
		m.mode = modeList
		m.input.Blur()
		m.statusMsg = ""
		return m, nil

	case "enter":
		// Create the todo
		title := strings.TrimSpace(m.input.Value())
		if title != "" {
			_, err := m.service.CreateTodo(title)
			if err != nil {
				m.statusMsg = fmt.Sprintf("Error: %v", err)
			} else {
				m.statusMsg = "Todo created!"
				m.todos, _ = m.service.ListTodos()
				m.updateTableRows()
			}
		}
		m.mode = modeList
		m.input.Blur()
		return m, nil
	}

	// Handle input updates
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View renders the UI
func (m model) View() string {
	var s strings.Builder

	// Render different views based on mode
	switch m.mode {
	case modeSplash:
		return m.renderSplashView()
	case modeList:
		// Header
		headerStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			Padding(1, 0)

		s.WriteString(headerStyle.Render("📝 TODO LIST"))
		s.WriteString("\n\n")
		s.WriteString(m.renderListView())
	case modeAdd:
		// Header
		headerStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			Padding(1, 0)

		s.WriteString(headerStyle.Render("📝 TODO LIST"))
		s.WriteString("\n\n")
		s.WriteString(m.renderAddView())
	}

	return s.String()
}

// renderSplashView renders the splash screen with logo
func (m *model) renderSplashView() string {
	var s strings.Builder

	// Style the logo
	logoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7D56F4")).
		Bold(true).
		Align(lipgloss.Center)

	// Add some padding at the top
	s.WriteString("\n\n")
	s.WriteString(logoStyle.Render(logo))
	s.WriteString("\n\n")

	// Tagline
	taglineStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Italic(true).
		Align(lipgloss.Center)

	tagline := "Your tasks, beautifully organized in the terminal"
	s.WriteString(taglineStyle.Render(tagline))
	s.WriteString("\n\n")

	// Loading indicator
	loadingStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#04B575")).
		Align(lipgloss.Center)

	loading := "Press any key to continue..."
	s.WriteString(loadingStyle.Render(loading))

	return s.String()
}

// renderListView renders the list of todos
func (m *model) renderListView() string {
	var s strings.Builder

	// Stats
	total, completed, pending, _ := m.service.GetStats()
	statsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Padding(0, 0, 1, 0)

	stats := fmt.Sprintf("Total: %d | Completed: %d | Pending: %d", total, completed, pending)
	s.WriteString(statsStyle.Render(stats))
	s.WriteString("\n")

	// Table
	if len(m.todos) == 0 {
		emptyStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true)
		s.WriteString(emptyStyle.Render("No todos yet. Press 'a' to add one!"))
		s.WriteString("\n")
	} else {
		s.WriteString(m.table.View())
		s.WriteString("\n")
	}

	// Status message
	if m.statusMsg != "" {
		s.WriteString("\n")
		statusStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Italic(true)
		s.WriteString(statusStyle.Render(m.statusMsg))
		s.WriteString("\n")
	}

	// Help
	s.WriteString("\n")
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Padding(1, 0, 0, 0)

	help := "↑/↓ navigate • enter/space toggle • a add • d delete • q quit"
	s.WriteString(helpStyle.Render(help))

	return s.String()
}

// renderAddView renders the add todo view
func (m *model) renderAddView() string {
	var s strings.Builder

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7D56F4")).
		Bold(true)

	s.WriteString(labelStyle.Render("Add New Todo"))
	s.WriteString("\n\n")
	s.WriteString(m.input.View())
	s.WriteString("\n\n")

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262"))

	help := "enter to save • esc to cancel"
	s.WriteString(helpStyle.Render(help))

	return s.String()
}
