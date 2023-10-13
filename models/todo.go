package models

import "fmt"

type Todo struct {
	Description string
	Done        bool
}

func (t Todo) String() string {
	return fmt.Sprintf("%s  - %t", t.Description, t.Done)
}

func NewTodo(description string) Todo {
	return Todo{Description: description, Done: false}
}
