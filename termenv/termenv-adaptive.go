package termenv

import (
	"fmt"

	"github.com/muesli/termenv"
)

// This demonstrates Termenv's KILLER FEATURE:
// Automatic color downgrading based on terminal capabilities

func mainOne() {
	output := termenv.NewOutput(nil)

	fmt.Println("=== Termenv's Magic: Adaptive Color Conversion ===\n")

	// You specify a true color (24-bit RGB)
	// Termenv automatically converts it based on what your terminal supports

	myColor := output.Color("#FF6B9D") // A nice pink color
	_ = myColor                        // We won't use it directly here

	fmt.Println("Your desired color: #FF6B9D (a pretty pink)")
	fmt.Println()

	// Show what happens in different terminal types
	fmt.Println("How this color appears in different terminals:\n")

	// 1. True Color terminal (modern terminals)
	trueColorOutput := termenv.NewOutput(nil)
	trueColorOutput.Profile = termenv.TrueColor
	fmt.Printf("TrueColor Terminal (24-bit):  %s\n",
		trueColorOutput.String("████████").
			Foreground(trueColorOutput.Color("#FF6B9D")))
	fmt.Println("   → Uses exact RGB: #FF6B9D")
	fmt.Println()

	// 2. 256-color terminal
	ansi256Output := termenv.NewOutput(nil)
	ansi256Output.Profile = termenv.ANSI256
	fmt.Printf("ANSI256 Terminal (8-bit):     %s\n",
		ansi256Output.String("████████").
			Foreground(ansi256Output.Color("#FF6B9D")))
	fmt.Println("   → Converts to closest 256-color match")
	fmt.Println()

	// 3. 16-color terminal (old school)
	ansiOutput := termenv.NewOutput(nil)
	ansiOutput.Profile = termenv.ANSI
	fmt.Printf("ANSI Terminal (4-bit):        %s\n",
		ansiOutput.String("████████").
			Foreground(ansiOutput.Color("#FF6B9D")))
	fmt.Println("   → Converts to closest 16-color match")
	fmt.Println()

	// 4. No color terminal
	asciiOutput := termenv.NewOutput(nil)
	asciiOutput.Profile = termenv.Ascii
	fmt.Printf("ASCII Terminal (no color):    %s\n",
		asciiOutput.String("████████"))
	fmt.Println("   → No colors at all")
	fmt.Println()

	// The Point: You write ONE line of code
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("THE MAGIC:")
	fmt.Println()
	fmt.Println("  You write: output.Color(\"#FF6B9D\")")
	fmt.Println("  Termenv automatically adapts to:")
	fmt.Println("    ✓ Modern terminals → Full RGB")
	fmt.Println("    ✓ Older terminals  → Closest match")
	fmt.Println("    ✓ Basic terminals  → Basic colors")
	fmt.Println("    ✓ No-color mode    → Plain text")
	fmt.Println()
	fmt.Println("YOUR APP WORKS EVERYWHERE! 🎉")
	fmt.Println()

	// Show actual terminal capabilities
	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Your Terminal Capabilities:")
	fmt.Printf("  Profile: %s\n", output.Profile)

	switch output.Profile {
	case termenv.TrueColor:
		fmt.Println("  🎨 Your terminal supports 16 million colors!")
	case termenv.ANSI256:
		fmt.Println("  🎨 Your terminal supports 256 colors")
	case termenv.ANSI:
		fmt.Println("  🎨 Your terminal supports 16 colors")
	case termenv.Ascii:
		fmt.Println("  📝 Your terminal has no color support")
	}

	fmt.Println()
	fmt.Println("Termenv detected this automatically and adapts")
	fmt.Println("all colors you use to match these capabilities!")
}
