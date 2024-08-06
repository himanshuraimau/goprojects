package cmd

import (
	"github.com/spf13/cobra"
	"encoding/csv"
	"errors"
	"os"
)


// Define the completeCmd
var completeCmd = &cobra.Command{
    Use:   "complete <taskid>",
    Short: "Mark a task as complete",
    Args:  cobra.ExactArgs(1), // Require exactly one argument
    Run: func(cmd *cobra.Command, args []string) {
        taskID := args[0] // Get the task ID from the arguments

        if err := completeTask(taskID); err != nil {
            cmd.PrintErrln("Error completing task:", err)
        } else {
            cmd.Println("Task marked as complete successfully")
        }
    },
}

// add the completeCmd to rootCmd
func init() {
    rootCmd.AddCommand(completeCmd)
}

// completeTask marks a task as complete
func completeTask(taskID string) error {
    // Open the tasks CSV file for reading
    file, err := os.OpenFile("data/tasks.csv", os.O_RDWR, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    rows, err := reader.ReadAll()
    if err != nil {
        return err
    }

    var updatedRows [][]string
    taskFound := false

    // Process rows to find and update the task
    for _, row := range rows {
        if row[0] == taskID {
            // Update the task's IsComplete status
            row[3] = "true"
            taskFound = true
        }
        updatedRows = append(updatedRows, row)
    }

    if !taskFound {
        return errors.New("task not found")
    }

    // Re-open the file for writing and save the updated rows
    file, err = os.OpenFile("data/tasks.csv", os.O_RDWR|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    err = writer.WriteAll(updatedRows)
    if err != nil {
        return err
    }

    return nil
}
