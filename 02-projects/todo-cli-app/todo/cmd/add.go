package cmd

import (
	"fmt"

	"github.com/r-koma/go-learning-journal/02-projects/todo-cli-app/todo/internal/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := viper.GetString("datafile")

		m, err := todo.NewManager(file)
		if err != nil {
			return err
		}

		task := args[0]
		if err := m.Add(task); err != nil {
			return err
		}

		fmt.Printf("Added: \"%s\"\n", task)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
