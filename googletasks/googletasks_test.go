package googletasks

import (
	"regexp"
	"testing"
)

func TestGetAllTasksMD(t *testing.T) {
	input := "This"
	want := regexp.MustCompile(`\bThat\b`)
	msg, err := GetAllTasksGoogle(input)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`This = %q, %v, want match for %#q, nil`, msg, err, want)
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
