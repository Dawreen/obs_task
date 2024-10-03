package markdowntasks

import (
	"testing"
)

func TestGetAllTasksMdPath(t *testing.T) {
	want_tasks := map[string]MdTask{
		"files/file_tasks.md|task 1":                   {"task 1", false},
		"files/file_tasks.md|task 2 in header 2":       {"task 2 in header 2", false},
		"files/file_tasks.md|Task done":                {"Task done", true},
		"files/file_tasks.md|Task todo in TODO header": {"Task todo in TODO header", false},
	}
	input := "files"
	allTasksMdMap, err := getAllTasksMdPath(input)
	if err != nil {
		t.Fatalf(`Got error: %v`, err)
	}
	for key, value := range want_tasks {
		if allTasksMdMap[key] != value {
			t.Fatalf(`%v is not present in returned tasks (= %v)`, value, allTasksMdMap)
		}
	}
}

func TestDoneTaskMD(t *testing.T) {
	wantBefore := MdTask{
		"task 1",
		false,
	}
	wantAfter := MdTask{
		"task 1",
		true,
	}
	path := "files/file_tasks.md"
	taskTitle := "task 1"
	key := path + "|" + taskTitle

	allTasksMap, err := getAllTasksMdPath(".")
	_, ok := allTasksMap[key]
	if !ok {
		return // task 1 does not exist
	}
	err = DoneTaskMd(path, taskTitle)
	if err != nil {
		t.Fatalf(`Got error: %v`, err)
	}

	allTasksMap, err = getAllTasksMdPath(".")
	if err != nil {
		t.Fatalf(`Got error: %v`, err)
	}
	if allTasksMap[key].Status == true {
		return
	}
	if allTasksMap[key].Status == false {
		t.Fatalf(`Task %v was NOT updated to %v, error: %v`, wantBefore, wantAfter, err)
	}
}
