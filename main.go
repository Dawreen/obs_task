package main

import (
	"fmt"
	"log"

	"obsidian_tasks/googletasks"
	"obsidian_tasks/markdowntasks"

	"google.golang.org/api/tasks/v1"
)

func main() {
	path := "."
	// taskListTitle := "Obsidian Tasks"
	taskListTitle := "Test obsidian_tasks"
	taskListId, err := googletasks.GetTasksListId(taskListTitle)
	if err != nil {
		log.Fatalf(`Got error: %v`, err)
	}

	fmt.Println("Starting search for tasks in markdown files!")
	allTasksMd, err := markdowntasks.GetAllTasksMdPath(path)
	if err != nil {
		log.Fatalf(`Got error: %v`, err)
	}
	fmt.Printf("%v tasks found\n\n", len(allTasksMd))

	fmt.Println("Search on GoogleTasks!")
	allTasksGoogle := googletasks.GetAllTasksGoogle(taskListTitle)
	fmt.Printf("%v tasks found\n\n", len(allTasksGoogle))

	fmt.Println("Update all tasks!")

	for key, value := range allTasksGoogle {
		if value.Status == "completed" && !allTasksMd[key].Status {
			fmt.Println("Completed task found")
			// update markdown
			err = markdowntasks.DoneTaskMd(value.Notes, value.Title)
			if err != nil {
				log.Fatalf(`Got error: %v`, err)
			}
			delete(allTasksMd, key)
		}
		if value.Status == "needsAction" && allTasksMd[key].Status {
			value.Status = "completed"
			_, err = googletasks.DoneTaskGoogle(taskListId, value.Id, &value)
			if err != nil {
				log.Fatalf(`No update on GoogleTasks! error: %v`, err)
			}
			delete(allTasksMd, key)
		}
		if value.Status == "needsAction" && !allTasksMd[key].Status {
			delete(allTasksMd, key)
		}
	}

	fmt.Println("Add new tasks to GoogleTasks!")

	for _, value := range allTasksMd {
		if !value.Status {
			taskGoogle := tasks.Task{
				Title: value.Title,
				Notes: value.Path,
			}
			_, err = googletasks.AddTaskGoogle(taskListId, &taskGoogle)
			if err != nil {
				log.Fatalf(`Failed to create task! error: %v`, err)
			}
		}
		// TODO gestion parent
	}
}
