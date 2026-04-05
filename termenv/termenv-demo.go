package termenv

import (
	"fmt"

	"github.com/muesli/termenv"
)

func mainFour() {
	// Termenv automatically detects your terminal's color capabilities
	// and adapts the output accordingly

	output := termenv.NewOutput(nil)

	fmt.Println("=== Termenv Demo: Adaptive Terminal Styling ===\n")

	// 1. Basic Colors
	fmt.Println("1. Basic Colors:")
	fmt.Println(output.String("   Red text").Foreground(termenv.ANSIRed))
	fmt.Println(output.String("   Green text").Foreground(termenv.ANSIGreen))
	fmt.Println(output.String("   Blue text").Foreground(termenv.ANSIBlue))
	fmt.Println(output.String("   Yellow background").Background(termenv.ANSIYellow))
	fmt.Println()

	// 2. Bold, Italic, Underline
	fmt.Println("2. Text Styles:")
	fmt.Println(output.String("   Bold text").Bold())
	fmt.Println(output.String("   Italic text").Italic())
	fmt.Println(output.String("   Underlined text").Underline())
	fmt.Println(output.String("   Bold + Red").Bold().Foreground(termenv.ANSIRed))
	fmt.Println()

	// 3. True Color (RGB) - adapts based on terminal support
	fmt.Println("3. True Color (RGB) - adapts to your terminal:")
	pink := output.Color("#FF1493")
	purple := output.Color("#9370DB")
	cyan := output.Color("#00CED1")

	fmt.Println(output.String("   Hot Pink").Foreground(pink))
	fmt.Println(output.String("   Medium Purple").Foreground(purple))
	fmt.Println(output.String("   Dark Cyan").Foreground(cyan))
	fmt.Println()

	// 4. The Magic: Terminal Detection
	fmt.Println("4. Your Terminal Info:")
	fmt.Printf("   Color Profile: %s\n", output.Profile)
	fmt.Printf("   Supports True Color: %v\n", output.Profile == termenv.TrueColor)
	fmt.Printf("   Supports 256 Colors: %v\n", output.Profile >= termenv.ANSI256)
	fmt.Println()

	// 5. Gradients
	fmt.Println("5. Color Gradients:")

	text := "    This is a gradient from red to blue!"
	for i, char := range text {
		// Calculate color between start and end
		ratio := float64(i) / float64(len(text))
		color := interpolateColor(output, ratio)
		fmt.Print(output.String(string(char)).Foreground(color))
	}
	fmt.Println("\n")

	// 6. Background + Foreground combinations
	fmt.Println("6. Styled Boxes:")
	successBg := output.Color("#10B981")
	errorBg := output.Color("#EF4444")
	warningBg := output.Color("#F59E0B")

	fmt.Println(output.String("  SUCCESS  ").
		Background(successBg).
		Foreground(termenv.ANSIWhite).
		Bold())

	fmt.Println(output.String("  ERROR    ").
		Background(errorBg).
		Foreground(termenv.ANSIWhite).
		Bold())

	fmt.Println(output.String("  WARNING  ").
		Background(warningBg).
		Foreground(termenv.ANSIBlack).
		Bold())

	fmt.Println()
}

// Helper function to interpolate between two colors
func interpolateColor(output *termenv.Output, ratio float64) termenv.Color {
	// Simple gradient from red to blue
	// In production, you'd want proper color space interpolation
	if ratio < 0.33 {
		return output.Color("#FF0000") // Red
	} else if ratio < 0.66 {
		return output.Color("#8000FF") // Purple
	}
	return output.Color("#0000FF") // Blue
}
