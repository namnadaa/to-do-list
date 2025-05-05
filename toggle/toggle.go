package toggle

import (
	"bufio"
	"fmt"
	"todolist/color"
	"todolist/history"
	"todolist/storage"
	"todolist/task"
)

// ToggleMenu displays a submenu for toggling task statuses.
func ToggleMenu(reader *bufio.Reader) {
	for {
		fmt.Println(color.Blue("\n=== Toggle Menu ==="))
		fmt.Println(color.Blue("1.") + " Mark task")
		fmt.Println(color.Blue("2.") + " Unmark task")
		fmt.Println(color.Blue("3.") + " Mark all")
		fmt.Println(color.Blue("4.") + " Unmark all")
		fmt.Println(color.Blue("5.") + " Back to menu")
		fmt.Print(color.Blue("\nChoose an action: "))

		input := storage.ReadInput(reader)

		switch input {
		case "1":
			markSingleTask(reader)
		case "2":
			unmarkSingleTask(reader)
		case "3":
			storage.WithSave(func() {
				markAllTask()
			})
		case "4":
			storage.WithSave(func() {
				unmarkAllTask()
			})
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
	input := storage.ReadInput(reader)

	number, err := storage.ConvertValue(input)
	if err != nil {
		fmt.Printf(color.Magenta("Error: %s\n"), err)
		return
	}

	if number >= 1 && number <= len(task.List) {
		if !task.List[number-1].Completed {
			history.Record(history.Action{
				Type:  history.Toggle,
				Index: number - 1,
			})

			storage.WithSave(func() {
				task.List[number-1].Completed = true
			})

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
	input := storage.ReadInput(reader)

	number, err := storage.ConvertValue(input)
	if err != nil {
		fmt.Printf(color.Magenta("Error: %s\n"), err)
		return
	}

	if number >= 1 && number <= len(task.List) {
		if task.List[number-1].Completed {
			history.Record(history.Action{
				Type:  history.Toggle,
				Index: number - 1,
			})

			storage.WithSave(func() {
				task.List[number-1].Completed = false
			})

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
			history.Record(history.Action{
				Type:  history.Toggle,
				Index: i,
			})

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
			history.Record(history.Action{
				Type:  history.Toggle,
				Index: i,
			})

			task.List[i].Completed = false
			count++
		}
	}
	fmt.Printf(color.Green("Marked %d task(s) as not completed.\n"), count)
}
