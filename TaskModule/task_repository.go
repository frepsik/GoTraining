package taskmodule

import (
	"errors"

	"github.com/google/uuid"
)

// Repository - это прослойка, где осуществляется работа с данными внутри системы, БД к примеру под капотом использует отдельные механизмы как грамотно данные сохранять при различных операциях
// в связи с этим, эту прослойку не выносят в отдельный файл и не вызывают в service, в моём случае здесь нужно будет реализовать это в service. Здесь просто сама работа с данными, с небольшими проверками,
// не входные валидации - это уровень Serivce

type TaskRepository struct {
	tasks map[uuid.UUID]Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[uuid.UUID]Task),
	}
}

// Добавление задачи
func (tr *TaskRepository) Add(task Task) error {
	tr.tasks[task.Id] = task
	return nil
}

// Получение задачи по id
func (tr *TaskRepository) GetById(idTask uuid.UUID) (Task, error) {
	task, ok := tr.tasks[idTask]
	if !ok {
		return Task{}, errors.New("Task not found")
	}
	return task, nil
}

// Получение всех задач
func (tr *TaskRepository) Get() []Task {
	result := []Task{}

	for _, t := range tr.tasks {
		result = append(result, t)
	}

	return result
}

// Получение завершённых задач
func (tr *TaskRepository) GetCompleated() []Task {
	result := []Task{}
	for _, t := range tr.tasks {
		if t.isCompleted == true {
			result = append(result, t)
		}
	}
	return result
}

// Получение не завершённых задач
func (tr *TaskRepository) GetNotCompleated() []Task {
	result := []Task{}

	for _, t := range tr.tasks {
		if t.isCompleted != true {
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
	task.isCompleted = true
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
