package markdowntasks

import (
	"regexp"
	"testing"
)

func TestGetAllTasksMD(t *testing.T) {
	input := "This"
	want := regexp.MustCompile(`\bThat\b`)
	msg, err := GetAllTasksMD(input)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`This = %q, %v, want match for %#q, nil`, msg, err, want)
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
