package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/davidtom/tarea/pkg/task"
)

var (
	addDescShort = "creates a new task"
	addDescLong  = "creates a new task"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add <task>",
	Short: addDescShort,
	Long:  addDescLong,
	Args:  cobra.MinimumNArgs(1),
	Run:   addTask,
}

func addTask(cmd *cobra.Command, args []string) {
	// TODO: pass in all flags as second arg
	t := task.NewTask(args[0], priority)

	fmt.Printf("%+v", *t)
}
