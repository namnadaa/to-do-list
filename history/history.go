package history

import (
	"fmt"
	"todolist/color"
	"todolist/task"
)

// ActionType defines a type for representing different kinds of user actions.
type ActionType string

const (
	Add    ActionType = "add"
	Sort   ActionType = "sort"
	Toggle ActionType = "toggle"
	Delete ActionType = "delete"
	Edit   ActionType = "edit"
)

// Action represents a user operation that can be undone.
type Action struct {
	Type       ActionType
	TaskData   task.Task
	Index      int
	PrevText   string
	PrevState  []task.Task
	SubActions []Action
}

// History stores a list of actions performed, in order to support undo functionality.
var History []Action

// Record saves the provided action into the history stack.
func Record(action Action) {
	History = append(History, action)
}

// Undo reverts the last recorded action, if possible.
func Undo() {
	if len(History) == 0 {
		fmt.Println(color.Red("Nothing to undo."))
		return
	}

	last := History[len(History)-1]
	History = History[:len(History)-1]

	switch last.Type {
	case Add:
		if len(task.List) > 0 {
			task.List = task.List[:len(task.List)-1]
		} else {
			fmt.Println(color.Magenta("Undo failed: invalid insertion index."))
		}
	case Sort:
		if last.PrevState != nil {
			task.List = last.PrevState
		} else {
			fmt.Println(color.Magenta("Undo failed: invalid insertion index."))
		}
	case Toggle:
		if len(last.SubActions) > 0 {
			for i := len(last.SubActions) - 1; i >= 0; i-- {
				a := last.SubActions[i]
				task.List[a.Index].Completed = !task.List[a.Index].Completed
			}
		} else if last.Index >= 0 && last.Index < len(task.List) {
			task.List[last.Index].Completed = !task.List[last.Index].Completed
		} else {
			fmt.Println(color.Magenta("Undo failed: invalid insertion index."))
		}
	case Delete:
		if last.Index >= 0 && last.Index <= len(task.List) {
			task.List = append(task.List[:last.Index], append([]task.Task{last.TaskData}, task.List[last.Index:]...)...)
		} else {
			fmt.Println(color.Magenta("Undo failed: invalid insertion index."))
		}
	case Edit:
		if last.Index >= 0 && last.Index < len(task.List) {
			task.List[last.Index].Task = last.PrevText
		} else {
			fmt.Println(color.Magenta("Undo failed: invalid insertion index."))
		}
	default:
		fmt.Println(color.Magenta("Undo failed: unknown action type."))
	}
	fmt.Println(color.Green("Undo last action."))
}
