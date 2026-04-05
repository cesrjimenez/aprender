# Termenv Demo Guide

## What IS Termenv?

**Termenv is a color and styling library that automatically adapts to your terminal's capabilities.**

Think of it like this:
- You want to use beautiful colors in your CLI app
- Different terminals support different color modes (16 colors, 256 colors, true color)
- Termenv detects what YOUR terminal supports and converts colors automatically
- Your app works perfectly on EVERY terminal!

## The Problem It Solves

**Without Termenv:**
```go
// Hard-coded ANSI escape sequences
fmt.Printf("\033[31mError\033[0m")  // Might not work everywhere!
```

**With Termenv:**
```go
output := termenv.NewOutput(nil)
output.String("Error").Foreground(termenv.ANSIRed)  // Works everywhere!
```

## Files Created

### 1. `termenv-demo.go` - Feature Overview
Shows all the main features of Termenv.

**Run it:**
```bash
go run termenv-demo.go
```

**What it demonstrates:**
- Basic colors (red, green, blue)
- Text styles (bold, italic, underline)
- True color support (RGB hex colors like #FF1493)
- Terminal capability detection
- Gradients
- Background + foreground combinations

### 2. `termenv-dashboard.go` - Practical Example
A real-world status dashboard showing why Termenv is useful.

**Run it:**
```bash
go run termenv-dashboard.go
```

**What it demonstrates:**
- Service status indicators (colored dots)
- Progress bars with dynamic colors
- Professional-looking dashboard
- Everything adapts to your terminal automatically

### 3. `termenv-adaptive.go` - The Magic Explained
Shows how Termenv adapts colors to different terminal types.

**Run it:**
```bash
go run termenv-adaptive.go
```

**What it demonstrates:**
- How one color (#FF6B9D) appears in different terminals
- True color vs 256-color vs 16-color vs no-color
- Terminal capability detection
- Why this matters for compatibility

### 4. `termenv-comparison.go` - Before & After
Direct comparison of manual ANSI codes vs Termenv.

**Run it:**
```bash
go run termenv-comparison.go
```

**What it demonstrates:**
- Ugly ANSI codes vs clean Termenv code
- Real-world logger example
- Code readability improvement

## For Your Video (2-3 minutes)

### Demo Flow

**1. Start with the Problem (30 seconds)**
```bash
go run termenv-comparison.go
```
Show the ugly ANSI codes vs clean Termenv syntax

**2. Show the Features (45 seconds)**
```bash
go run termenv-demo.go
```
"Look at all these styling options - colors, bold, italic, RGB colors..."

**3. Show the Magic (45 seconds)**
```bash
go run termenv-adaptive.go
```
"Here's the killer feature - automatic color adaptation!"
Show how one color looks different in different terminal modes

**4. Real World Example (60 seconds)**
```bash
go run termenv-dashboard.go
```
"Here's a real dashboard - notice the colored status indicators and progress bars"
"This works perfectly on any terminal because Termenv handles the conversion"

### Key Talking Points

✅ **"The Foundation of Terminal Styling"**
- Lipgloss (which you saw in Bubbletea) is built ON TOP of Termenv
- Termenv is the low-level engine

✅ **"Write Once, Run Anywhere"**
- You write: `output.Color("#FF6B9D")`
- Termenv converts it for:
    - Modern terminals → Full 24-bit RGB
    - Older terminals → Closest 256-color match
    - Ancient terminals → Closest 16-color match
    - No-color mode → Plain text

✅ **"Why This Matters"**
- Your users have different terminals (iTerm, Terminal.app, Windows Terminal, SSH sessions)
- Without Termenv: broken colors on some terminals
- With Termenv: perfect colors everywhere

✅ **"The Professional Choice"**
- Used by: Glamour, Lipgloss, and many other CLI tools
- Industry standard for terminal color handling
- Part of the Charm ecosystem

## Quick Setup

```bash
# Install
go get github.com/muesli/termenv

# Run any demo
go run termenv-demo.go
go run termenv-dashboard.go
go run termenv-adaptive.go
go run termenv-comparison.go
```

## How It Fits With The Other Tools

**The Stack:**
1. **Termenv** (bottom layer) - Color adaptation and styling
2. **Lipgloss** (middle layer) - Style composition (uses Termenv)
3. **Bubbletea** (top layer) - Interactive apps (uses Lipgloss)
4. **Glamour** (parallel) - Markdown rendering (uses Termenv)

**The Flow:**
```
Your App
   ↓
Bubbletea (interactive)
   ↓
Lipgloss (styling)
   ↓
Termenv (color adaptation)
   ↓
Terminal
```

## One-Liner Explanation

**"Termenv makes your terminal colors work perfectly on any terminal by automatically detecting capabilities and converting colors appropriately."**