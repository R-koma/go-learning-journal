package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/r-koma/go-learning-journal/02-projects/todo-cli-app/todo/internal/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		file := viper.GetString("datafile")
		m, err := todo.NewManager(file)
		if err != nil {
			return err
		}

		items := m.GetAll()
		if len(items) == 0 {
			fmt.Println("No tasks found")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tTask\tStatus\tCreated")

		all, _ := cmd.Flags().GetBool("all")

		for _, item := range items {
			if !all && item.Done {
				continue
			}

			status := "[ ]"
			if item.Done {
				status = "[x]"
			}
			data := item.CreatedAt.Format("2006-01-02 15:04")

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", item.ID, item.Task, status, data)
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "Show all tasks including completing")
}
