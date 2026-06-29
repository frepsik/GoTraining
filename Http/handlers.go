package http

import (
	"encoding/json"
	taskmodule "goTraining/TaskModule"
	"net/http"
	"time"
)

//Файл, со всеми endPoint (паттернами), где будут существовать различные виды обработки запросов, как прослойка, из которой вызывают другие методы, по сути можно было бы сделать
//что отсюда я буду вызывать прослойку service, где будет происходить валидация некоторая и работа, а далее уже вызываться repository, но пока в планах упрощённую версию сделать

type HttpHandlers struct {
	taskRepository *taskmodule.TaskRepository
}

func NewHttpHandlers(taskRepository *taskmodule.TaskRepository) *HttpHandlers {
	return &HttpHandlers{
		taskRepository: taskRepository,
	}
}

/*
pattern: /tasks
method: Post
info: JSON in Request Body

success:

	-status code: 201 Created
	-response body: JSON with created object

failed:

	-status code: 400, 409, 500...
	-response body: JSON with error message + time
*/
func (h *HttpHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}

	if err := taskDTO.ValidationForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}
}

/*
pattern: /tasks{idTask}
method: Get
info: pattern

success:

	-status code: 200 Ok
	-response body: JSON with get Task

failed:

	-status code: 400, 409, 500...
	-response body: JSON with error message + time
*/
func (h *HttpHandlers) HandleGetTaskById(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks
method: Get
info: -

success:

	-status code: 200 Ok
	-response body: JSON with get Tasks

failed:

	-status code: 400, 500...
	-response body: JSON with error message + time
*/
func (h *HttpHandlers) HandleGetTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks?isCompleted=true
method: Get
info: Query parametrs

success:

	-status code: 200 Ok
	-response body: JSON with get CompleatedTasks

failed:

	-status code: 400, 409, 500...
	-response body: JSON with error message + time
*/
func (h *HttpHandlers) HandleGetCompleatedTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks?isCompleted=false
method: Get
info: Query parametrs

success:

	-status code: 200 Ok
	-response body: JSON with get CompleatedTasks

failed:

	-status code: 400, 409, 500...
	-response body: JSON with error message + time
*/
func (h *HttpHandlers) HandleGetUnCompleatedTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{IdTask}
method: PATCH
info: pattern + JSON with neead param

success:

	-status code: 200 Ok
	-response body: JSON with get CompleatedTasks

failed:

	-status code: 400, 409, 500...
	-response body: JSON with error message + time
*/
func (h *HttpHandlers) HandlePatchTaskCompleated(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{IdTask}
method: DELETE
info: pattern

success:

	-status code: 204 No body
	-response body: -

failed:

	-status code: 400, 409, 500...
	-response body: JSON with error message + time
*/
func (h *HttpHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
