package cmd

import (
    "github.com/spf13/cobra"
    "github.com/himanshuraimau/goprojects/todo-list/internal/task" // Ensure this path matches your project structure
)

var addCmd = &cobra.Command{
    Use:   "add <description>",
    Short: "Add a new task",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        description := args[0]
        if err := tasks.AddTask(description); err != nil {
            cmd.PrintErrln("Error adding task:", err)
        } else {
            cmd.Println("Task added successfully")
        }
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
