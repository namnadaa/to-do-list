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
)

// autosaveEnable indicates whether autosave is enabled for operations wrapped in withSave.
var autosaveEnable = false

// ANSI color codes used for colored console output.
var (
	Blue    = "\033[34m"       // Menu headers and system messages
	Green   = "\033[32m"       // Successful actions and confirmations
	Yellow  = "\033[38;5;214m" // Warnings and user prompts
	Red     = "\033[31m"       // Cancellations and denials
	Magenta = "\033[35m"       // Errors and invalid actions
	Reset   = "\033[0m"        // Reset to default color
)

// blue returns the input text wrapped in ANSI blue color codes.
func blue(text string) string {
	return Blue + text + Reset
}

// green returns the input text wrapped in ANSI green color codes.
func green(text string) string {
	return Green + text + Reset
}

// yellow returns the input text wrapped in ANSI yellow color codes.
func yellow(text string) string {
	return Yellow + text + Reset
}

// red returns the input text wrapped in ANSI red color codes.
func red(text string) string {
	return Red + text + Reset
}

// magenta returns the input text wrapped in ANSI magenta color codes.
func magenta(text string) string {
	return Magenta + text + Reset
}

// Task represents a single to-do item with a title and completion status.
type Task struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// List stores all tasks in memory.
var List []Task

// addTask appends a new task to the task list.
func addTask(task string) {
	List = append(List, Task{Task: task, Completed: false})
}

// showList displays all tasks and a visual progress bar.
func showList() {
	var count int

	fmt.Printf("%s  %-7s %-s\n", blue("#"), blue("Status"), blue("Task"))

	for i, task := range List {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
			count++
		}

		number := blue(strconv.Itoa(i + 1))
		fmt.Printf("%-3s  %-6s %-s\n", number, status, task.Task)
	}

	progressBar(count)
}

// colorProgressBar returns a colored version of the progress bar based on the ratio.
func colorProgressBar(progressRatio float64, bar string) string {
	percent := progressRatio * 100

	switch {
	case percent < 33:
		return red(bar)
	case percent <= 66:
		return yellow(bar)
	default:
		return green(bar)
	}
}

// progressBar displays a visual representation of task completion status.
func progressBar(count int) {
	fmt.Println(blue("\nProgress:"))

	barWidth := 10

	if len(List) == 0 {
		fmt.Println(red("[----------]") + " 0.0% " + " (0/0)")
		return
	}

	progressRatio := float64(count) / float64(len(List))
	filled := int(progressRatio * float64(barWidth))
	progressBar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", barWidth-filled) + "]"
	fmt.Printf("%s %.1f%%  (%d/%d)\n", colorProgressBar(progressRatio, progressBar), progressRatio*100, count, len(List))
}

// toggleMenu displays a submenu for toggling task statuses.
func toggleMenu(reader *bufio.Reader) {
	for {
		fmt.Println(blue("\n=== Toggle Menu ==="))
		fmt.Println(blue("1.") + " Mark task")
		fmt.Println(blue("2.") + " Unmark task")
		fmt.Println(blue("3.") + " Mark all")
		fmt.Println(blue("4.") + " Unmark all")
		fmt.Println(blue("5.") + " Back to menu")
		fmt.Print(blue("\nChoose an action: "))

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
			fmt.Println(red("Invalid choice. Please try again."))
		}
	}
}

// markSingleTask marks a single task as completed.
func markSingleTask(reader *bufio.Reader) {
	fmt.Print("Enter the task number: ")
	input := readInput(reader)

	number, err := convertValue(input)
	if err != nil {
		fmt.Printf(magenta("Error: %s\n"), err)
		return
	}

	if number >= 1 && number <= len(List) {
		if !List[number-1].Completed {
			List[number-1].Completed = true
			fmt.Println(green("Task marked as complete."))
		} else {
			fmt.Println(yellow("Task is already marked as completed."))
		}
	} else {
		fmt.Println(red("Invalid task number."))
	}
}

// unmarkSingleTask marks a single task as not completed.
func unmarkSingleTask(reader *bufio.Reader) {

	fmt.Print("Enter the task number: ")
	input := readInput(reader)

	number, err := convertValue(input)
	if err != nil {
		fmt.Printf(magenta("Error: %s\n"), err)
		return
	}

	if number >= 1 && number <= len(List) {
		if List[number-1].Completed {
			List[number-1].Completed = false
			fmt.Println(green("Task marked as not complete."))
		} else {
			fmt.Println(yellow("Task is already not marked as completed."))
		}
	} else {
		fmt.Println(red("Invalid task number."))
	}
}

