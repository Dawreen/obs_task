package googletasks

import (
	"regexp"
	"testing"
)

func TestGetAllTasksGoogle(t *testing.T) {
	// taskListTitle := "Test obsidian_tasks"
	taskTitle1 := "Task with subs"
	taskTitle2 := "sub_task1"
	taskTitle3 := "sub_task2"

	allTasksMap := GetAllTasksGoogle()

	_, ok1 := allTasksMap["|"+taskTitle1]
	_, ok2 := allTasksMap["|"+taskTitle2]
	_, ok3 := allTasksMap["dettaglio2|"+taskTitle3]

	if !ok1 {
		t.Fatalf(`%v missing in %v`, taskTitle1, allTasksMap)
	}
	if !ok2 {
		t.Fatalf(`%v missing in %v`, taskTitle2, allTasksMap)
	}
	if !ok3 {
		t.Fatalf(`%v missing in %v`, taskTitle3, allTasksMap)
	}
}

func TestDoneTaskMD(t *testing.T) {
	input := "This"
	want := regexp.MustCompile(`\bThat\b`)
	msg, err := DoneTaskGoogle(input)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`This = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
