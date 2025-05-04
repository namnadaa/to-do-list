package main

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
var autosaveEnable = false

// toggleMenu displays a submenu for toggling task statuses.
func toggleMenu(reader *bufio.Reader) {
	for {
		fmt.Println(color.Blue("\n=== Toggle Menu ==="))
		fmt.Println(color.Blue("1.") + " Mark task")
		fmt.Println(color.Blue("2.") + " Unmark task")
		fmt.Println(color.Blue("3.") + " Mark all")
		fmt.Println(color.Blue("4.") + " Unmark all")
		fmt.Println(color.Blue("5.") + " Back to menu")
		fmt.Print(color.Blue("\nChoose an action: "))

		input := readInput(reader)

		switch input {
		case "1":
			markSingleTask(reader)
		case "2":
			unmarkSingleTask(reader)
		case "3":
			markAllTask()
		case "4":
			unmarkAllTask()
		case "5":
			return
		default:
			fmt.Println(color.Red("Invalid choice. Please try again."))
		}
	}
}

// markSingleTask marks a single task as completed.
func markSingleTask(reader *bufio.Reader) {
	fmt.Print("Enter the task number: ")
	input := readInput(reader)

	number, err := convertValue(input)
	if err != nil {
		fmt.Printf(color.Magenta("Error: %s\n"), err)
		return
	}

	if number >= 1 && number <= len(task.List) {
		if !task.List[number-1].Completed {
			task.List[number-1].Completed = true
			fmt.Println(color.Green("Task marked as complete."))
		} else {
			fmt.Println(color.Yellow("Task is already marked as completed."))
		}
	} else {
		fmt.Println(color.Red("Invalid task number."))
	}
}

// unmarkSingleTask marks a single task as not completed.
func unmarkSingleTask(reader *bufio.Reader) {

	fmt.Print("Enter the task number: ")
	input := readInput(reader)

	number, err := convertValue(input)
	if err != nil {
		fmt.Printf(color.Magenta("Error: %s\n"), err)
		return
	}

	if number >= 1 && number <= len(task.List) {
		if task.List[number-1].Completed {
			task.List[number-1].Completed = false
			fmt.Println(color.Green("Task marked as not complete."))
		} else {
			fmt.Println(color.Yellow("Task is already not marked as completed."))
		}
	} else {
		fmt.Println(color.Red("Invalid task number."))
	}
}

// markAllTask marks all tasks in the list as completed.
func markAllTask() {
	var count int
	for i := range task.List {
		if !task.List[i].Completed {
			task.List[i].Completed = true
			count++
		}
	}
	fmt.Printf(color.Green("Marked %d task(s) as completed.\n"), count)
}

// unmarkAllTask marks all tasks in the list as not completed.
func unmarkAllTask() {
	var count int
	for i := range task.List {
		if task.List[i].Completed {
			task.List[i].Completed = false
			count++
		}
	}
	fmt.Printf(color.Green("Marked %d task(s) as not completed.\n"), count)
}

// readInput reads a line of input from the console and trims whitespace.
func readInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%s", color.Magenta(fmt.Sprintf("Error reading input: %v", err)))
	}
	return strings.TrimSpace(input)
}

// convertValue converts a string input into an integer.
func convertValue(number string) (int, error) {
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

// loadTasks reads a file and deserializes its content into the task list.
func loadTasks(fileName string) error {
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

// withSave executes an action and saves the task list if autosave is enabled.
func withSave(action func()) {
	action()

	if autosaveEnable {
		err := saveTasks("tasks.json")
		if err != nil {
			log.Printf(color.Magenta("[ERROR] Autosave failed: %v"), err)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	err := loadTasks("tasks.json")
	if err != nil {
		log.Fatalf("%s", color.Magenta(fmt.Sprintf("[ERROR] Failed to load tasks: %v", err)))
	}

	for {
		fmt.Println(color.Blue("\n=== To-Do Menu ==="))
		fmt.Println(color.Blue("1.") + " Add task")
		fmt.Println(color.Blue("2.") + " Show tasks")
		fmt.Println(color.Blue("3.") + " Toggle menu")
		fmt.Println(color.Blue("4.") + " Delete task")
		fmt.Println(color.Blue("5.") + " Edit task")
		fmt.Println(color.Blue("6.") + " Exit")
		fmt.Print(color.Blue("\nChoose an action: "))

		input := readInput(reader)

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			title := readInput(reader)
			withSave(func() {
				task.AddTask(title)
			})
			fmt.Printf(color.Green("Task #%d added!\n"), len(task.List))
		case "2":
			fmt.Println(color.Blue("\n=== Task list ==="))
			task.ShowList()
		case "3":
			withSave(func() {
				toggleMenu(reader)
			})
		case "4":
			fmt.Print("Enter the task number: ")
			number := readInput(reader)
			n, err := convertValue(number)
			if err != nil {
				fmt.Printf(color.Magenta("Error: %s\n"), err)
				break
			}

			fmt.Printf("%s #%d%s\n", color.Blue("You are about to delete task"), n, color.Blue("."))
			fmt.Print(color.Yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(readInput(reader))

			if confirm == "y" {
				withSave(func() {
					task.DeleteTask(n - 1)
				})
			} else if confirm == "n" {
				fmt.Println(color.Red("Action canceled."))
			} else {
				fmt.Println(color.Red("Invalid choice, please enter 'y' or 'n'."))
			}
		case "5":
			fmt.Print("Enter the task number: ")
			number := readInput(reader)
			n, err := convertValue(number)
			if err != nil {
				fmt.Printf(color.Magenta("Error: %s\n"), err)
				break
			}

			fmt.Print("Enter new task text: ")
			newText := readInput(reader)

			fmt.Printf("%s #%d %s \"%s\"%s\n", color.Blue("You are about to change task"), n, color.Blue("to:"), newText, color.Blue("."))
			fmt.Print(color.Yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(readInput(reader))

			if confirm == "y" {
				withSave(func() {
					task.TaskEditing(n-1, newText)
				})
			} else if confirm == "n" {
				fmt.Println(color.Red("Task not changed."))
			} else {
				fmt.Println(color.Red("Invalid choice, please enter 'y' or 'n'."))
			}
		case "6":
			fmt.Println(color.Blue("Exiting..."))
			return
		default:
			fmt.Println(color.Red("Invalid choice. Please try again."))
		}
	}
}
