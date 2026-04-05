# Glamour Demo Guide

## What is Glamour?

Glamour renders markdown beautifully in the terminal with:
- Syntax highlighting for code blocks
- Beautiful headers and formatting
- Automatic terminal color adaptation
- Support for tables, lists, and more

## Files Created

### 1. `glamour-demo.go` - Simple README Renderer
The simplest example - just renders your README.md with auto-detected styling.

**Run it:**
```bash
go run glamour-demo.go
```

**What it shows:**
- Basic Glamour usage
- Auto-style detection (adapts to your terminal)
- Word wrapping

### 2. `glamour-advanced.go` - Style Switcher
Shows different Glamour themes and how to switch between them.

**Run it:**
```bash
# Auto-detect (default)
go run glamour-advanced.go

# Try different themes
go run glamour-advanced.go dark
go run glamour-advanced.go light
go run glamour-advanced.go pink
go run glamour-advanced.go notty  # Plain text, no colors
```

**What it shows:**
- Multiple theme options
- Command-line arguments
- Style customization

### 3. `glamour-bubbletea-integration.go` - Integrated Help System
Shows how to integrate Glamour INTO a Bubbletea app with scrollable help docs.

**Run it:**
```bash
go run glamour-bubbletea-integration.go
```

**What it shows:**
- Glamour + Bubbletea working together
- Scrollable markdown content with viewport
- Real-world use case (help documentation)
- How you could add this to your TODO app

## For Your Video

### Demo Flow

**1. Start Simple (30 seconds)**
```bash
go run glamour-demo.go
```
Show: "This is your plain markdown README..."

**2. Show the Magic (30 seconds)**
"Now watch what Glamour does to it!"
Show the beautifully rendered output with:
- Colored headers
- Formatted code blocks
- Beautiful tables
- Styled lists

**3. Show Themes (30 seconds)**
```bash
go run glamour-advanced.go dark
go run glamour-advanced.go pink
```
"It even has multiple themes!"

**4. Show Integration (1 minute)**
```bash
go run glamour-bubbletea-integration.go
```
"And you can integrate it into your Bubbletea apps for beautiful help docs!"

### Key Points to Mention

✅ **"Every CLI tool needs documentation"**
✅ **"Glamour makes your markdown look professional"**
✅ **"Automatically adapts to terminal capabilities"**
✅ **"Works perfectly with Bubbletea"**

### Talking Points

- "You've probably seen this in tools like Glow, Soft Serve, and others"
- "It's the same library that powers many modern CLI tools"
- "No more ugly plain text help docs"
- "Your users can read your README right in the terminal"

## Quick Setup

```bash
# Install dependencies
go get github.com/charmbracelet/glamour

# Run any demo
go run glamour-demo.go
go run glamour-advanced.go
go run glamour-bubbletea-integration.go
```

## Integration with Your TODO App

You could add a help mode to your TODO app:
1. Press `h` to show help
2. Renders beautiful markdown docs
3. Scrollable with viewport
4. Press `esc` to go back

See `glamour-bubbletea-integration.go` for the full implementation!