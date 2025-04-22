package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Task structure
type Task struct {
	Task      string `json:"task"`
	Completed bool   `json:"done"`
}

// Task list
var List []Task

// Add a task
func addTask(task string) {
	List = append(List, Task{Task: task, Completed: false})
}

// Show the list
func showList() {
	var count int

	for i, task := range List {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
			count++
		}
		fmt.Printf("%s %d. %s\n", status, i+1, task.Task)
	}
	fmt.Printf("\nCompleted: %d/%d\n", count, len(List))

	barWidth := 10
	progressRatio := float64(count) / float64(len(List))
	filled := int(progressRatio * float64(barWidth))
	progressBar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", barWidth-filled) + "]"
	fmt.Printf("%s\n", progressBar)
}

// Submenu toggle task status
func toggleMenu(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Toggle Menu ---")
		fmt.Println("1. Mark one task")
		fmt.Println("2. Unmark one task")
		fmt.Println("3. Mark all tasks as completed")
		fmt.Println("4. Unmark all tasks")
		fmt.Println("5. Back to main menu")
		fmt.Print("\nChoose an action: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		input = strings.TrimSpace(input)

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
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Mark a task as completed
func markSingleTask(reader *bufio.Reader) {
	fmt.Print("Enter the task number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	input = strings.TrimSpace(input)

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Only a number can be entered: %v\n", err)
		return
	}

	if number >= 1 && number <= len(List) {
		if !List[number-1].Completed {
			List[number-1].Completed = true
			fmt.Println("Marked as complete.")
		} else {
			fmt.Println("The task is already marked as completed.")
		}
	} else {
		fmt.Println("Invalid number.")
	}
}

// Unmark a task as not completed
func unmarkSingleTask(reader *bufio.Reader) {

	fmt.Print("Enter the task number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	input = strings.TrimSpace(input)

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Only a number can be entered: %v\n", err)
		return
	}

	if number >= 1 && number <= len(List) {
		if List[number-1].Completed {
			List[number-1].Completed = false
			fmt.Println("Marked as not complete.")
		} else {
			fmt.Println("The task is already not marked as completed.")
		}
	} else {
		fmt.Println("Invalid number.")
	}
}

// Mark all task as completed
func markAllTask() {
	var count int
	for i := range List {
		if !List[i].Completed {
			List[i].Completed = true
			count++
		}
	}
	fmt.Printf("Marked %d task(s) as completed.\n", count)
}

// Unmark all task as completed
func unmarkAllTask() {
	var count int
	for i := range List {
		if List[i].Completed {
			List[i].Completed = false
			count++
		}
	}
	fmt.Printf("Marked %d task(s) as not completed.\n", count)
}

// Delete a task from the list
func deleteTask(number int) {
	if number >= 0 && number < len(List) {
		List = append(List[:number], List[number+1:]...)
		fmt.Println("Task deleted.")
	} else {
		fmt.Println("Invalid number.")
	}
}

// Edit task by number
func taskEditing(number int, task string) {
	if number >= 0 && number < len(List) {
		List[number].Task = task
		fmt.Println("Task modified.")
	} else {
		fmt.Println("Invalid number.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- To-Do Menu ---")
		fmt.Println("1. Add a task")
		fmt.Println("2. Show task list")
		fmt.Println("3. Toggle menu")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Edit a task")
		fmt.Println("6. Exit")
		fmt.Print("\nChoose an action: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			title, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading title: %v", err)
			}
			title = strings.TrimSpace(title)

			addTask(title)
			fmt.Printf("Task #%d added!\n", len(List))
		case "2":
			fmt.Println("\nTask list:")
			showList()
		case "3":
			toggleMenu(reader)
		case "4":
			fmt.Print("Enter the task number: ")
			number, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading number: %v", err)
			}
			number = strings.TrimSpace(number)

			n, err := strconv.Atoi(number)
			if err != nil {
				fmt.Printf("Only a number can be entered: %v\n", err)
				break
			}

			fmt.Printf("You are about to delete task #%d\n", n)
			fmt.Print("Are you sure? (y/n): ")
			confirm, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading title: %v", err)
			}
			confirm = strings.TrimSpace(strings.ToLower(confirm))

			if confirm == "y" {
				deleteTask(n - 1)
			} else if confirm == "n" {
				fmt.Println("Task deletion canceled.")
			} else {
				fmt.Println("Invalid choice, please enter 'y' or 'n'.")
			}
		case "5":
			fmt.Print("Enter the task number: ")
			number, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading number: %v", err)
			}
			number = strings.TrimSpace(number)

			n, err := strconv.Atoi(number)
			if err != nil {
				fmt.Printf("Only a number can be entered: %v\n", err)
				break
			}

			fmt.Print("Enter new task text: ")
			newText, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading title: %v", err)
			}
			newText = strings.TrimSpace(newText)

			fmt.Printf("You are about to change task #%d to: \"%s\"\n", n, newText)
			fmt.Print("Are you sure? (y/n): ")
			confirm, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading title: %v", err)
			}
			confirm = strings.TrimSpace(strings.ToLower(confirm))

			if confirm == "y" {
				taskEditing(n-1, newText)
			} else if confirm == "n" {
				fmt.Println("Changes have been canceled.")
			} else {
				fmt.Println("Invalid choice, please enter 'y' or 'n'.")
			}
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
