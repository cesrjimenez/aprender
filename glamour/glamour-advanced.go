package glamour

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
)

func mainSeven() {
	// Check if a style argument is provided
	style := "auto"
	if len(os.Args) > 1 {
		style = os.Args[1]
	}

	// Read the README file
	content, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Printf("Error reading README: %v\n", err)
		os.Exit(1)
	}

	// Create renderer based on the style argument
	var r *glamour.TermRenderer

	switch style {
	case "dark":
		r, err = glamour.NewTermRenderer(
			glamour.WithStylePath("dark"),
			glamour.WithWordWrap(100),
		)
	case "light":
		r, err = glamour.NewTermRenderer(
			glamour.WithStylePath("light"),
			glamour.WithWordWrap(100),
		)
	case "pink":
		r, err = glamour.NewTermRenderer(
			glamour.WithStylePath("pink"),
			glamour.WithWordWrap(100),
		)
	case "notty":
		r, err = glamour.NewTermRenderer(
			glamour.WithStylePath("notty"),
			glamour.WithWordWrap(100),
		)
	default:
		// Auto-detects terminal capabilities
		r, err = glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(100),
		)
	}

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

// Usage examples:
// go run glamour-advanced.go           # Auto-detect style
// go run glamour-advanced.go dark      # Dark theme
// go run glamour-advanced.go light     # Light theme
// go run glamour-advanced.go pink      # Pink theme
// go run glamour-advanced.go notty     # Plain text (no colors)
