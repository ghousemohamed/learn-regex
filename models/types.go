package models

type TestCase struct {
	Text     string
	Expected bool
}

type Lesson struct {
	Title       string
	Description string
	Task        string
	TestCases   []TestCase
	Completed   bool
}

type PracticeProblem struct {
	Title       string
	Description string
	Examples    string
	TestCases   []TestCase
	Completed   bool
}

type CompletionState int

const (
	Welcome CompletionState = iota
	Learning
	Practicing
	Completed
	Success
)

type WelcomeOption int

const (
	StartLearning WelcomeOption = iota
	Practice
	Quit
)

type Progress struct {
	CurrentLesson     int      `json:"current_lesson"`
	Completed         []string `json:"completed_lessons"`
	PracticeIndex     int      `json:"practice_index"`
	CompletedPractice []string `json:"completed_practice"`
}