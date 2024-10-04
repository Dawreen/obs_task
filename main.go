package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"obsidian_tasks/googletasks"
	"obsidian_tasks/markdowntasks"

	"google.golang.org/api/tasks/v1"
)

func main() {
	path := "."
	taskListTitle := "Obsidian Tasks"
	// taskListTitle := "Test obsidian_tasks"

	fmt.Println("Start synch " + taskListTitle)

	taskListId, err := googletasks.GetTasksListId(taskListTitle)
	if err != nil {
		log.Fatalf(`Got error: %v`, err)
	}

	allTasksMd, err := markdowntasks.GetAllTasksMdPath(path)
	if err != nil {
		log.Fatalf(`Got error: %v`, err)
	}
	fmt.Printf("Found %v tasks on Obsidian\n\n", len(allTasksMd))

	allTasksGoogle := googletasks.GetAllTasksGoogle(taskListTitle)
	fmt.Printf("Found %v tasks on GoogleTasks\n\n", len(allTasksGoogle))

	mdIdMap := make(map[string]string)

	i := 1
	for key, value := range allTasksGoogle {
		fmt.Printf("\rOn checking: %d/%d", i, len(allTasksGoogle))
		i++
		if value.Status == "completed" && !allTasksMd[key].Status {
			// update markdown
			err = markdowntasks.DoneTaskMd(value.Notes, value.Title)
			if err != nil {
				continue
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
			if value.Parent == "" {
				mdIdMap[key] = value.Id
			}
			delete(allTasksMd, key)
		}
	}

	i = 0
	fmt.Println("\nAdding new tasks!")
	for key, value := range allTasksMd {
		fmt.Printf("\rAdding: %d/%d", i, len(allTasksMd))
		i++
		if !value.Status {
			taskGoogle := tasks.Task{
				Title: value.Title,
				Notes: value.Path,
			}
			newTaskGoogle, err := googletasks.AddTaskGoogle(taskListId, &taskGoogle)
			if err != nil {
				log.Fatalf(`Failed to create task! error: %v`, err)
			}
			mdIdMap[key] = newTaskGoogle.Id
		}
	}

	i = 0
	fmt.Println("\nSetting parent")
	for key, value := range mdIdMap {
		fmt.Printf("\rConnecting: %d/%d", i, len(mdIdMap))
		i++
		pathTask := strings.Split(key, "|")
		taskTitle := pathTask[1]
		titleMd := filepath.Base(strings.TrimSuffix(pathTask[0], ".md"))
		if taskTitle != titleMd {
			// is child
			keyParent := pathTask[0] + "|" + filepath.Base(strings.TrimSuffix(pathTask[0], ".md"))
			_, err := googletasks.SetParentGoogle(taskListId, value, mdIdMap[keyParent])
			if err != nil {
				log.Fatalf(`Failed to move task! error: %v
					%v
					%v
					%v
					%v`, err, taskListId, value, mdIdMap[keyParent], keyParent)
			}
		}
	}
	fmt.Println("\nSynch completed!")
}
