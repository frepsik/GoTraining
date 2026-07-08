package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	httpHandlers *HttpHandlers
}

func NewHttpServer(httpHandlers *HttpHandlers) *HttpServer {
	return &HttpServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HttpServer) StartServer() error {
	//Расписал более корректные маршруты, в рамках работы с запросами HTTP, без необходимости использовать if, else
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)
	router.Path("/tasks/{taskId}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTaskById)
	router.Path("/tasks").Methods("GET").Queries("isCompleated", "true").HandlerFunc(s.httpHandlers.HandleGetCompleatedTasks)
	router.Path("/tasks").Methods("GET").Queries("isCompleated", "false").HandlerFunc(s.httpHandlers.HandleGetUnCompleatedTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTasks)
	router.Path("/tasks/{taskId}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandlePatchTaskCompleated)
	router.Path("/tasks/{taskId}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)

	//Возвращаем либо ошибку запуска, сервера, либо nil, если сервер успешно запустился и сидит слушает

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
