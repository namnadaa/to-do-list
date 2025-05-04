package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"todolist/color"
	"todolist/storage"
	"todolist/task"
	"todolist/toggle"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	err := storage.LoadTasks("tasks.json")
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

		input := storage.ReadInput(reader)

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			title := storage.ReadInput(reader)
			storage.WithSave(func() {
				task.AddTask(title)
			})
			fmt.Printf(color.Green("Task #%d added!\n"), len(task.List))
		case "2":
			fmt.Println(color.Blue("\n=== Task list ==="))
			task.ShowList()
		case "3":
			storage.WithSave(func() {
				toggle.ToggleMenu(reader)
			})
		case "4":
			fmt.Print("Enter the task number: ")
			number := storage.ReadInput(reader)
			n, err := storage.ConvertValue(number)
			if err != nil {
				fmt.Printf(color.Magenta("Error: %s\n"), err)
				break
			}

			fmt.Printf("%s #%d%s\n", color.Blue("You are about to delete task"), n, color.Blue("."))
			fmt.Print(color.Yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(storage.ReadInput(reader))

			if confirm == "y" {
				storage.WithSave(func() {
					task.DeleteTask(n - 1)
				})
			} else if confirm == "n" {
				fmt.Println(color.Red("Action canceled."))
			} else {
				fmt.Println(color.Red("Invalid choice, please enter 'y' or 'n'."))
			}
		case "5":
			fmt.Print("Enter the task number: ")
			number := storage.ReadInput(reader)
			n, err := storage.ConvertValue(number)
			if err != nil {
				fmt.Printf(color.Magenta("Error: %s\n"), err)
				break
			}

			fmt.Print("Enter new task text: ")
			newText := storage.ReadInput(reader)

			fmt.Printf("%s #%d %s \"%s\"%s\n", color.Blue("You are about to change task"), n, color.Blue("to:"), newText, color.Blue("."))
			fmt.Print(color.Yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(storage.ReadInput(reader))

			if confirm == "y" {
				storage.WithSave(func() {
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
