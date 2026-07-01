package service

import (
	repo "goTraining/Repo"
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
		return taskmodule.Task{}, err
	}

	return newTask, nil
}

func (ts *TaskService) GetTaskbyId(uuid.UUID) {

}

func (ts *TaskService) GetTasks() {

}

// Пока не знаю, как сюда передавать query параметры, под вопросом
func (ts *TaskService) GetCompleatedTasks() {

}

// Пока не знаю, как сюда передавать query параметры, под вопросом
func (ts *TaskService) GetUnCompleatedTasks() {

}

// Пока не знаю, как сюда передавать json файл, в теории полями, если я изменяю только статус
func (ts *TaskService) PatchTaskCompleated(idTask uuid.UUID, taskStatus bool) {

}

func (ts *TaskService) DeleteTask(idTask uuid.UUID) {

}
