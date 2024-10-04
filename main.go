package main

import (
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
	taskListId, err := googletasks.GetTasksListId(taskListTitle)
	if err != nil {
		log.Fatalf(`Got error: %v`, err)
	}

	allTasksMd, err := markdowntasks.GetAllTasksMdPath(path)
	if err != nil {
		log.Fatalf(`Got error: %v`, err)
	}

	allTasksGoogle := googletasks.GetAllTasksGoogle(taskListTitle)

	mdIdMap := make(map[string]string)

	for key, value := range allTasksGoogle {
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
			mdIdMap[key] = value.Id
			delete(allTasksMd, key)
		}
	}

	for key, value := range allTasksMd {
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

	for key, value := range mdIdMap {
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
}
