package main

import (
	"fmt"
	http "goTraining/Http"
	repo "goTraining/Repo"
	service "goTraining/Service"
	storage "goTraining/Storage"
)

func main() {
	//Создаём объект работы с фалйом
	storage := storage.NewJsonStorage("data/tasks.json")
	tasks, err := storage.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Создание репозитория
	taskRepository := repo.NewTaskRepository(tasks, storage)

	//Создание сервиса
	taskService := service.NewTaskService(taskRepository)

	//Создание handlers
	handlers := http.NewHttpHandlers(taskService)

	//Создание http сервера
	httpServer := http.NewHttpServer(handlers)

	//Запуск сервера
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Fail start http server:", err)
		return
	}

}
