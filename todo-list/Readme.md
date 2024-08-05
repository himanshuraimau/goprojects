Certainly! Here's a comprehensive `README.md` for your Todo App project:

---

# Todo App

Welcome to the Todo App project! This CLI-based application allows you to manage tasks directly from your terminal. You can perform basic CRUD (Create, Read, Update, Delete) operations on tasks, and these tasks are stored in a CSV file.

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
- **Timediff**: For displaying relative friendly time differences (e.g., 1 hour ago, 10 minutes ago).
- **OS**: For file operations.
- **Syscall**: For file locking to prevent concurrent read/writes.

## Project Structure

```plaintext
todo-app/
├── cmd/
│   └── root.go
│   └── add.go
│   └── list.go
│   └── complete.go
│   └── delete.go
├── internal/
│   ├── tasks/
│   │   ├── task.go
│   │   ├── repository.go
│   │   └── util.go
├── data/
│   └── tasks.csv
├── go.mod
└── go.sum
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

## Technical Considerations

### Stderr vs Stdout

Make sure to write any diagnostics or errors to the stderr stream and write output to stdout.

### File Locking

The underlying data file should be locked by the process to prevent concurrent read/writes. This can be achieved using the `flock` system call in Unix-like systems to obtain an exclusive lock on the file.

Example code for file locking:

```go
func loadFile(filepath string) (*os.File, error) {
    f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
    if (err != nil) {
        return nil, fmt.Errorf("failed to open file for reading: %w", err)
    }

    // Exclusive lock obtained on the file descriptor
    if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
        _ = f.Close()
        return nil, err
    }

    return f, nil
}

func closeFile(f *os.File) error {
    syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
    return f.Close()
}
```

## Contribution

Contributions to this project are welcome! If you have any ideas for improvements or would like to contribute code, feel free to submit a pull request.

---

This `README.md` provides a comprehensive guide to setting up, using, and understanding the structure and functionality of your Todo App.