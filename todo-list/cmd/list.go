package cmd

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/spf13/cobra"
)

// Define the listCmd
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        showAll, _ := cmd.Flags().GetBool("all")
        if err := listTasks(showAll); err != nil {
            cmd.PrintErrln("Error listing tasks:", err)
        }
    },
}

func init() {
    // Add flags to listCmd
    listCmd.Flags().BoolP("all", "a", false, "Show all tasks including completed ones")
    // Add listCmd to rootCmd
    rootCmd.AddCommand(listCmd)
}

// listTasks displays tasks based on the showAll flag
func listTasks(showAll bool) error {
    // Open the tasks CSV file for reading
    file, err := os.Open("data/tasks.csv")
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    rows, err := reader.ReadAll()
    if err != nil {
        return err
    }

    // Determine the format for output
    format := "%-5s %-45s %-20s %s\n"
    if showAll {
        format = "%-5s %-45s %-20s %-5s\n"
    }

    fmt.Printf(format, "ID", "Task", "Created", "Done")
    for _, row := range rows {
        id := row[0]
        description := row[1]
        createdAt, _ := time.Parse(time.RFC3339, row[2])
        isComplete, _ := strconv.ParseBool(row[3])

        // Display task only if it is not complete or if --all flag is set
        if showAll || !isComplete {
            diff := time.Since(createdAt)
            humanReadableTime := formatDuration(diff)
            done := strconv.FormatBool(isComplete)
            if !showAll {
                done = ""
            }
            fmt.Printf(format, id, description, humanReadableTime, done)
        }
    }

    return nil
}

// formatDuration formats a time.Duration into a human-readable string
func formatDuration(d time.Duration) string {
    if d < time.Minute {
        return fmt.Sprintf("%d seconds ago", int(d.Seconds()))
    } else if d < time.Hour {
        return fmt.Sprintf("%d minutes ago", int(d.Minutes()))
    } else if d < 24*time.Hour {
        return fmt.Sprintf("%d hours ago", int(d.Hours()))
    } else {
        return fmt.Sprintf("%d days ago", int(d.Hours()/24))
    }
}
