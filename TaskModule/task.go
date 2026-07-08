package taskmodule

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID  `json:"id"`
	Head        string     `json:"head"`
	Description string     `json:"description"`
	IsCompleted bool       `json:"isCompleted"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt"`
}

func NewTask(head string, description string) Task {
	return Task{
		Id:          uuid.New(),
		Head:        head,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}
