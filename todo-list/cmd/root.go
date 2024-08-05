package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "tasks",
    Short: "Tasks is a CLI task manager",
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    // Add subcommands here
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(listCmd)
    rootCmd.AddCommand(completeCmd)
    rootCmd.AddCommand(deleteCmd)
}
