# To-Do List CLI

A feature-rich command-line To-Do list application written in Go.
This project demonstrates modular architecture, colorized terminal output, autosaving, undo functionality, and export features.

---

## Features

- Add, edit, and delete tasks
- Mark and unmark tasks as completed
- Batch toggle all tasks
- Autosave with toggleable state
- Undo the last action (add, edit, delete, toggle, sort)
- Sort tasks by completion
- Export task list to a plain text file
- Save the task list to a custom file (Save As...)
- Colored and aligned terminal output
- Visual progress bar
- Modular structure (packages: task, storage, toggle, history, show, filemenu, color)

---

## How It Works

The application runs in the terminal and displays an interactive text-based menu.
Users interact with the app by selecting options and typing input via the keyboard.

---

## Menus Overview

### Main Menu
- Add task  
- Show tasks  
- Toggle menu (mark/unmark)  
- Delete task  
- Edit task  
- Undo last action  
- File menu  
- Exit  

### Toggle Menu
- Mark or unmark a single task  
- Mark or unmark all tasks  
- Return to main menu  

### File Menu
- Toggle autosave on/off  
- Save task list to custom file (Save As...)  
- Export task list to `.txt`  
- Return to main menu  

---

## Export Example

When exporting, tasks are formatted like this:

```
Status  Task
[x]     Buy bread
[ ]     Call a friend
[ ]     Write code
```

---

## Undo Support

Undo is available for:
- Task creation  
- Task deletion  
- Task editing  
- Task completion toggling (including mass toggle)  
- Sorting  

Each undo reverts only the last action.

---

## Data Storage

Tasks are stored in `tasks.json` in JSON format.  
On startup, tasks are automatically loaded.  
Autosave is enabled by default and can be toggled.

---

## Progress Bar

Displays the percentage of completed tasks:

```
#   Status  Task
1   [x]     Buy bread
2   [ ]     Call a friend
3   [ ]     Write code

Progress:
[###-------] 33.3%  (1/3)
```

---

## Running the App

```bash
go run main.go
```

---

## Project Structure

```
todolist/
├── cmd/
│   └── main.go           # Entry point      
├── pkg/
│   ├── color/            # ANSI color helpers      
│   ├── filemenu/         # File menu options
│   ├── history/          # Undo functionality  
│   ├── show/             # Display list and progress 
│   ├── storage/          # Autosave, save/load, export  
│   ├── task/             # Task logic and model  
│   ├── toggle/           # Toggle task completion                
├── go.mod                # Module definition  
├── README.md
└── tasks.json            # Task data (ignored by Git)  
```
