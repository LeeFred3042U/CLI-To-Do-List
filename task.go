package main

import "encoding/json"

// Task represents a single to-do task
type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// MarkDone marks a task as completed
func (t *Task) MarkDone() {
	t.Done = true
}

// ToJSON converts a Task to JSON format (optional)
func (t Task) ToJSON() string {
	data, _ := json.MarshalIndent(t, "", "  ")
	return string(data)
}
