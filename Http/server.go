package http

import (
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
	router.Path("/tasks/{idTask}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTaskById)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTasks)
	router.Path("/tasks").Methods("GET").Queries("isCompleted", "true").HandlerFunc(s.httpHandlers.HandleGetCompleatedTasks)
	router.Path("/tasks").Methods("GET").Queries("isCompleted", "false").HandlerFunc(s.httpHandlers.HandleGetUnCompleatedTasks)
	router.Path("/tasks/{idTask}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandlePatchTaskCompleated)
	router.Path("/tasks/{idTask}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)

	//Возвращаем либо ошибку запуска, сервера, либо nil, если сервер успешно запустился и сидит слушает
	return http.ListenAndServe(":9091", router)
}
