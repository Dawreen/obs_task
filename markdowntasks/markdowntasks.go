package markdowntasks

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type MdTask struct {
	Title  string
	Id     string
	Status bool
}

func getAllTasksMdPath(rootPath string) ([]MdTask, error) {
	var allTasks []MdTask

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".md") {
			fileTasks, err := getAllTasksMd(path)
			if err != nil {
				return err
			}
			allTasks = append(allTasks, fileTasks...)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return allTasks, nil
}

func getAllTasksMd(filePath string) ([]MdTask, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var allTasks []MdTask
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "- [ ]") {
			task := strings.TrimSpace(strings.TrimPrefix(line, "- [ ]"))
			allTasks = append(allTasks,
				MdTask{
					task,
					filePath + "|" + task,
					false,
				})
		}
		if strings.Contains(line, "- [X]") {
			task := strings.TrimSpace(strings.TrimPrefix(line, "- [X]"))
			allTasks = append(allTasks,
				MdTask{
					task,
					filePath + "|" + task,
					true,
				})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allTasks, nil
}

func DoneTaskMD(input string) (string, error) {
	return "That", nil
}
