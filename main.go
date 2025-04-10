package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Структура задачи
type Task struct {
	Task      string `json:"task"`
	Completed bool   `json:"done"`
}

// Список задач
var List []Task

// Добавление задачи
func addTask(task string) {
	List = append(List, Task{Task: task, Completed: false})
}

// Показать список
func showList() {
	for i, task := range List {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%s %d. %s\n", status, i+1, task.Task)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- To-Do Menu ---")
		fmt.Println("1. Add a task")
		fmt.Println("2. Show task list")
		fmt.Println("3. Exit")
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
			fmt.Println("Task added!")
		case "2":
			fmt.Println("Task list:")
			showList()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
