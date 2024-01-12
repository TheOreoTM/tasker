/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

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

		filePath := "todolist.json"
		task := args[0]

		fmt.Println(filePath)
		fmt.Println(task)

		// Read the JSON file
		file, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Unmarshal the JSON data into a map
		var data map[string]interface{}
		err = json.Unmarshal(file, &data)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Add the task to the list
		tasks, ok := data["tasks"].([]interface{})
		if !ok {
			tasks = []interface{}{}
		}
		tasks = append(tasks, task)
		data["tasks"] = tasks

		// Marshal the updated data back to JSON
		updatedData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}

		os.WriteFile(filePath, updatedData, 0644)

		// Write the updated data back to the file
		// Write the updated data back to the file
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s added to your todo list: %s\n", task, args)
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)

	addTaskCmd.Flags().StringP("list", "l", "main", "The list to add the task to")
}
