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

func (js *JsonStorage) Load() (map[uuid.UUID]taskmodule.Task, error) {
	file, err := os.Open(js.path)
	if err != nil {
		//Проверка на то, что ошибка связана с тем, что файла не существует
		if errors.Is(err, os.ErrNotExist) {
			return make(map[uuid.UUID]taskmodule.Task), nil
		}
		//Если любая другая ошибка при попытке открыть файл
		return nil, fmt.Errorf("%w: open file: %w", ErrStorage, err)
	}
	defer file.Close()

	var loadValues []taskmodule.Task

	//Возвращаем также произвольную ошибку, если возникла проблема при десериализации json файла
	if err := json.NewDecoder(file).Decode(&loadValues); err != nil {
		return nil, fmt.Errorf("%w: decode json: %w", ErrStorage, err)
	}

	tasksMap := make(map[uuid.UUID]taskmodule.Task, len(loadValues))

	for _, value := range loadValues {
		tasksMap[value.Id] = value
	}

	return tasksMap, nil
}

func (js *JsonStorage) Save(tasks map[uuid.UUID]taskmodule.Task) error {
	file, err := os.Create(js.path)
	if err != nil {
		return fmt.Errorf("fail open file: %w %w", err, ErrStorage)
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
