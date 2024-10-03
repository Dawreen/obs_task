package main

import (
	"fmt"
	"obsidian_tasks/googletasks"
	"obsidian_tasks/markdowntasks"
)

func main() {
	fmt.Println(markdowntasks.GetAllTasksMD("Hello"))
	fmt.Println(markdowntasks.DoneTaskMD("input string"))

	fmt.Println(googletasks.GetAllTasksGoogle("Hello"))
	fmt.Println(googletasks.DoneTaskGoogle("input string"))
}
