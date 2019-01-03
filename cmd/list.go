package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/davidtom/tarea/pkg/storage"
)

// TODO: make all descriptions more descriptive
var (
	listDescShort = "lists tasks"
	listDescLong  = "list tasks"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   addDescShort,
	Long:    addDescLong,
	Args:    cobra.NoArgs,
	Run:     listTask,
}

func listTask(cmd *cobra.Command, args []string) {
	// TODO: eventually I should pass flags in here (ie --project personal)
	// there are global flags set already too though!
	tasks := storage.Get()

	// TODO: pass tasks to a printer package method
	fmt.Println(tasks)
}
