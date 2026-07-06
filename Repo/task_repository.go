package repo

import (
	storage "goTraining/Storage"
	taskmodule "goTraining/TaskModule"
	"sync"

	"github.com/google/uuid"
)

// Repository - это прослойка, где осуществляется работа с данными внутри системы, БД к примеру под капотом использует отдельные механизмы как грамотно данные сохранять при различных операциях
// в связи с этим, эту прослойку не выносят в отдельный файл и не вызывают в service, в моём случае здесь нужно будет реализовать это в service. Здесь просто сама работа с данными, с небольшими проверками,
// не входные валидации - это уровень Serivce

type TaskRepository struct {
	tasks   map[uuid.UUID]taskmodule.Task
	storage *storage.JsonStorage
	mtx     sync.RWMutex
}

func NewTaskRepository(tasks map[uuid.UUID]taskmodule.Task, storage *storage.JsonStorage) *TaskRepository {
	return &TaskRepository{
		tasks:   tasks,
		storage: storage,
	}
}

// Добавление задачи
func (tr *TaskRepository) Add(task taskmodule.Task) error {

	tr.mtx.Lock()
	defer tr.mtx.Unlock()

	//Если таска уже существует по такмоу uuid
	if _, ok := tr.tasks[task.Id]; ok {
		return ErrTaskAlreadyExists
	}

	tr.tasks[task.Id] = task

	if err := tr.storage.Save(tr.tasks); err != nil {
		return err
	}

	return nil
}

// Получение задачи по id
func (tr *TaskRepository) GetById(idTask uuid.UUID) (taskmodule.Task, error) {

	tr.mtx.RLock()
	defer tr.mtx.RUnlock()

	task, ok := tr.tasks[idTask]
	if !ok {
		return taskmodule.Task{}, ErrSearchTaskById
	}
	return task, nil
}

// Получение всех задач
func (tr *TaskRepository) Get() []taskmodule.Task {

	tr.mtx.RLock()
	defer tr.mtx.RUnlock()

	result := []taskmodule.Task{}

	for _, t := range tr.tasks {
		result = append(result, t)
	}

	return result
}

// Получение завершённых задач
func (tr *TaskRepository) GetCompleated() []taskmodule.Task {

	tr.mtx.RLock()
	defer tr.mtx.RUnlock()

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

	tr.mtx.RLock()
	defer tr.mtx.RUnlock()

	result := []taskmodule.Task{}

	for _, t := range tr.tasks {
		if t.IsCompleted != true {
			result = append(result, t)
		}
	}
	return result
}

// Изменить статус на выполнено у текущей задачи
func (tr *TaskRepository) PatchTaskStatus(id uuid.UUID, status bool) (taskmodule.Task, error) {

	tr.mtx.Lock()
	defer tr.mtx.Unlock()

	task, ok := tr.tasks[id]
	if !ok {
		return taskmodule.Task{}, ErrSearchTaskById
	}
	task.IsCompleted = status
	tr.tasks[id] = task

	if err := tr.storage.Save(tr.tasks); err != nil {
		return taskmodule.Task{}, err
	}

	return task, nil
}

// Удалить задачу
func (tr *TaskRepository) Delete(id uuid.UUID) error {

	tr.mtx.Lock()
	defer tr.mtx.Unlock()

	if _, ok := tr.tasks[id]; !ok {
		return ErrSearchTaskById
	}

	delete(tr.tasks, id)

	if err := tr.storage.Save(tr.tasks); err != nil {
		return err
	}

	return nil
}
