package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetTodoFile() ([]byte, error) {
	filePath := "todolist.json"

	// Read the JSON file
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("No file found, creating one...")
		// Create file
		_, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}

		return GetTodoFile()
	}

	return file, nil
}

func GetTodos() (Tasks, error) {
	file, err := GetTodoFile()
	if err != nil {
		fmt.Println(err)
		return Tasks{}, err
	}

	// Unmarshal the JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println(err)
		return Tasks{}, err
	}

	// Add the task to the list
	tasks, ok := data["tasks"].([]interface{})
	if !ok {
		tasks = []interface{}{}
	}

	var tasksString []string
	for _, task := range tasks {
		tasksString = append(tasksString, task.(string))
	}

	return Tasks{Tasks: tasksString}, nil

}

type Tasks struct {
	Tasks []string `json:"tasks"`
}
