package termenv

import (
	"fmt"

	"github.com/muesli/termenv"
)

// This shows the problem Termenv solves

func maingo() {
	fmt.Println("=== The Problem Termenv Solves ===\n")

	// WITHOUT Termenv - Manual ANSI codes (fragile!)
	fmt.Println("❌ WITHOUT Termenv (manual ANSI codes):")
	fmt.Println("   Hard-coded escape sequences:")
	fmt.Printf("   \033[31mRed text\033[0m\n") // What if terminal doesn't support this?
	fmt.Printf("   \033[1;34mBold Blue\033[0m\n")
	fmt.Printf("   \033[48;5;200mPink background\033[0m\n")
	fmt.Println()
	fmt.Println("   Problems:")
	fmt.Println("   • Might not work on all terminals")
	fmt.Println("   • Ugly, hard to read code")
	fmt.Println("   • You have to remember ANSI codes")
	fmt.Println("   • No automatic adaptation")
	fmt.Println()

	// WITH Termenv - Clean and adaptive!
	output := termenv.NewOutput(nil)

	fmt.Println("✅ WITH Termenv (clean and adaptive):")
	fmt.Println("   Same output, but properly adapted:")
	fmt.Println("  ", output.String("Red text").Foreground(termenv.ANSIRed))
	fmt.Println("  ", output.String("Bold Blue").Bold().Foreground(termenv.ANSIBlue))
	fmt.Println("  ", output.String("Pink background").Background(output.Color("#FF69B4")))
	fmt.Println()
	fmt.Println("   Benefits:")
	fmt.Println("   • Works on ALL terminals")
	fmt.Println("   • Clean, readable code")
	fmt.Println("   • Use hex colors (#FF69B4)")
	fmt.Println("   • Automatic color conversion")
	fmt.Println()

	// Show code comparison
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("CODE COMPARISON:")
	fmt.Println()
	fmt.Println("Without Termenv:")
	fmt.Println(`  fmt.Printf("\033[31mError\033[0m\n")  // Ugly!`)
	fmt.Println()
	fmt.Println("With Termenv:")
	fmt.Println(`  output.String("Error").Foreground(termenv.ANSIRed)  // Clean!`)
	fmt.Println()

	// Real world example
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("REAL WORLD EXAMPLE:")
	fmt.Println()

	// Building a logger
	errorColor := output.Color("#EF4444")
	successColor := output.Color("#10B981")
	warningColor := output.Color("#F59E0B")
	infoColor := output.Color("#3B82F6")

	fmt.Println(output.String(" ERROR ").
		Background(errorColor).
		Foreground(termenv.ANSIWhite).
		Bold(), "Failed to connect to database")

	fmt.Println(output.String(" SUCCESS ").
		Background(successColor).
		Foreground(termenv.ANSIWhite).
		Bold(), "Deployment completed")

	fmt.Println(output.String(" WARNING ").
		Background(warningColor).
		Foreground(termenv.ANSIBlack).
		Bold(), "API rate limit approaching")

	fmt.Println(output.String(" INFO ").
		Background(infoColor).
		Foreground(termenv.ANSIWhite).
		Bold(), "Starting server on port 3000")

	fmt.Println()
	fmt.Println("These colors work EVERYWHERE because Termenv")
	fmt.Println("automatically converts them to what your terminal supports!")
}
