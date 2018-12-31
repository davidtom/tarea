package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// TODO: do either gopkg.* files belong in .gitignore?
// TODO: best way to add vendor to .gitignore?

// TODO: which flags should be persistent (global) and which local?
var (
	// define persistent/global flags
	priority string
)

func init() {
	// TODO: make sure my usage strings follow conventions (see blog post)
	// TODO: make default priority configurable via config
	rootCmd.PersistentFlags().StringVarP(&priority, "priority", "p", "medium", "task priority level [default: medium")
}

var rootCmd = &cobra.Command{
	Use:   "tarea",
	Short: "tarea is a command line todo list that goes where you do",
	Long: `tarea is a command line todo list that goes where you do.
Built with Go`,
}

// Execute - runs all commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
