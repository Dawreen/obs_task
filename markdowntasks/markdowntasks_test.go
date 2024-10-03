package markdowntasks

import (
	"regexp"
	"slices"
	"testing"
)

func TestGetAllTasksMD(t *testing.T) {
	want_tasks := []string{
		"task 1",
		"task 2 in header 2",
		"Task done",
		"Task todo in TODO header",
	}
	input := "files"
	msg, err := getAllTasksMDPath(input)
	if err != nil {
		t.Fatalf(`Got error: %v`, err)
	}
	for _, task := range want_tasks {
		if !slices.Contains(msg, task) {
			t.Fatalf(`%q is not present in returned tasks (= %q)`, task, msg)
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
