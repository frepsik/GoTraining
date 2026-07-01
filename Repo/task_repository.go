package repo

import (
	"errors"
	storage "goTraining/Storage"
	taskmodule "goTraining/TaskModule"

	"github.com/google/uuid"
)

// Repository - это прослойка, где осуществляется работа с данными внутри системы, БД к примеру под капотом использует отдельные механизмы как грамотно данные сохранять при различных операциях
// в связи с этим, эту прослойку не выносят в отдельный файл и не вызывают в service, в моём случае здесь нужно будет реализовать это в service. Здесь просто сама работа с данными, с небольшими проверками,
// не входные валидации - это уровень Serivce

type TaskRepository struct {
	tasks   map[uuid.UUID]taskmodule.Task
	storage *storage.JsonStorage
}

func NewTaskRepository(storage *storage.JsonStorage) *TaskRepository {
	return &TaskRepository{
		tasks:   make(map[uuid.UUID]taskmodule.Task),
		storage: storage,
	}
}

// Добавление задачи
func (tr *TaskRepository) Add(task taskmodule.Task) error {
	tr.tasks[task.Id] = task

	if err := tr.storage.Save(tr.tasks); err != nil {
		return err
	}
	return nil
}

// Получение задачи по id
func (tr *TaskRepository) GetById(idTask uuid.UUID) (taskmodule.Task, error) {
	task, ok := tr.tasks[idTask]
	if !ok {
		return taskmodule.Task{}, errors.New("Task not found")
	}
	return task, nil
}

// Получение всех задач
func (tr *TaskRepository) Get() []taskmodule.Task {
	result := []taskmodule.Task{}

	for _, t := range tr.tasks {
		result = append(result, t)
	}

	return result
}

// Получение завершённых задач
func (tr *TaskRepository) GetCompleated() []taskmodule.Task {
	result := []taskmodule.Task{}
	for _, t := range tr.tasks {
		if t.IsCompleted == true {
			result = append(result, t)
		}
	}
	return result
}

// Получение не завершённых задач
func (tr *TaskRepository) GetNotCompleated() []taskmodule.Task {
	result := []taskmodule.Task{}

	for _, t := range tr.tasks {
		if t.IsCompleted != true {
			result = append(result, t)
		}
	}
	return result
}

// Изменить статус на выполнено у текущей задачи
func (tr *TaskRepository) PatchTaskCompleated(id uuid.UUID) error {
	task, ok := tr.tasks[id]
	if !ok {
		return errors.New("Task not found")
	}
	task.IsCompleted = true
	tr.tasks[id] = task
	return nil
}

// Удалить задачу
func (tr *TaskRepository) Delete(id uuid.UUID) error {
	if _, ok := tr.tasks[id]; !ok {
		return errors.New("Task not found")
	}

	delete(tr.tasks, id)

	return nil
}
