package markdowntasks

import (
	"regexp"
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
	input := "This"
	want := regexp.MustCompile(`\bThat\b`)
	msg, err := DoneTaskMD(input)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`This = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
