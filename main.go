package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	// Internal
	"github.com/ghousemohamed/regex-in-the-terminal/data"
	"github.com/ghousemohamed/regex-in-the-terminal/models"
	"github.com/ghousemohamed/regex-in-the-terminal/storage"
)

type model struct {
	lessons         []models.Lesson
	practices       []models.PracticeProblem
	current         int
	practiceIndex   int
	input           textinput.Model
	err             error
	width           int
	height          int
	quitting        bool
	state           models.CompletionState
	selectedOption  models.WelcomeOption
}

var (
	
	docStyle = lipgloss.NewStyle().
		Align(lipgloss.Center)

	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7B2CBF")).
		Padding(0, 1).
		MarginBottom(1)

	lessonStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4A5568")).
		PaddingLeft(1)

	tocStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 2).
		MarginLeft(2)

	completedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981"))

	incompletedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E11D48"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981"))

	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 2).
		MarginBottom(1).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0080"))

	gradientColors = []string{
		"#FF0080",
		"#FF359A",
		"#FF65B5",
		"#FF94D0",
		"#FFC2EB",
	}

	mainContentStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7B2CBF")).
		Padding(1, 2).
		MarginRight(2)
)

var progressFile = filepath.Join(os.Getenv("HOME"), ".regex_tutorial_progress.json")

