package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a title for the task.")
			return
		}
		addTask(args[0])
	},
}

func addTask(title string) {
	newTask := Task{ID: taskID, Title: title, Completed: false}
	tasks = append(tasks, newTask)
	taskID++
	saveTasks()
	fmt.Printf("Added task: %+v\n", newTask)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
