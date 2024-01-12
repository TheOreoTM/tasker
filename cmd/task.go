package cmd

import (
	"fmt"
	"time"
)

type Task struct {
	Completed bool      `json:"completed"`
	Text      string    `json:"text"`
	Index     int       `json:"index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t *Task) Uncomplete() {
	t.Completed = false
}

func (t *Task) UpdateText(s string) {
	t.Text = s
}

func (t *Task) String() string {
	return fmt.Sprintf("%s\n", t.Text)
}

func (t *Task) Format() string {
	task := t.String()

	paddedIndex := fmt.Sprintf("%03d", t.Index)

	task = fmt.Sprintf("%s. %s", paddedIndex, task)

	if t.Completed {
		task = "[x] " + task
	} else {
		task = "[ ] " + task
	}

	return task
}
