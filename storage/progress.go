package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ghousemohamed/regex-in-the-terminal/models"
)

var progressFile = filepath.Join(os.Getenv("HOME"), ".regex_tutorial_progress.json")

func SaveProgress(current int, practiceIndex int, lessons []models.Lesson, practices []models.PracticeProblem) error {
	var completed []string
	var completedPractice []string

	for i, l := range lessons {
		if l.Completed {
			completed = append(completed, fmt.Sprintf("%d", i))
		}
	}

	for i, p := range practices {
		if p.Completed {
			completedPractice = append(completedPractice, fmt.Sprintf("%d", i))
		}
	}

	progress := models.Progress{
		CurrentLesson:     current,
		Completed:         completed,
		PracticeIndex:     practiceIndex,
		CompletedPractice: completedPractice,
	}

	data, err := json.Marshal(progress)
	if err != nil {
		return err
	}

	return os.WriteFile(progressFile, data, 0644)
}

func LoadProgress() (models.Progress, error) {
	var progress models.Progress
	data, err := os.ReadFile(progressFile)
	if err != nil {
		if os.IsNotExist(err) {
			return progress, nil
		}
		return progress, err
	}

	err = json.Unmarshal(data, &progress)
	return progress, err
}

func ClearSpecificProgress(clearType string) error {
	progress, err := LoadProgress()
	if err != nil && !os.IsNotExist(err) {
		return err
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