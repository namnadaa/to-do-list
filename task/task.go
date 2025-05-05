package task

import (
	"fmt"
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
