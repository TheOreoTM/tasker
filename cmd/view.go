/*
Copyright © 2024 TheOreoTM
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View your todo list",
	Long: `View your todo list. For example:

	tasker task view`,

	Run: func(cmd *cobra.Command, args []string) {
		// open the json file and print the list

		todolist := NewTodoList()
		tasks := todolist.GetTasks()

		if len(tasks) == 0 {
			fmt.Println("You have no tasks in your list.")
			return
		}

		for _, task := range tasks {
			fmt.Printf(task.Format())
		}

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
