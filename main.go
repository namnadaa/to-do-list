package main

import (
	"fmt"
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
	// list := []Task{
	// 	{Task: "do smthng", Completed: true},
	// 	{Task: "let code", Completed: false},
	// 	{Task: "sleep", Completed: true},
	// }
	// data, err := json.MarshalIndent(list, "", "  ")
	// if err != nil {
	// 	fmt.Println("Ошибка сериализации:", err)
	// 	return
	// }
	// fmt.Println(string(data))

	addTask("do smthng")
	addTask("sleep")
	addTask("coding")
	showList()
}
