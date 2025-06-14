package show

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
	"todolist/pkg/color"
	"todolist/pkg/history"
	"todolist/pkg/storage"
	"todolist/pkg/task"
)

// ShowMenu displays the list of tasks and presents a submenu with additional options.
func ShowMenu(reader *bufio.Reader) {
	for {
		fmt.Println(color.Blue("\n=== Task List ==="))
		showList()
		fmt.Println(color.Blue("\n--- Show Menu ---"))
		fmt.Println(color.Blue("1.") + "Sort by completed")
		fmt.Println(color.Blue("2.") + "Back to menu")
		fmt.Print(color.Blue("\nChoose an action: "))

		input := storage.ReadInput(reader)

		switch input {
		case "1":
			storage.WithSave(func() {
				sortTask()
			})
		case "2":
			return
		default:
			fmt.Println(color.Red("Invalid choice. Please try again."))
		}
	}
}

// ShowList displays all tasks and a visual progress bar.
func showList() {
	var count int

	msg := fmt.Sprintf("%-4s%-8s%-s\n", "#", "Status", "Task")
	fmt.Print(color.Blue(msg))

	for i, task := range task.List {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
			count++
		}

		number := fmt.Sprintf("%-4d", i+1)
		fmt.Printf("%s%-8s%-s\n", color.Blue(number), status, task.Task)
	}

	progressBar(count)
}

// colorProgressBar returns a colored version of the progress bar based on the ratio.
func colorProgressBar(progressRatio float64, bar string) string {
	percent := progressRatio * 100

	switch {
	case percent < 33:
		return color.Red(bar)
	case percent <= 66:
		return color.Yellow(bar)
	default:
		return color.Green(bar)
	}
}

// ProgressBar displays a visual representation of task completion status.
func progressBar(count int) {
	fmt.Println(color.Blue("\nProgress:"))

	barWidth := 10

	if len(task.List) == 0 {
		fmt.Println(color.Red("[----------]") + " 0.0% " + " (0/0)")
		return
	}

	progressRatio := float64(count) / float64(len(task.List))
	filled := int(progressRatio * float64(barWidth))
	progressBar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", barWidth-filled) + "]"
	fmt.Printf("%s %.1f%%  (%d/%d)\n", colorProgressBar(progressRatio, progressBar), progressRatio*100, count, len(task.List))
}

// sortTask sorts the task list so that completed tasks appear before uncompleted ones.
func sortTask() {
	prev := make([]task.Task, len(task.List))
	copy(prev, task.List)

	alreadySort := sort.SliceIsSorted(task.List, func(i, j int) bool {
		return task.List[i].Completed && !task.List[j].Completed
	})

	if alreadySort {
		fmt.Println(color.Yellow("List is already sorted."))
		return
	}

	sort.Slice(task.List, func(i, j int) bool {
		return task.List[i].Completed && !task.List[j].Completed
	})

	history.Record(history.Action{
		Type:      history.Sort,
		PrevState: prev,
	})

	fmt.Println(color.Green("List sorted."))
}