// markAllTask marks all tasks in the list as completed.
func markAllTask() {
	var count int
	for i := range List {
		if !List[i].Completed {
			List[i].Completed = true
			count++
		}
	}
	fmt.Printf(green("Marked %d task(s) as completed.\n"), count)
}

// unmarkAllTask marks all tasks in the list as not completed.
func unmarkAllTask() {
	var count int
	for i := range List {
		if List[i].Completed {
			List[i].Completed = false
			count++
		}
	}
	fmt.Printf(green("Marked %d task(s) as not completed.\n"), count)
}

// deleteTask removes a task from the list by its index.
func deleteTask(number int) {
	if number >= 0 && number < len(List) {
		List = append(List[:number], List[number+1:]...)
		fmt.Println(green("Task deleted."))
	} else {
		fmt.Println(red("Invalid task number."))
	}
}

// taskEditing modifies the text of a task by its index.
func taskEditing(number int, task string) {
	if number >= 0 && number < len(List) {
		List[number].Task = task
		fmt.Println(green("Task updated."))
	} else {
		fmt.Println(red("Invalid task number."))
	}
}

// readInput reads a line of input from the console and trims whitespace.
func readInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("%s", magenta(fmt.Sprintf("Error reading input: %v", err)))
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
	data, err := json.MarshalIndent(List, "", "  ")
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
			List = make([]Task, 0)
			return nil
		} else {
			return fmt.Errorf("task loading error: %v", err)
		}
	}

	err = json.Unmarshal(data, &List)
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
			log.Printf(magenta("[ERROR] Autosave failed: %v"), err)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	err := loadTasks("tasks.json")
	if err != nil {
		log.Fatalf("%s", magenta(fmt.Sprintf("[ERROR] Failed to load tasks: %v", err)))
	}

	for {
		fmt.Println(blue("\n=== To-Do Menu ==="))
		fmt.Println(blue("1.") + " Add task")
		fmt.Println(blue("2.") + " Show tasks")
		fmt.Println(blue("3.") + " Toggle menu")
		fmt.Println(blue("4.") + " Delete task")
		fmt.Println(blue("5.") + " Edit task")
		fmt.Println(blue("6.") + " Exit")
		fmt.Print(blue("\nChoose an action: "))

		input := readInput(reader)

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			title := readInput(reader)
			withSave(func() {
				addTask(title)
			})
			fmt.Printf(green("Task #%d added!\n"), len(List))
		case "2":
			fmt.Println(blue("\n=== Task list ==="))
			showList()
		case "3":
			withSave(func() {
				toggleMenu(reader)
			})
		case "4":
			fmt.Print("Enter the task number: ")
			number := readInput(reader)
			n, err := convertValue(number)
			if err != nil {
				fmt.Printf(magenta("Error: %s\n"), err)
				break
			}

			fmt.Printf("%s #%d%s\n", blue("You are about to delete task"), n, blue("."))
			fmt.Print(yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(readInput(reader))

			if confirm == "y" {
				withSave(func() {
					deleteTask(n - 1)
				})
			} else if confirm == "n" {
				fmt.Println(red("Action canceled."))
			} else {
				fmt.Println(red("Invalid choice, please enter 'y' or 'n'."))
			}
		case "5":
			fmt.Print("Enter the task number: ")
			number := readInput(reader)
			n, err := convertValue(number)
			if err != nil {
				fmt.Printf(magenta("Error: %s\n"), err)
				break
			}

			fmt.Print("Enter new task text: ")
			newText := readInput(reader)

			fmt.Printf("%s #%d %s \"%s\"%s\n", blue("You are about to change task"), n, blue("to:"), newText, blue("."))
			fmt.Print(yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(readInput(reader))

			if confirm == "y" {
				withSave(func() {
					taskEditing(n-1, newText)
				})
			} else if confirm == "n" {
				fmt.Println(red("Task not changed."))
			} else {
				fmt.Println(red("Invalid choice, please enter 'y' or 'n'."))
			}
		case "6":
			fmt.Println(blue("Exiting..."))
			return
		default:
			fmt.Println(red("Invalid choice. Please try again."))
		}
	}
}
