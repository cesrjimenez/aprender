package termenv

import (
	"fmt"
	"time"

	"github.com/muesli/termenv"
)

// This example shows WHY Termenv is useful:
// Building a status dashboard that looks good on ANY terminal

func mainThree() {
	output := termenv.NewOutput(nil)

	// Clear screen
	fmt.Print("\033[2J\033[H")

	// Title
	title := output.String(" 🚀 DEPLOYMENT DASHBOARD ").
		Background(output.Color("#7D56F4")).
		Foreground(termenv.ANSIWhite).
		Bold()
	fmt.Println(title)
	fmt.Println()

	// Services status
	renderService(output, "API Server", "running", 99.9)
	renderService(output, "Database", "running", 100.0)
	renderService(output, "Cache", "degraded", 85.2)
	renderService(output, "Queue", "running", 98.1)
	renderService(output, "CDN", "error", 0.0)

	fmt.Println()

	// Metrics
	renderMetric(output, "CPU Usage", 45.2, 100)
	renderMetric(output, "Memory", 67.8, 100)
	renderMetric(output, "Disk", 23.1, 100)

	fmt.Println()

	// Footer
	timestamp := time.Now().Format("15:04:05")
	footer := output.String(fmt.Sprintf(" Last updated: %s ", timestamp)).
		Foreground(output.Color("#626262"))
	fmt.Println(footer)
}

func renderService(output *termenv.Output, name, status string, uptime float64) {
	// Status indicator
	var statusText string
	var statusColor termenv.Color

	switch status {
	case "running":
		statusText = "●"
		statusColor = output.Color("#10B981") // Green
	case "degraded":
		statusText = "●"
		statusColor = output.Color("#F59E0B") // Orange
	case "error":
		statusText = "●"
		statusColor = output.Color("#EF4444") // Red
	default:
		statusText = "○"
		statusColor = output.Color("#626262") // Gray
	}

	indicator := output.String(statusText).Foreground(statusColor)
	serviceName := output.String(fmt.Sprintf("%-15s", name)).Bold()
	statusDisplay := output.String(fmt.Sprintf("%-10s", status)).Foreground(statusColor)
	uptimeDisplay := output.String(fmt.Sprintf("%.1f%%", uptime))

	fmt.Printf(" %s %s %s uptime: %s\n",
		indicator,
		serviceName,
		statusDisplay,
		uptimeDisplay,
	)
}

func renderMetric(output *termenv.Output, name string, value, max float64) {
	percentage := (value / max) * 100

	// Choose color based on value
	var barColor termenv.Color
	if percentage < 50 {
		barColor = output.Color("#10B981") // Green
	} else if percentage < 80 {
		barColor = output.Color("#F59E0B") // Orange
	} else {
		barColor = output.Color("#EF4444") // Red
	}

	// Build progress bar
	barWidth := 30
	filled := int(float64(barWidth) * (value / max))

	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	styledBar := output.String(bar).Foreground(barColor)
	metricName := output.String(fmt.Sprintf("%-12s", name)).Bold()
	valueDisplay := output.String(fmt.Sprintf("%5.1f%%", percentage))

	fmt.Printf(" %s [%s] %s\n", metricName, styledBar, valueDisplay)
}
