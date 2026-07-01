package storage

import (
	"encoding/json"
	"errors"
	"fmt"
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
		//Проверка на то, что ошибка связана с тем, что файла не существует
		if errors.Is(err, os.ErrNotExist) {
			return nil, make(map[uuid.UUID]taskmodule.Task)
		}
		//Если любая другая ошибка при попытке открыть файл
		return fmt.Errorf("%w: open file: %w", ErrStorage, err), nil
	}
	defer file.Close()

	var loadValues []taskmodule.Task

	//Возвращаем также произвольную ошибку, если возникла проблема при десериализации json файла
	if err := json.NewDecoder(file).Decode(&loadValues); err != nil {
		return fmt.Errorf("%w: decode json: %w", ErrStorage, err), nil
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
