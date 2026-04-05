package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize the store
	store := NewMemoryStore()

	// Initialize the service
	service := NewTodoService(store)

	// Create the model
	m := initialModel(service)

	// Start the program with alternate screen (fullscreen mode)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
