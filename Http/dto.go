package http

import (
	"encoding/json"
	"errors"
	"time"
)

//DTO - data transfer object

type TaskDTO struct {
	Head        string
	Description string
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
	status string
}

func (pt PatchTaskDTO) ValidationForPatch() error {
	if pt.status == "" {
		return errors.New("status is empty")
	} else {
		return nil
	}
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
