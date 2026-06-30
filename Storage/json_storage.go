package storage

import (
	"encoding/json"
	taskmodule "goTraining/TaskModule"
	"os"

	"github.com/google/uuid"
)

type JsonStorage struct {
	path string
}

func NewJsonStorage(path string) *JsonStorage {

	return &JsonStorage{
		path: path,
	}
}

func (js *JsonStorage) Load() (error, map[uuid.UUID]taskmodule.Task) {
	file, err := os.Open(js.path)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	var loadValues []taskmodule.Task
	if err := json.NewDecoder(file).Decode(&loadValues); err != nil {
		return err, nil
	}

	tasksMap := make(map[uuid.UUID]taskmodule.Task, len(loadValues))

	for _, value := range loadValues {
		tasksMap[value.Id] = value
	}

	return nil, tasksMap
}

func (js *JsonStorage) Save(tasks map[uuid.UUID]taskmodule.Task) error {
	file, err := os.Create(js.path)
	if err != nil {
		return err
	}
	defer file.Close()

	values := make([]taskmodule.Task, 0, len(tasks))

	for _, value := range tasks {
		values = append(values, value)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(values)
}
