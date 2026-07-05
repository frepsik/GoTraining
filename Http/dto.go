package http

import (
	"encoding/json"
	"errors"
	"time"
)

//DTO - data transfer object

type TaskDTO struct {
	Head        string `json:"head"`
	Description string `json:"description"`
}

func (t TaskDTO) ValidationForCreate() error {
	if t.Description == "" {
		return errors.New("Description is empty")
	}
	if t.Head == "" {
		return errors.New("Head is empty")
	}
	return nil
}

type PatchTaskDTO struct {
	status bool
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(message string) ErrorDTO {
	return ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
