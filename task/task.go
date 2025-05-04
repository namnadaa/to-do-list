package task

import (
	"fmt"
	"strconv"
	"strings"
	"todolist/color"
)

// Task represents a single to-do item with a title and completion status.
type Task struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// List stores all tasks in memory.
var List []Task

// AddTask appends a new task to the task list.
func AddTask(task string) {
	List = append(List, Task{Task: task, Completed: false})
}

// ShowList displays all tasks and a visual progress bar.
func ShowList() {
	var count int

	fmt.Printf("%s  %-7s %-s\n", color.Blue("#"), color.Blue("Status"), color.Blue("Task"))

	for i, task := range List {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
			count++
		}

		number := color.Blue(strconv.Itoa(i + 1))
		fmt.Printf("%-3s  %-6s %-s\n", number, status, task.Task)
	}

	ProgressBar(count)
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
func ProgressBar(count int) {
	fmt.Println(color.Blue("\nProgress:"))

	barWidth := 10

	if len(List) == 0 {
		fmt.Println(color.Red("[----------]") + " 0.0% " + " (0/0)")
		return
	}

	progressRatio := float64(count) / float64(len(List))
	filled := int(progressRatio * float64(barWidth))
	progressBar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", barWidth-filled) + "]"
	fmt.Printf("%s %.1f%%  (%d/%d)\n", colorProgressBar(progressRatio, progressBar), progressRatio*100, count, len(List))
}

// DeleteTask removes a task from the list by its index.
func DeleteTask(number int) {
	if number >= 0 && number < len(List) {
		List = append(List[:number], List[number+1:]...)
		fmt.Println(color.Green("Task deleted."))
	} else {
		fmt.Println(color.Red("Invalid task number."))
	}
}

// TaskEditing modifies the text of a task by its index.
func TaskEditing(number int, task string) {
	if number >= 0 && number < len(List) {
		List[number].Task = task
		fmt.Println(color.Green("Task updated."))
	} else {
		fmt.Println(color.Red("Invalid task number."))
	}
}
