/*
Copyright © 2024 TheOreoTM
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TodoList struct {
	Tasks []Task
}

func (t *TodoList) LoadTodos(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &t.Tasks)

	if err != nil {
		return err
	}

	return nil
}

func (t *TodoList) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := json.MarshalIndent(t.Tasks, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(byteValue)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoList) GetPath() string {
	return "todolist.json"
}

func (t *TodoList) CompleteTask(i int) error {
	if i < 0 || i >= len(t.Tasks) {
		return fmt.Errorf("Index out of bounds")
	}

	t.Tasks[i].Complete()
	t.Tasks[i].UpdatedAt = time.Now()

	err := t.Save(t.GetPath())
	if err != nil {
		return err
	}
	return nil
}

func (t *TodoList) GetTask(i int) (Task, error) {
	if i < 0 || i >= len(t.Tasks) {
		return Task{}, fmt.Errorf("Index out of bounds")
	}

	return t.Tasks[i], nil
}

func (t *TodoList) GetTasks() []Task {
	return t.Tasks
}

func (t *TodoList) AddTask(s string) error {
	task := Task{
		Text:      s,
		Completed: false,
		Index:     len(t.Tasks),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Tasks = append(t.Tasks, task)
	err := t.Save(t.GetPath())
	if err != nil {
		return err
	}
	return nil
}

func NewTodoList() *TodoList {
	t := &TodoList{}
	err := t.LoadTodos(t.GetPath())
	if err != nil {
		fmt.Printf("Error loading todos: %s\n", err)
		return &TodoList{}
	}
	return t
}
