# To-Do List CLI

A simple command-line application for managing a task list, written in Go.  
This project demonstrates modular architecture, submenu support, automatic data saving, and task progress visualization.

---

## Features

- Add, edit, and delete tasks  
- Mark and unmark tasks as completed  
- Bulk actions (mark/unmark all)  
- Submenu for task status control  
- Automatic saving to `tasks.json` after every change  
- Task progress bar  
- Task loading at startup  
- Console-based input

---

## How It Works

The program runs in the terminal and provides a text-based menu.  
The user selects actions, enters data via keyboard, and the app updates the task list while saving changes automatically.

---

## Processing Stages

- **Add Task**  
  Input the task text and add it to the list.

- **Edit Task**  
  Modify the text of an existing task with confirmation.

- **Delete Task**  
  Delete a task by number with confirmation.

- **Toggle Menu (submenu)**  
  Controls the completion status of tasks:
  - Mark a single task
  - Unmark a single task
  - Mark all
  - Unmark all

- **Autosave (`withSave`)**  
  Every change is wrapped in a function that automatically saves the task list to `tasks.json`.

---

## Data Storage

Tasks are serialized in JSON format and saved to a `tasks.json` file in the root directory.  
The data is loaded automatically when the application starts.

---

## Progress Bar

Visually shows the percentage of completed tasks:
```
[x] 1. Buy bread
[ ] 2. Call a friend
[ ] 3. Write code

[##––––] 33.3%  (1/3)
```
---

## Input Source

All input is entered manually via the console.  
All actions are selected through the text menu.

---

## Example Run

```bash
go run main.go
```