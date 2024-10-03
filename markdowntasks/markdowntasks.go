package markdowntasks

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/maps"
)

type MdTask struct {
	Title  string
	Status bool
}

func getAllTasksMdPath(rootPath string) (map[string]MdTask, error) {
	allTasksMap := make(map[string]MdTask)

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".md") {
			fileTasksMap, err := getAllTasksMd(path)
			if err != nil {
				return err
			}
			maps.Copy(allTasksMap, fileTasksMap)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return allTasksMap, nil
}

func getAllTasksMd(filePath string) (map[string]MdTask, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	allTasksMap := make(map[string]MdTask)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "- [ ]") {
			task := strings.TrimSpace(strings.TrimPrefix(line, "- [ ]"))
			allTasksMap[filePath+"|"+task] = MdTask{task, false}
		}
		if strings.Contains(line, "- [X]") {
			task := strings.TrimSpace(strings.TrimPrefix(line, "- [X]"))
			allTasksMap[filePath+"|"+task] = MdTask{task, true}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allTasksMap, nil
}

func DoneTaskMd(filePath string, task string) error {
	// Apri il file in modalità lettura e scrittura
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crea un buffer per leggere e scrivere le righe del file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "- [ ] "+task) {
			// Trovato il task, aggiorna lo stato
			line = strings.Replace(line, "- [ ]", "- [X]", 1)
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Scrivi le righe modificate nel file
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
