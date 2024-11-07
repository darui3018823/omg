package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Description string
	Completed   bool
}

var tasks []Task

func addTask(description string) {
	tasks = append(tasks, Task{Description: description, Completed: false})
	fmt.Println("Task added:", description)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("To-Do List:")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "[X]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
	}
}

func completeTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index].Completed = true
	fmt.Println("Task completed:", tasks[index].Description)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Add Task\n2. List Tasks\n3. Complete Task\n4. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := scanner.Text()
			addTask(description)
		case "2":
			listTasks()
		case "3":
			fmt.Print("Enter task number to complete: ")
			scanner.Scan()
			var index int
			fmt.Sscanf(scanner.Text(), "%d", &index)
			completeTask(index - 1) // Adjusting for zero-based index
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