func evaluateRegex(pattern string, testCases []models.TestCase) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid regex pattern: %v", err)
	}

	for _, tc := range testCases {
		matches := re.MatchString(tc.Text)
		if matches != tc.Expected {
			if tc.Expected {
				return fmt.Errorf("pattern should match '%s' but doesn't", tc.Text)
			}
			return fmt.Errorf("pattern shouldn't match '%s' but does", tc.Text)
		}
	}
	return nil
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter your regex pattern"
	ti.Focus()

	m := model{
		lessons:   data.GetLessons(),
		practices: data.GetPracticeProblems(),
		input:     ti,
		state:     models.Welcome,
	}

	if progress, err := storage.LoadProgress(); err == nil {
		m.current = progress.CurrentLesson
		m.practiceIndex = progress.PracticeIndex

		for _, lessonID := range progress.Completed {
			var id int
			if _, err := fmt.Sscanf(lessonID, "%d", &id); err == nil && id < len(m.lessons) {
				m.lessons[id].Completed = true
			}
		}

		for _, practiceID := range progress.CompletedPractice {
			var id int
			if _, err := fmt.Sscanf(practiceID, "%d", &id); err == nil && id < len(m.practices) {
				m.practices[id].Completed = true
			}
		}
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			storage.SaveProgress(m.current, m.practiceIndex, m.lessons, m.practices)
			return m, tea.Quit
		case "ctrl+r":
			if m.state == models.Practicing {
				storage.ClearSpecificProgress("practice")
				newM := resetModel(m.width, m.height)

				if progress, err := storage.LoadProgress(); err == nil {
					newM.current = progress.CurrentLesson
					for _, lessonID := range progress.Completed {
						var id int
						if _, err := fmt.Sscanf(lessonID, "%d", &id); err == nil && id < len(newM.lessons) {
							newM.lessons[id].Completed = true
						}
					}
				}
				newM.state = models.Practicing
				return newM, nil
			} else if m.state == models.Learning {
				storage.ClearSpecificProgress("learning")
				newM := resetModel(m.width, m.height)

				if progress, err := storage.LoadProgress(); err == nil {
					newM.practiceIndex = progress.PracticeIndex
					for _, practiceID := range progress.CompletedPractice {
						var id int
						if _, err := fmt.Sscanf(practiceID, "%d", &id); err == nil && id < len(newM.practices) {
							newM.practices[id].Completed = true
						}
					}
				}
				newM.state = models.Learning
				return newM, nil
			}
		case "up", "k":
			if m.state == models.Welcome {
				if m.selectedOption > models.StartLearning {
					m.selectedOption--
				} else {
					m.selectedOption = models.Quit
				}
			}
		case "down", "j":
			if m.state == models.Welcome {
				if m.selectedOption < models.Quit {
					m.selectedOption++
				} else {
					m.selectedOption = 0
				}
			}
		case "enter":
			if m.state == models.Success {
				return m, tea.Quit
			}

			if m.state == models.Welcome {
				switch m.selectedOption {
				case models.StartLearning:
					m.state = models.Learning
				case models.Practice:
					m.state = models.Practicing
				case models.Quit:
					m.quitting = true
					return m, tea.Quit
				}
				return m, nil
			}

			if m.state == models.Practicing {
				if err := evaluateRegex(m.input.Value(), m.practices[m.practiceIndex].TestCases); err != nil {
					m.err = err
				} else {
					m.practices[m.practiceIndex].Completed = true
					storage.SaveProgress(m.current, m.practiceIndex, m.lessons, m.practices)
					if m.practiceIndex < len(m.practices)-1 {
						m.practiceIndex++
					}
					m.input.SetValue("")
					m.err = nil
				}
				return m, nil
			}

			if err := evaluateRegex(m.input.Value(), m.lessons[m.current].TestCases); err != nil {
				m.err = err
			} else {
				m.lessons[m.current].Completed = true
				m.err = nil
				storage.SaveProgress(m.current, m.practiceIndex, m.lessons, m.practices)
				if getCompletedLessons(m) == len(m.lessons) {
					m.state = models.Success
				} else {
					if m.current == len(m.lessons) -1 {
						m.current = 0
					} else {
						m.current++
					}
				}
			}
			m.input.SetValue("")
		case "tab":
			if m.state == models.Learning {
				if m.current == len(m.lessons) - 1 {
					m.current = 0
				} else {
					m.current++
				}
			} else if m.state == models.Practicing {
				if m.practiceIndex == len(m.practices) - 1 {
					m.practiceIndex = 0
				} else {
					m.practiceIndex++
				}
			}
			m.input.SetValue("")
			m.err = nil
		case "shift+tab":
			if m.state == models.Learning && m.current > 0 {
				m.current--
				m.input.SetValue("")
				m.err = nil
			} else if m.state == models.Practicing && m.practiceIndex > 0 {
				m.practiceIndex--
				m.input.SetValue("")
				m.err = nil
			}
		case "esc":
			if m.state == models.Learning || m.state == models.Practicing || m.state == models.Success {
				m.state = models.Welcome
				m.input.SetValue("")
				m.err = nil
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	if m.state == models.Learning || m.state == models.Practicing {
		m.input, cmd = m.input.Update(msg)
	}
	return m, cmd
}

func gradientText(text string) string {
	words := strings.Split(text, "")
	coloredText := make([]string, len(words))
	
	for i, char := range words {
		colorIndex := (i * len(gradientColors)) / len(words)
		if colorIndex >= len(gradientColors) {
			colorIndex = len(gradientColors) - 1
		}
		coloredText[i] = lipgloss.NewStyle().
			Foreground(lipgloss.Color(gradientColors[colorIndex])).
			Render(char)
	}
	
	return strings.Join(coloredText, "")
}

func getCompletedLessons(m model) int {
	completedLessons := 0
	for _, lesson := range m.lessons {
		if lesson.Completed {
			completedLessons++
		}
	}
	return completedLessons
}

func (m model) View() string {
	totalWidth := m.width
	if totalWidth == 0 {
		totalWidth = 120 // fallback width
	}

	leftColumnWidth := (totalWidth * 60) / 100
	rightColumnWidth := (totalWidth * 40) / 100

	// Adjust header width
	headerStyle := headerStyle.Copy().Width(totalWidth - 4) // subtract margin
	headerText := "「 Learn Regex in the Terminal 」"
	header := headerStyle.Render(gradientText(headerText))

	if m.quitting {
		return header + "\n" + "Progress saved. Thanks for learning regex!\n"
	}

	if m.state == models.Success {
		successMsg := lipgloss.JoinVertical(lipgloss.Center,
			"🎉 Congratulations! 🎉",
			"",
			"You've mastered all the regex lessons!",
			"",
			successStyle.Render("Final Stats:"),
			fmt.Sprintf("Completed all %d lessons", len(m.lessons)),
			"",
			"You're now ready to tackle real-world regex challenges!",
			"",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")).Render("Press ENTER to exit or ctrl+r to restart or Esc to go to main screen"),
		)

		return lipgloss.JoinVertical(lipgloss.Center,
			header,
			lipgloss.NewStyle().
				Width(totalWidth - 4).
				Margin(2).
				Padding(2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#10B981")).
				Align(lipgloss.Center).
				Render(successMsg),
		)
	}

	if m.state == models.Welcome {
		progress, _ := storage.LoadProgress()
		hasLessonProgress := progress.CurrentLesson > 0 || len(progress.Completed) > 0
		hasPracticeProgress := progress.PracticeIndex > 0 || len(progress.CompletedPractice) > 0

		var welcomeMsg strings.Builder
		welcomeMsg.WriteString(gradientText("Welcome to the Interactive Regex Tutorial!") + "\n\n")

		if hasLessonProgress {
			completedLessons := len(progress.Completed)
			totalLessons := len(m.lessons)
			welcomeMsg.WriteString(fmt.Sprintf("Tutorial Progress: %d out of %d lessons completed\n", completedLessons, totalLessons))
			welcomeMsg.WriteString(fmt.Sprintf("Last tutorial: Lesson %d: %s\n\n", 
				progress.CurrentLesson+1, 
				m.lessons[progress.CurrentLesson].Title))
		}

		if hasPracticeProgress {
			completedPractice := len(progress.CompletedPractice)
			totalPractice := len(m.practices)
			welcomeMsg.WriteString(fmt.Sprintf("Practice Progress: %d out of %d problems completed\n", completedPractice, totalPractice))
			welcomeMsg.WriteString(fmt.Sprintf("Last practice: Problem %d: %s\n\n",
				progress.PracticeIndex+1,
				m.practices[progress.PracticeIndex].Title))
		}

		welcomeOptions := []string{
			"Continue Learning",
			"Practice Problems",
			"Quit",
		}

		hasProgress := len(progress.Completed) > 0 || progress.CurrentLesson > 0
		if !hasProgress {
			welcomeOptions[0] = "Start Learning"
		}

		for i, option := range welcomeOptions {
			cursor := " "
			if models.WelcomeOption(i) == m.selectedOption {
				cursor = ">"
			}
			style := lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))
			if models.WelcomeOption(i) == m.selectedOption {
				style = style.Bold(true).Foreground(lipgloss.Color("#7B2CBF"))
			}
			welcomeMsg.WriteString(fmt.Sprintf("%s %s\n", cursor, style.Render(option)))
		}

		welcomeMsg.WriteString("\nUse ↑/↓ arrows to select and Enter to confirm")

		return lipgloss.JoinVertical(lipgloss.Center,
			header,
			lipgloss.NewStyle().
				Width(totalWidth - 4).
				Margin(1).
				Padding(1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#874BFD")).
				Render(welcomeMsg.String()),
		)
	}

	if m.state == models.Practicing {
		totalWidth := m.width
		if totalWidth == 0 {
			totalWidth = 120 // fallback width
		}

		leftColumnWidth := (totalWidth * 60) / 100
		rightColumnWidth := (totalWidth * 40) / 100

		doc := strings.Builder{}
		doc.WriteString(header + "\n")

		// Left column content
		var mainContent strings.Builder
		currentProblem := m.practices[m.practiceIndex]
		mainContent.WriteString(titleStyle.Render(currentProblem.Title) + "\n\n")
		mainContent.WriteString(lessonStyle.Render(currentProblem.Description) + "\n")
		mainContent.WriteString(lessonStyle.Render("Examples:\n" + currentProblem.Examples) + "\n\n")
		mainContent.WriteString(lipgloss.NewStyle().
			PaddingLeft(1).
			Render(m.input.View()))

		leftCol := mainContentStyle.
			Width(leftColumnWidth - 6).
			Render(mainContent.String())

		if m.err != nil {
			leftCol = lipgloss.JoinVertical(lipgloss.Left,
				leftCol,
				lipgloss.NewStyle().
					PaddingLeft(2).
					PaddingTop(1).
					Render(errorStyle.Render(m.err.Error())),
			)
		}

		// Right column (Practice Problems List)
		tocStyle := tocStyle.Copy().Width(rightColumnWidth - 6)
		var toc strings.Builder
		toc.WriteString(gradientText("Practice Problems") + "\n\n")

		for i, p := range m.practices {
			status := "○"
			style := incompletedStyle
			
			if p.Completed {
				status = "✓"
				style = completedStyle
			}

			problemTitle := fmt.Sprintf("%s Problem %d: %s", status, i+1, p.Title)
			if i == m.practiceIndex {
				problemTitle += " (current)"
				style = style.Bold(true)
				if !p.Completed {
					status = "○"
				}
				problemTitle = fmt.Sprintf("%s Problem %d: %s (current)", status, i+1, p.Title)
			}

			toc.WriteString(style.Render(problemTitle) + "\n")
		}

		rightCol := tocStyle.Render(toc.String())

		// Join columns
		columns := lipgloss.JoinHorizontal(lipgloss.Top,
			leftCol,
			rightCol,
		)

		doc.WriteString(lipgloss.NewStyle().
			Align(lipgloss.Center).
			Render(columns))

		doc.WriteString("\n\nPress ctrl+r to reset progress • tab to skip problem • shift+tab for previous problem • esc for main menu\n")

		return docStyle.Copy().Width(totalWidth).Render(doc.String())
	}

	// Main tutorial view
	doc := strings.Builder{}
	doc.WriteString(header + "\n")

	// Left column content
	var mainContent strings.Builder
	currentLesson := m.lessons[m.current]
	mainContent.WriteString(titleStyle.Render(currentLesson.Title) + "\n\n")
	mainContent.WriteString(lessonStyle.Render(currentLesson.Description) + "\n")
	mainContent.WriteString(lessonStyle.Render(currentLesson.Task) + "\n\n")
	mainContent.WriteString(lipgloss.NewStyle().
		PaddingLeft(1).
		Render(m.input.View()))

	leftCol := mainContentStyle.
		Width(leftColumnWidth - 6).  // Account for borders and margin
		Render(mainContent.String())

	if m.err != nil {
		leftCol = lipgloss.JoinVertical(lipgloss.Left,
			leftCol,
			lipgloss.NewStyle().
				PaddingLeft(2).
				PaddingTop(1).
				Render(errorStyle.Render(m.err.Error())),
		)
	}

	// Right column (Table of Contents)
	tocStyle := tocStyle.Copy().Width(rightColumnWidth - 6)  // Account for borders and margin
	var toc strings.Builder
	toc.WriteString(gradientText("Table of Contents") + "\n\n")

	for i, l := range m.lessons {
		status := "○"
		style := incompletedStyle
		
		if l.Completed {
			status = "✓"
			style = completedStyle
		}

		lessonTitle := fmt.Sprintf("%s Lesson %d: %s", status, i+1, l.Title)
		if i == m.current {
			lessonTitle += " (current)"
			style = style.Bold(true)
			if !l.Completed {
				status = "○"
			}
			lessonTitle = fmt.Sprintf("%s Lesson %d: %s (current)", status, i+1, l.Title)
		}

		toc.WriteString(style.Render(lessonTitle) + "\n")
	}

	rightCol := tocStyle.Render(toc.String())

	// Join columns with proper spacing
	columns := lipgloss.JoinHorizontal(lipgloss.Top,
		leftCol,
		rightCol,
	)

	doc.WriteString(lipgloss.NewStyle().
		Align(lipgloss.Center).
		Render(columns))

	doc.WriteString("\n\nPress ctrl+r to reset progress • tab to skip lesson • shift+tab for previous lesson • esc for main menu\n")

	return docStyle.Copy().Width(totalWidth).Render(doc.String())
}

// Add this function to create a new model while preserving dimensions
func resetModel(width, height int) model {
	m := initialModel()
	m.width = width
	m.height = height
	return m
}

func clearSpecificProgress(_ model, clearType string) error {
	progress, err := storage.LoadProgress()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		progress = models.Progress{}
	}

	if clearType == "practice" {
		progress.PracticeIndex = 0
		progress.CompletedPractice = nil
	} else if clearType == "learning" {
		progress.CurrentLesson = 0
		progress.Completed = nil
	}

	data, err := json.Marshal(progress)
	if err != nil {
		return err
	}
	
	return os.WriteFile(progressFile, data, 0644)
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}