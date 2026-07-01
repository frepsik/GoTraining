package taskmodule

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id            uuid.UUID
	head          string
	description   string
	IsCompleted   bool
	dateCreated   time.Time
	dateCompleate time.Time
}

func NewTask(head string, description string) Task {
	return Task{
		Id:          uuid.New(),
		head:        head,
		description: description,
		IsCompleted: false,
		dateCreated: time.Now(),
	}
}
