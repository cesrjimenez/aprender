package glamour

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
)

func mainFive() {
	// Read the README file
	content, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Printf("Error reading README: %v\n", err)
		os.Exit(1)
	}

	// Create a glamour renderer with a dark style
	// Available styles: "dark", "light", "notty", "auto"
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),  // Automatically adapts to terminal
		glamour.WithWordWrap(80), // Wrap text at 80 characters
	)
	if err != nil {
		fmt.Printf("Error creating renderer: %v\n", err)
		os.Exit(1)
	}

	// Render the markdown
	output, err := r.Render(string(content))
	if err != nil {
		fmt.Printf("Error rendering markdown: %v\n", err)
		os.Exit(1)
	}

	// Print the rendered output
	fmt.Print(output)
}
