package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- To-Do App ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. Show Tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, strings.TrimSpace(task))
			fmt.Println("Task Added")
		case "2":
			fmt.Println("\n Your Tasks")
			if len(tasks) == 0 {
				fmt.Println("No tasks yet!")
			} else {
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "3":
			fmt.Println("Goodbye")
			return

		default:
			fmt.Println("Invalid choice , try again ")
		}
	}
}
