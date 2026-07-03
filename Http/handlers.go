package http

import (
	"encoding/json"
	"errors"
	"fmt"
	repo "goTraining/Repo"
	service "goTraining/Service"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//Файл, со всеми endPoint (паттернами), где будут существовать различные виды обработки запросов, как прослойка, из которой вызывают другие методы, по сути можно было бы сделать
//что отсюда я буду вызывать прослойку service, где будет происходить валидация некоторая и работа, а далее уже вызываться repository, но пока в планах упрощённую версию сделать

type HttpHandlers struct {
	taskService *service.TaskService
}

func NewHttpHandlers(taskService *service.TaskService) *HttpHandlers {
	return &HttpHandlers{
		taskService: taskService,
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

	//Пробуем десерелизовать, то что прислал клиент
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		//Отправляем обратно, с error, потмоу что прислал данные некорректные
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}

	//Проверяем на адекватность отправленные значения
	if err := taskDTO.ValidationForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		//Отправляем обратно, с error, потмоу что прислал данные некорректные
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}

	//Вызываем service, который под капотом создаст task, в последствии отправит на запись в repository
	task, err := h.taskService.CreateTask(taskDTO.Head, taskDTO.Description)

	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		//Custom error
		if errors.Is(err, service.ErrInternalServer) {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}

	bytes, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("fail to write http response:", err)
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
	taskIdString := mux.Vars(r)["taskId"]
	taskid, err := uuid.Parse(taskIdString)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTaskbyId(taskid)

	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, repo.ErrSearchTaskById) {
			//404 статус
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
			return
		} else {
			//500 статус
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}

	//Превращаем в массив байт определённого формата, а именно json
	bytes, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}

	//Указываем статус 200
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("fail to write http response:", err)
	}
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
	tasks := h.taskService.GetTasks()
	bytes, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("fail to write http response:", err)
	}
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
	// Не надо никак принимать query параметры, потому что если сработал необходиый handler, значит параметры таковы и были
	compleatedTasks := h.taskService.GetCompleatedTasks()
	bytes, err := json.MarshalIndent(compleatedTasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("fail to write http response:", err)
	}
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
	uncompleatedTasks := h.taskService.GetUnCompleatedTasks()
	bytes, err := json.MarshalIndent(uncompleatedTasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("fail to write http response:", err)
	}
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
	var statusTask = PatchTaskDTO{}

	//Первичная валидация входящий данных
	if err := json.NewDecoder(r.Body).Decode(&statusTask); err != nil {
		errDTO := NewErrorDTO(err.Error())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}
	taskIdString := mux.Vars(r)["taskId"]
	taskId, err := uuid.Parse(taskIdString)
	if err != nil {
		errDTO := NewErrorDTO(err.Error())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}

	patchTask, err := h.taskService.PatchTaskCompleated(taskId, statusTask.status)
	if err != nil {
		errDTO := NewErrorDTO(err.Error())

		if errors.Is(err, repo.ErrSearchTaskById) {
			//404 статус
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
			return
		} else if errors.Is(err, service.ErrInternalServer) {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return

		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}

	bytes, err := json.MarshalIndent(patchTask, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		fmt.Println("fail to write http response:", err)
	}

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
	taskIdString := mux.Vars(r)["taskId"]
	taskId, err := uuid.Parse(taskIdString)
	if err != nil {
		errDTO := NewErrorDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := h.taskService.DeleteTask(taskId); err != nil {
		errDTO := NewErrorDTO(err.Error())
		if errors.Is(err, repo.ErrSearchTaskById) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}

}
