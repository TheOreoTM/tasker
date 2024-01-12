/*
Copyright © 2024 TheOreoTM
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as complete",
	Long: `Mark a task as complete. For example:
	
	tasker task complete 1`,

	Run: func(cmd *cobra.Command, args []string) {
		todoList := NewTodoList()
		indexInt, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("Invalid task number.")
			return
		}

		task, err := todoList.GetTask(indexInt)

		if err != nil {
			fmt.Println("Task not found.")
			return
		}

		todoList.CompleteTask(indexInt)

		fmt.Printf("Task completed successfully!\n%s\n", task.String())
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
