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
	"todolist/color"
	"todolist/task"
)

// autosaveEnable indicates whether autosave is enabled for operations wrapped in withSave.
var autosaveEnable = true

// readInput reads a line of input from the console and trims whitespace.
func ReadInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%s", color.Magenta(fmt.Sprintf("Error reading input: %v", err)))
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

	if autosaveEnable {
		err := saveTasks("tasks.json")
		if err != nil {
			log.Printf(color.Magenta("[ERROR] Autosave failed: %v"), err)
		}
	}
}
