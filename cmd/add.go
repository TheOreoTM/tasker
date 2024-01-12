/*
Copyright © 2024 TheOreoTM
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var List string

// addTaskCmd represents the add command
var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your todo list",
	Long: `Add a task to your todo list. For example: 
	tasker task add "Buy milk"`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		todoList := NewTodoList()

		err := todoList.AddTask(args[0])

		if err != nil {
			panic(err)
		}

		fmt.Printf("Task added successfully!\n %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)

	addTaskCmd.Flags().StringP("list", "l", "main", "The list to add the task to")
}
