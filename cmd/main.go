package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"todolist/pkg/color"
	"todolist/pkg/filemenu"
	"todolist/pkg/history"
	"todolist/pkg/show"
	"todolist/pkg/storage"
	"todolist/pkg/task"
	"todolist/pkg/toggle"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	err := storage.LoadTasks("tasks.json")
	if err != nil {
		msg := fmt.Sprintf("[ERROR] Failed to load tasks: %v", err)
		log.Fatal(color.Magenta(msg))
	}

	for {
		fmt.Println(color.Blue("\n=== To-Do Menu ==="))
		fmt.Println(color.Blue("1.") + " Add task")
		fmt.Println(color.Blue("2.") + " Show tasks")
		fmt.Println(color.Blue("3.") + " Toggle menu")
		fmt.Println(color.Blue("4.") + " Delete task")
		fmt.Println(color.Blue("5.") + " Edit task")
		fmt.Println(color.Blue("6.") + " Undo action")
		fmt.Println(color.Blue("7.") + " File menu")
		fmt.Println(color.Blue("8.") + " Exit")
		fmt.Print(color.Blue("\nChoose an action: "))

		input := storage.ReadInput(reader)

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			title := storage.ReadInput(reader)
			storage.WithSave(func() {
				task.AddTask(title)
				history.Record(history.Action{
					Type:     history.Add,
					TaskData: task.List[len(task.List)-1],
				})
			})
			msg := fmt.Sprintf("Task #%d added!\n", len(task.List))
			fmt.Print(color.Green(msg))
		case "2":
			show.ShowMenu(reader)
		case "3":
			toggle.ToggleMenu(reader)
		case "4":
			fmt.Print("Enter the task number: ")
			number := storage.ReadInput(reader)
			n, err := storage.ConvertValue(number)
			if err != nil {
				msg := fmt.Sprintf("Error: %v\n", err)
				fmt.Print(color.Magenta(msg))
				break
			}

			fmt.Printf("%s #%d%s\n", color.Blue("You are about to delete task"), n, color.Blue("."))
			fmt.Print(color.Yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(storage.ReadInput(reader))

			if confirm == "y" {
				storage.WithSave(func() {
					deleted := task.List[n-1]
					task.DeleteTask(n - 1)
					history.Record(history.Action{
						Type:     history.Delete,
						TaskData: deleted,
						Index:    n - 1,
					})
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
				msg := fmt.Sprintf("Error: %v\n", err)
				fmt.Print(color.Magenta(msg))
				break
			}

			fmt.Print("Enter new task text: ")
			newText := storage.ReadInput(reader)

			fmt.Printf("%s #%d %s \"%s\"%s\n", color.Blue("You are about to change task"), n, color.Blue("to:"), newText, color.Blue("."))
			fmt.Print(color.Yellow("Are you sure? (y/n): "))
			confirm := strings.ToLower(storage.ReadInput(reader))

			if confirm == "y" {
				oldText := task.List[n-1].Task
				storage.WithSave(func() {
					task.TaskEditing(n-1, newText)
					history.Record(history.Action{
						Type:     history.Edit,
						Index:    n - 1,
						PrevText: oldText,
					})
				})
			} else if confirm == "n" {
				fmt.Println(color.Red("Task not changed."))
			} else {
				fmt.Println(color.Red("Invalid choice, please enter 'y' or 'n'."))
			}
		case "6":
			storage.WithSave(func() {
				history.Undo()
			})
		case "7":
			filemenu.FileMenu(reader)
		case "8":
			fmt.Println(color.Blue("Exiting..."))
			return
		default:
			fmt.Println(color.Red("Invalid choice. Please try again."))
		}
	}
}
