package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	tasks, _ := loadTasks() // Ensure loadTasks() is correctly defined in storage.go

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <add/list/done/delete> [task]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Provide a task description")
			return
		}
		task := Task{ID: len(tasks) + 1, Text: os.Args[2]} // Ensure Task struct exists in tasks.go
		tasks = append(tasks, task)
		saveTasks(tasks) // Ensure saveTasks() is correctly defined in storage.go
		fmt.Println("Task added!")

	case "list":
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, task := range tasks {
			fmt.Printf("[%d] %s (Done: %v)\n", task.ID, task.Text, task.Done)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Provide a task ID to mark as done")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil || id < 1 || id > len(tasks) {
			fmt.Println("Invalid task ID")
			return
		}
		tasks[id-1].MarkDone()
		saveTasks(tasks)
		fmt.Println("Task marked as done!")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Provide a task ID to delete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil || id < 1 || id > len(tasks) {
			fmt.Println("Invalid task ID")
			return
		}
		tasks = append(tasks[:id-1], tasks[id:]...)
		saveTasks(tasks)
		fmt.Println("Task deleted!")

	default:
		fmt.Println("Unknown command")
	}
}
