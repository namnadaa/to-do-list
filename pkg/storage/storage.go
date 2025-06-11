package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
	"todolist/pkg/color"
	"todolist/pkg/task"
)

// autosaveEnable indicates whether autosave is enabled for operations wrapped in withSave.
var AutosaveEnable = true

// readInput reads a line of input from the console and trims whitespace.
func ReadInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		msg := fmt.Sprintf("Error reading input: %v", err)
		log.Fatal(color.Magenta(msg))
	}
	return strings.TrimSpace(input)
}

// ConvertValue converts a string input into an integer.
func ConvertValue(number string) (int, error) {
	n, err := strconv.Atoi(number)
	if err != nil {
		return 0, fmt.Errorf("only a number can be entered: %v", err)
	}
	return n, nil
}

// SetAutoSave toggles the autosave setting and displays a message
// indicating whether autosave has been enabled or disabled.
func SetAutoSave() {
	AutosaveEnable = !AutosaveEnable
	if AutosaveEnable {
		fmt.Println(color.Green("Autosave enebled."))
	} else {
		fmt.Println(color.Yellow("Autosave disabled."))
	}
}

// saveTasks serializes the task list and writes it to a file.
func saveTasks(fileName string) error {
	data, err := json.MarshalIndent(task.List, "", "  ")
	if err != nil {
		return fmt.Errorf("serialization error: %v", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("task saving error: %v", err)
	}
	return nil
}

// LoadTasks reads a file and deserializes its content into the task list.
func LoadTasks(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			task.List = make([]task.Task, 0)
			return nil
		} else {
			return fmt.Errorf("task loading error: %v", err)
		}
	}

	err = json.Unmarshal(data, &task.List)
	if err != nil {
		return fmt.Errorf("deserialization error: %v", err)
	}

	return nil
}

// WithSave executes an action and saves the task list if autosave is enabled.
func WithSave(action func()) {
	action()

	if AutosaveEnable {
		err := saveTasks("tasks.json")
		if err != nil {
			msg := fmt.Sprintf("[ERROR] Autosave failed: %v", err)
			log.Print(color.Magenta(msg))
		}
	}
}

// SaveAs saves the current task list to a user-defined file.
func SaveAs(reader *bufio.Reader) error {
	fmt.Print("Enter file name to save as: ")
	fileName := ReadInput(reader)

	if strings.TrimSpace(fileName) == "" {
		return fmt.Errorf("file name cannot be empty")
	}

	err := saveTasks(fileName)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	fmt.Println(color.Green("Tasks saved as:"), fileName)
	return nil
}

// ExportToText exports the task list to a plain text file.
func ExportToText(reader *bufio.Reader) error {
	fmt.Print("Enter file name to export: ")
	fileName := ReadInput(reader)

	if strings.TrimSpace(fileName) == "" {
		return fmt.Errorf("file name cannot be empty")
	}

	msg := fmt.Sprintf("%-8s%s", "Status", "Task")
	lines := []string{msg}

	for _, t := range task.List {
		status := "[ ]"
		if t.Completed {
			status = "[x]"
		}
		line := fmt.Sprintf("%-8s%s", status, t.Task)
		lines = append(lines, line)
	}

	data := strings.Join(lines, "\n")
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("task export error: %v", err)
	}

	fmt.Println(color.Green("Tasks exported to:"), fileName)
	return nil
}
