package taskmodule

import "github.com/google/uuid"

type Task struct {
	id          uuid.UUID
	head        string
	description string
	isCompleted bool
}

func NewTask(head string, description string) *Task {
	return &Task{
		id:          uuid.New(),
		head:        head,
		description: description,
		isCompleted: false,
	}
}

func GetTask(id uuid.UUID) Task {}
