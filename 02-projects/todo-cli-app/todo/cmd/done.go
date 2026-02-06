package cmd

import (
	"fmt"
	"strconv"

	"github.com/r-koma/go-learning-journal/02-projects/todo-cli-app/todo/internal/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneCmd = &cobra.Command{
	Use:   "done [ID]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %s", args[0])
		}

		file := viper.GetString("datafile")
		m, err := todo.NewManager(file)
		if err != nil {
			return err
		}

		if err := m.MarkAsDone(id); err != nil {
			return err
		}
		fmt.Printf("Task %d marked as done!\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
