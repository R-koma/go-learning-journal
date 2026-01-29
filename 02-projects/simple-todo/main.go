package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func (t *Task) MarkAsDone(){
	t.Done = true
	fmt.Printf("You did great job: %s\n", t.Title)
}

func (t *Task) Display() {
	status := "[ ]"
	if t.Done {
		status = "[x]" 
	}
	fmt.Printf("%s [%03d] %-20s\n", status, t.ID, t.Title)
}

func main() {
	fmt.Println("start")
	tasks := []*Task{
		{ID: 1, Title: "Study", Done: false},
		{ID: 2, Title: "Eat", Done: false},
		{ID: 3, Title: "Work", Done: false},
	}

	printAllTasks(tasks)
	tasks[0].MarkAsDone()
	printAllTasks(tasks)
}

func printAllTasks(tasks []*Task) {
	for _, t := range tasks {
		t.Display()
	}
}