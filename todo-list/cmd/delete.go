package cmd

import (
    "encoding/csv"
    "errors"
    "os"
    "github.com/spf13/cobra"
)

// Define the deleteCmd
var deleteCmd = &cobra.Command{
    Use:   "delete <taskid>",
    Short: "Delete a task",
    Args:  cobra.ExactArgs(1), // Require exactly one argument
    Run: func(cmd *cobra.Command, args []string) {
        taskID := args[0] // Get the task ID from the arguments

        if err := deleteTask(taskID); err != nil {
            cmd.PrintErrln("Error deleting task:", err)
        } else {
            cmd.Println("Task deleted successfully")
        }
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}

func deleteTask(taskID string) error {
    
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

   // remove the task
    for _, row := range rows {
        if row[0] == taskID {
            taskFound = true
            continue // Skip adding this row to the updatedRows slice
        }
        updatedRows = append(updatedRows, row)
    }

    if !taskFound {
        return errors.New("task not found")
    }
     // repoen the file in write mode
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
