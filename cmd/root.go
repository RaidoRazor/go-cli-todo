package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var taskID int

const filename = "tasks.csv"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli-todo",
	Short: "A brief description of your application",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func saveTasks() {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"ID", "Title", "Completed"})
	if err != nil {
		return
	}

	for _, task := range tasks {
		err := writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Title,
			strconv.FormatBool(task.Completed),
		})
		if err != nil {
			return
		}
	}
}

func loadTasks() {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, record := range records[1:] {
		id, _ := strconv.Atoi(record[0])
		completed := record[2] == "true"
		tasks = append(tasks, Task{ID: id, Title: record[1], Completed: completed})
		if id >= taskID {
			taskID = id + 1
		}
	}
}

func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Printf("Deleted task with ID %d\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

func listTasks() {
	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "[X]"
		}
		fmt.Printf("%d: %s %s\n", task.ID, status, task.Title)
	}
}

func completeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			saveTasks()
			fmt.Printf("Marked task %d as completed.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	loadTasks()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
