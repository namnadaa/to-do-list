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
	for i, task := range List {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%s %d. %s\n", status, i+1, task.Task)
	}
}

// Mark a task as completed
func markComplete(number int) {
	if number >= 0 && number < len(List) {
		List[number].Completed = true
		fmt.Println("Marked as complete.")
	} else {
		fmt.Println("Invalid number.")
	}
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
		fmt.Println("3. Mark task as completed")
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
			fmt.Println("Task list:")
			showList()
		case "3":
			fmt.Print("Enter the task number: ")
			number, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading number: %v", err)
			}
			number = strings.TrimSpace(number)
			n, err := strconv.Atoi(number)
			if err != nil {
				fmt.Printf("Only a number can be entered: %v\n", err)
			} else {
				markComplete(n - 1)
			}
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
			} else {
				deleteTask(n - 1)
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

			taskEditing(n-1, newText)
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
