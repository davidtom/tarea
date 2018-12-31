package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Task does something
// TODO: doc string
// TODO: where do I define this? main? pkg? this doesn't seem like the best spot
type Task struct {
	description string
}

// TODO: do either gopkg.* files belong in .gitignore?

// TODO: which flags should be persistent (global) and which local?
var (
	priority string
)

func init() {
	// TODO: make sure my usage strings follow conventions (see blog post)
	// TODO: make default priority configurable
	rootCmd.PersistentFlags().StringVarP(&priority, "priority", "p", "medium", "task priority level [default: medium")
}

var rootCmd = &cobra.Command{
	Use:   "tarea",
	Short: "tarea is a command line todo list that goes where you do",
	Long: `tarea is a command line todo list that goes where you do.
Built with Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

// Execute - runs all commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
