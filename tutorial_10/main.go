package main

import "fmt"

func main() {
    tasks := []string{}
    // Map to track if task is completed
    status := make(map[string]bool)

    fmt.Println("--- My Go Task Manager ---")

    for {
        var choice int
        fmt.Println("\n1. Add Task  2. List Tasks  3. Exit")
        fmt.Scan(&choice)

        if choice == 3 {
            break
        }

        switch choice {
        case 1:
            var name string
            fmt.Print("Enter task name: ")
            fmt.Scan(&name)
            tasks = append(tasks, name)
            status[name] = false
        case 2:
            fmt.Println("\nYour Tasks:")
            for i, task := range tasks {
                done := "❌"
                if status[task] {
                    done = "✅"
                }
                fmt.Printf("%d. %s [%s]\n", i+1, task, done)
            }
        default:
            fmt.Println("Invalid choice!")
        }
    }
}