package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// loadTasks reads tasks from tasks.json
func loadTasks() ([]Task, error) {
	file, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // Return empty if file doesn't exist
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// saveTasks writes tasks to tasks.json
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("tasks.json", data, 0644)
}