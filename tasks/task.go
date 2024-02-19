package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func ListTasks(tasks []Task) {

	if len(tasks) == 0 {
		fmt.Println("You have no tasks to complete!")
		return
	}

	for _, task := range tasks {
		if task.Completed {
			fmt.Println(task.ID, "[✓]", task.Title)
		} else {
			fmt.Println(task.ID, "[ ]", task.Title)
		}
	}
}

func AddTask(tasks []Task, title, description string) []Task {
	newTask := Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		Completed:   false,
	}

	return append(tasks, newTask)
}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)

	if err != nil {
		panic(err)
	}

	writer.Flush()

	if err != nil {
		panic(err)
	}
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			break
		}
	}
	return tasks
}
