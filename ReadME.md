# CLI To-Do List in Go

A simple command-line **To-Do List** built using **Go**, storing tasks in a JSON file.

## 📌 Features
- Add tasks  
- List tasks  
- Mark tasks as done  
- Delete tasks  
- Persistent storage using `tasks.json`  

## 🚀 Installation & Usage

### 1️⃣ Clone the Repository
    
    git clone https://github.com/your-username/cli-to-do-list.git

## 2️⃣ Run the Program

Since the project is split into multiple files, **you must include all Go files when running**:

    
    go run main.go storage.go task.go <command> [arguments]

## 📜 Commands & Usage

| Command             | Description                           | Example Usage                              |
|---------------------|---------------------------------------|--------------------------------------------|
| `add <task>`       | Add a new task                        | `go run *.go add "Finish Go project"`     |
| `list`             | Show all tasks                        | `go run *.go list`                        |
| `done <task_id>`   | Mark a task as completed              | `go run *.go done 1`                      |
| `delete <task_id>` | Delete a task                         | `go run *.go delete 1`                    |
## 🛠️ How It Works

1. Tasks are stored in `tasks.json` for persistence.  
2. `main.go` handles user input and executes the corresponding functions.  
3. `task.go` defines the `Task` struct and its methods (like marking a task as done).  
4. `storage.go` manages reading from and writing to `tasks.json`.  
5. Running the program with `go run main.go storage.go task.go <command>` executes the desired task operation.  
