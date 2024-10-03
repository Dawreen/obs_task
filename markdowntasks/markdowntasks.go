package markdowntasks

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func getAllTasksMDPath(rootPath string) ([]string, error) {
	var allTasks []string

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".md") {
			fileTasks, err := getAllTasksMD(path)
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

func getAllTasksMD(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var allTasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "- [ ]") {
			task := strings.TrimSpace(strings.TrimPrefix(line, "- [ ]"))
			allTasks = append(allTasks, task)
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
