package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/davidtom/tarea/pkg/storage"
)

// TODO: make all descriptions more descriptive
var (
	doneDescShort = "marks a task as complete"
	doneDescLong  = "marks a task as complete"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: doneDescShort,
	Long:  doneDescLong,
	Args:  cobra.MinimumNArgs(1),
	Run:   doneTask,
}

func doneTask(cmd *cobra.Command, args []string) {
	// TODO: need to map this all out better! storage.Get() right now is also what reassigns/resaves ids
	// I should sit down and plan out all of these methods better!
	// Whatever method I use here CANNOT update ids
	// I also need to think about how this all works when Im listing tasks from specific projects - how are ids
	// updated and assigned in that case???

	tasks := storage.Get()

	// TODO: pass tasks to a printer package method
	fmt.Println(tasks)
}
