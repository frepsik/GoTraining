package taskmodule

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	id             uuid.UUID
	head           string
	description    string
	isCompleted    bool
	dateCreated    time.Time
	dateCompleated time.Time
}

func NewTask(head string, description string) *Task {
	return &Task{
		id:          uuid.New(),
		head:        head,
		description: description,
		isCompleted: false,
		dateCreated: time.Now(),
	}
}
