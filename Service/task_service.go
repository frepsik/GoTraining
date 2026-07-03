package service

import (
	"errors"
	repo "goTraining/Repo"
	storage "goTraining/Storage"
	taskmodule "goTraining/TaskModule"

	"github.com/google/uuid"
)

type TaskService struct {
	taskRepository *repo.TaskRepository
}

func NewTaskService(taskRepository *repo.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (ts *TaskService) CreateTask(head string, description string) (taskmodule.Task, error) {
	newTask := taskmodule.NewTask(head, description)

	if err := ts.taskRepository.Add(newTask); err != nil {

		//В случае если произошла ошибка открытия файла, нужно сокрыть это на этом уровне
		if errors.Is(err, storage.ErrStorage) {
			return taskmodule.Task{}, ErrInternalServer
		}
		return taskmodule.Task{}, err
	}

	return newTask, nil
}

func (ts *TaskService) GetTaskbyId(taskId uuid.UUID) (taskmodule.Task, error) {
	task, err := ts.taskRepository.GetById(taskId)
	if err != nil {
		return taskmodule.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) GetTasks() []taskmodule.Task {
	return ts.taskRepository.Get()
}

// Не надо никак принимать query параметры, потому что если сработал необходиый handler, значит параметры таковы и были
func (ts *TaskService) GetCompleatedTasks() []taskmodule.Task {
	return ts.taskRepository.GetCompleated()
}

// Не надо никак принимать query параметры, потому что если сработал необходиый handler, значит параметры таковы и были
func (ts *TaskService) GetUnCompleatedTasks() []taskmodule.Task {
	return ts.GetUnCompleatedTasks()
}

// Пока не знаю, как сюда передавать json файл, в теории полями, если я изменяю только статус
func (ts *TaskService) PatchTaskCompleated(taskId uuid.UUID, taskStatus bool) (taskmodule.Task, error) {
	pathcTask, err := ts.taskRepository.PatchTaskStatus(taskId, taskStatus)
	if err != nil {
		if errors.Is(err, storage.ErrStorage) {
			return taskmodule.Task{}, ErrInternalServer
		}

		return taskmodule.Task{}, err
	}
	return pathcTask, nil
}

func (ts *TaskService) DeleteTask(idTask uuid.UUID) error {
	return ts.DeleteTask(idTask)
}
