package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/davidtom/tarea/pkg/storage"
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
	d := strings.Join(args, " ")

	// TODO: pass in all flags as second arg
	t := task.NewTask(d, priority)

	storage.Set(t)

	// TODO: add a log statement to indicate it was succesful? Need a printer package first
}
