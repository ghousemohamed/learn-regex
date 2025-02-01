# Learn Regex in the Terminal

![Learn Regex Demo](media/demo.gif)

An interactive command-line application to learn and practice regular expressions (regex) right in your terminal. Built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Features

- 🎓 Interactive tutorial with progressive lessons
- 💪 Practice problems to test your skills
- 💾 Progress tracking across sessions
- 🎨 Beautiful terminal UI with gradient text and modern design

## Installation

```
wget https://raw.githubusercontent.com/ghousemohamed/learn-regex/main/install.sh -O install.sh && bash install.sh
```

Or clone and build from source:

```
git clone https://github.com/ghousemohamed/learn-regex-in-the-terminal.git
cd learn-regex-in-the-terminal
go build
```

## Usage

Simply run the application in your terminal:

```
learn-regex
```

### Controls

- `↑`/`↓` or `j`/`k`: Navigate menu options
- `Enter`: Submit regex pattern / Select menu option
- `Tab`: Skip to next lesson/problem
- `Shift + Tab`: Go to previous lesson/problem
- `Ctrl + r`: Reset progress
- `Esc`: Return to main menu
- `Ctrl + c`: Save progress and quit

## Learning Path

The tutorial is structured to take you from regex basics to advanced patterns:

1. Start with basic character matching
2. Progress through special characters and metacharacters
3. Learn about quantifiers and groups
4. Practice with real-world examples
5. Test your skills with challenging problems

Your progress is automatically saved, allowing you to continue where you left off.

## Development

This project is built using:
- [Go](https://golang.org/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
