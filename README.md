
This is my submition to the 
https://roadmap.sh/projects/task-tracker challenge.


# Task CLI

A simple command-line application for managing tasks. This project enables users to add, update, delete, and list tasks, as well as change their statuses, while persisting data in a JSON file.

---

## Features

- Add new tasks with a description.
- Update task descriptions.
- Delete tasks by their ID.
- Mark tasks as "in progress" or "done."
- List all tasks or filter them by status (`todo`, `done`, or `in-progress`).
- Store tasks in a JSON file for persistence.

---

## Requirements

- The application should run from the command line.
- Task data should be stored in a JSON file in the current directory.
- Use positional arguments for commands and inputs.
- Handle errors and edge cases gracefully (e.g., invalid inputs, file access errors).
- No external libraries are used for this project.

---

## Usage

### Adding a Task
```bash
go run main.go add "Your task description"
# Output: Task added successfully (ID: 1)

# Using `go run` directly (without building the binary)
go run main.go add "Your task description"
```

### Listing Tasks
```bash
# List all tasks
tgo run main.golist

# List tasks by status
go run main.go list todo
go run main.go list done
go run main.go list in-progress

# Using `go run` directly
go run main.go list
go run main.go list todo
go run main.go list done
go run main.go list in-progress
```

### Updating a Task
```bash
go run main.go update <task_id> "Updated task description"
# Example
go run main.go update 1 "Buy groceries and cook dinner"

# Using `go run` directly
go run main.go update <task_id> "Updated task description"
```

### Deleting a Task
```bash
go run main.go delete <task_id>
# Example
go run main.go delete 1

# Using `go run` directly
go run main.go delete <task_id>
```

### Marking a Task
```bash
# Mark a task as in-progress
go run main.go mark-in-progress <task_id>

# Mark a task as done
go run main.go mark-done <task_id>

# Using `go run` directly
go run main.go mark-in-progress <task_id>
go run main.go mark-done <task_id>
```

---

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/your-repo/task-cli.git
   cd task-cli
   ```

2. Build the application:
   ```bash
   go build -o task-cli main.go
   ```

3. Add the binary to your system's PATH for easy access (optional).

---

## JSON File Storage

- The application uses a `tasks.json` file in the current directory to store tasks.
- If the file does not exist, it will be created automatically.
- The JSON file format:
  ```json
  [
    {
      "id": 1,
      "description": "Task description",
      "status": "todo",
      "createdAt": "2025-01-04T15:04:05Z",
      "updatedAt": "2025-01-04T15:04:05Z"
    }
  ]
  ```

---
