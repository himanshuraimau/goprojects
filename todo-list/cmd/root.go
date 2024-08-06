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

