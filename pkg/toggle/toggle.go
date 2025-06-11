package toggle

import (
	"bufio"
	"fmt"
	"todolist/pkg/color"
	"todolist/pkg/history"
	"todolist/pkg/storage"
	"todolist/pkg/task"
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
		msg := fmt.Sprintf("Error: %v\n", err)
		fmt.Print(color.Magenta(msg))
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
		msg := fmt.Sprintf("Error: %v\n", err)
		fmt.Print(color.Magenta(msg))
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
	var subActions []history.Action

	for i := range task.List {
		if !task.List[i].Completed {
			subActions = append(subActions, history.Action{
				Type:  history.Toggle,
				Index: i,
			})

			task.List[i].Completed = true
			count++
		}
	}

	if count > 0 {
		history.Record(history.Action{
			Type:       history.Toggle,
			SubActions: subActions,
		})
		msg := fmt.Sprintf("Marked %d task(s) as completed.\n", count)
		fmt.Print(color.Green(msg))
	} else {
		fmt.Println(color.Yellow("All tasks are alredy completed."))
	}
}

// unmarkAllTask marks all tasks in the list as not completed.
func unmarkAllTask() {
	var count int
	var subActions []history.Action

	for i := range task.List {
		if task.List[i].Completed {
			subActions = append(subActions, history.Action{
				Type:  history.Toggle,
				Index: i,
			})

			task.List[i].Completed = false
			count++
		}
	}

	if count > 0 {
		history.Record(history.Action{
			Type:       history.Toggle,
			SubActions: subActions,
		})
		msg := fmt.Sprintf("Marked %d task(s) as not completed.\n", count)
		fmt.Print(color.Green(msg))
	} else {
		fmt.Println(color.Yellow("All tasks are alredy not completed."))
	}
}
