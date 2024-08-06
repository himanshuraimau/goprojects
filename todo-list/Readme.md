
# Todo List
Welcome to the Todo List project! This CLI-based application allows you to manage tasks directly from your terminal. You can perform basic CRUD (Create, Read, Update, Delete) operations on tasks, and these tasks are stored in a CSV file.

## Overview

This project involves creating a CLI application that handles the following responsibilities:

1. **Add Task**: Create a new task with a given description.
2. **List Tasks**: Display all uncompleted tasks, with an option to list all tasks.
3. **Complete Task**: Mark a task as completed.
4. **Delete Task**: Remove a task from the list.

## Requirements

The application supports the following CLI commands:

### Add

The `add` command creates a new task in the underlying data store. It takes a positional argument with the task description.

```
$ tasks add <description>
```

For example:

```
$ tasks add "Tidy my desk"
```

### List

The `list` command returns a list of all uncompleted tasks. Use the `-a` or `--all` flag to list all tasks regardless of their completion status.

```
$ tasks list
$ tasks list -a
```

### Complete

The `complete` command marks a task as done.

```
$ tasks complete <taskid>
```

### Delete

The `delete` command removes a task from the data store.

```
$ tasks delete <taskid>
```

## Example CLI Commands

```bash
# Add a new task
$ tasks add "Write documentation for project"

# List all uncompleted tasks
$ tasks list

# List all tasks including completed ones
$ tasks list -a

# Mark a task as complete
$ tasks complete 1

# Delete a task
$ tasks delete 1
```

## Technologies Used

- **CSV**: For storing tasks in a simple CSV file.
- **Cobra**: For creating the command-line interface.
- **OS**: For file operations.
- **Syscall**: For file locking to prevent concurrent read/writes.

## Project Structure

```plaintext
todo-app/
├── cmd/
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── complete.go
│   └── delete.go
├── internal/
│   ├── tasks/
│   │   ├── task.go
│   │   ├── repository.go
│   │   └── util.go
├── data/
│   └── tasks.csv
├── go.mod
├── go.sum
└── main.go
```

### Description of Files and Directories

- **cmd/**: Contains the CLI commands.
  - `root.go`: Sets up the root command and subcommands.
  - `add.go`: Command for adding a new task.
  - `list.go`: Command for listing tasks.
  - `complete.go`: Command for marking a task as complete.
  - `delete.go`: Command for deleting a task.
  
- **internal/tasks/**: Contains core business logic and data handling.
  - `task.go`: Defines the Task struct and related methods.
  - `repository.go`: Handles reading from and writing to the CSV file.
  - `util.go`: Contains utility functions, such as file locking.

- **data/**: Stores the data file for tasks.
  - `tasks.csv`: The CSV file where tasks are stored.

- **go.mod**: Defines the module and its dependencies.
- **go.sum**: Contains checksums for dependencies.
- **main.go**: The entry point for the application, sets up and executes the root command.

## Getting Started

To get started with this project, follow these steps:

1. **Clone the repository**:

```sh
git clone https://github.com/himanshuraimau/goprojects/todo-list
cd todo-list
```

2. **Initialize Go modules**:

```sh
go mod init github.com/himanshuraimau/goprojects/todo-list
go mod tidy
```

3. **Run the CLI application**:

```sh
go run main.go
```

## Example Data File

An example CSV file looks like the following:

```plaintext
ID,Description,CreatedAt,IsComplete
1,My new task,2024-07-27T16:45:19-05:00,true
2,Finish this video,2024-07-27T16:45:26-05:00,true
3,Find a video editor,2024-07-27T16:45:31-05:00,false
```
