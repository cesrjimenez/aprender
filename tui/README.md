# 📝 TODO List TUI

> A beautiful terminal-based TODO list application built with Go and the [Charm](https://charm.sh) ecosystem.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## ✨ Features

- ✅ **Add, toggle, and delete todos** with ease
- 📊 **Real-time statistics** showing total, completed, and pending tasks
- ⌨️  **Vim-style keyboard navigation** for power users
- 🎨 **Beautiful terminal UI** with colors and styling
- 🏗️  **Clean architecture** following Service + Store pattern
- 🖥️  **Fullscreen mode** for distraction-free task management

## 🚀 Quick Start

```bash
# Install dependencies
go mod download

# Build the application
go build -o todo-tui

# Run it!
./todo-tui
```

## 🎮 Keyboard Shortcuts

### List Mode
| Key | Action |
|-----|--------|
| `↑` / `↓` | Navigate up/down |
| `Enter` / `Space` | Toggle todo completion |
| `a` / `n` | Add new todo |
| `d` / `x` | Delete selected todo |
| `q` | Quit application |

### Add Mode
| Key | Action |
|-----|--------|
| `Enter` | Save todo |
| `Esc` | Cancel and return |

## 🏗️ Architecture

This project follows the **Service + Store** pattern for clean separation of concerns:

```
┌─────────────────────────────────────────┐
│           Bubbletea UI (TUI)            │
│         (model.go - View Layer)         │
└────────────────┬────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────┐
│         Service Layer (Logic)           │
│   (service.go - Business Rules)         │
└────────────────┬────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────┐
│        Store Layer (Data)               │
│   (store.go - In-Memory Storage)        │
└─────────────────────────────────────────┘
```

### Why This Pattern?

1. **Testability** - Mock the store for isolated unit tests
2. **Flexibility** - Swap in-memory storage for SQLite, PostgreSQL, etc.
3. **Maintainability** - Business logic separate from UI and storage
4. **Scalability** - Clear separation makes the codebase easy to grow

## 📦 Project Structure

```
.
├── main.go      # Application entry point
├── store.go     # Data storage layer (in-memory)
├── service.go   # Business logic layer
├── model.go     # Bubbletea UI implementation
└── go.mod       # Dependencies
```

## 🛠️ Built With

- **[Bubbletea](https://github.com/charmbracelet/bubbletea)** - The Elm Architecture for Go TUIs
- **[Bubbles](https://github.com/charmbracelet/bubbles)** - TUI components (table, textinput)
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Style definitions for terminal rendering

## 🎯 Next Steps

- [ ] Add SQLite persistence
- [ ] Implement todo categories/tags
- [ ] Add due dates and reminders
- [ ] Add search/filter functionality
- [ ] Support custom color themes
- [ ] Add data import/export

## 🤝 Contributing

Contributions are welcome! Feel free to:

- 🐛 Report bugs
- 💡 Suggest new features
- 🔧 Submit pull requests

## 📄 License

MIT License - feel free to use this in your own projects!

## 🙏 Acknowledgments

Built with love using the amazing [Charm](https://charm.sh) ecosystem. Special thanks to the Charm team for creating such powerful and delightful terminal tools.

---

**Happy task managing! 🚀**