package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/Piero7210/my-go-cli/tasks"
)

func main() {
	// OpenFile is the generalized open call; most users will use Open or Create instead.
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	// Close closes the File, rendering it unusable for I/O.
	defer file.Close()

	var tasks []task.Task

	// Stat returns the FileInfo structure describing file.
	// It returns an error if the file does not exist or if there is an error accessing it.
	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	// Size returns the size in bytes for regular files; system-dependent for others.
	if info.Size() != 0 {
		// ReadAll reads from r until an error or EOF and returns the data it read.
		bytes, err := io.ReadAll(file)

		if err != nil {
			panic(err)
		}
		// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
		err = json.Unmarshal(bytes, &tasks)

		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)

	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the title of the task: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		fmt.Println("Enter the description of the task: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		tasks := task.AddTask(tasks, title, description)
		task.SaveTask(file, tasks)

	case "complete":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the ID of the task you want to delete: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)
		// Convert string to int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		tasks = task.CompleteTask(tasks, idInt)
		task.SaveTask(file, tasks)

	case "delete":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the ID of the task you want to delete: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)
		// Convert string to int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		tasks = task.DeleteTask(tasks, idInt)
		task.SaveTask(file, tasks)

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Use one of the following commands: [list, add, complete, delete]")
}
