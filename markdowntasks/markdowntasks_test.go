package markdowntasks

import (
	"slices"
	"testing"
)

func TestGetAllTasksMdPath(t *testing.T) {
	want_tasks := []MdTask{
		{"task 1", "files/file_tasks.md|task 1", false},
		{"task 2 in header 2", "files/file_tasks.md|task 2 in header 2", false},
		{"Task done", "files/file_tasks.md|Task done", true},
		{"Task todo in TODO header", "files/file_tasks.md|Task todo in TODO header", false},
	}
	input := "files"
	msg, err := getAllTasksMdPath(input)
	if err != nil {
		t.Fatalf(`Got error: %v`, err)
	}
	for _, task := range want_tasks {
		if !slices.Contains(msg, task) {
			t.Fatalf(`%v is not present in returned tasks (= %v)`, task, msg)
		}
	}
}

func TestDoneTaskMD(t *testing.T) {
	wantBefore := MdTask{
		"task 1",
		"files/file_tasks.md|task 1",
		false,
	}
	wantAfter := MdTask{
		"task 1",
		"files/file_tasks.md|task 1",
		true,
	}
	path := "files/file_tasks.md"
	taskTitle := "task 1"

	allTasks, err := getAllTasksMdPath(".")
	if !slices.Contains(allTasks, wantBefore) {
		return // task 1 is already done (upper test will fail)
	}
	err = DoneTaskMD(path, taskTitle)
	if err != nil {
		t.Fatalf(`Got error: %v`, err)
	}

	allTasks, err = getAllTasksMdPath(".")

	if err == nil && !slices.Contains(allTasks, wantAfter) {
		t.Fatalf(`Task %v was NOT updated to %v, error: %v`, wantBefore, wantAfter, err)
	} else if err == nil && slices.Contains(allTasks, wantAfter) {
		return
	} else {
		t.Fatalf(`No update was performed`)
	}
}
