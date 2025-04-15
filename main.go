package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

// Отметить задачу как выполненую
func markComplete(number int) {
	if number >= 0 && number < len(List) {
		List[number].Completed = true
		fmt.Println("Marked as complete.")
	} else {
		fmt.Println("Invalid number.")
	}
}

// Удалить задачу из списка
func deleteTask(number int) {
	if number > 0 && number <= len(List) {
		List = append(List[:number], List[number+1:]...)
		fmt.Println("Task deleted.")
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
		fmt.Println("5. Exit")
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
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
