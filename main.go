package main

import (
	"fmt"
	handlerforpayandsavemoney "goTraining/HandlerForPayAndSaveMoney"
	"net/http"
)

//Здесь рассматриваю небольшую логику работы http обработчиков и запуск сервера с обработчиками.
//В данной программе рассмотрены фукнции обработчики, на возможные запросы с клиента, к примеру: оплата услуги, отмена оплаты услуги.
//Какая логика, у нас есть http.HandelFunc, этим мы говорим, что у нас есть точка, которая ожидает запрос определённого паттерна (например /pay) и может
//обработать его определённым образом, посредствам своего Handler, что мы указываем в качестве второго аргумента функции (например payHandler).
//Если углубиться в логику работы Handler, у него есть первый аргумент, что отвечает к примеру, за то, чтобы переслать определённый ответ клиенту в качесве []byte (определённый сериализованный поток данных),
//опять же таки здесь и должна располагаться определённая обработка запроса (к примеру проведение оплаты. Здесь условный пример, логики оплаты нет).

//Вернувшись к базовой логике, сначала, мы грубо говоря, сначала указываем точки маршрутизации, какие запросы сюда могут дойти (http.HandelFunc).
//Далее нам необходимо запустить сам сервер, который будет ожидать принимать запросы. Делается это посредствам: http.ListenAndServe(":9091", nil) - где первый аргумент,
//это порт на котором запускаем, и адрес, а второй, необходим для более тонкой настройки обработчиков

//Когда приходит запрос от клиента, сервер создаёт как раз таки горутину и начинает искать нужный EndPoint, и по нему вызывает handler

// Пример обработчика на запрос оплаты определённой услуги
func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата произведена успешно!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка", err.Error())
	} else {
		fmt.Println("Оплата была корректно проведена")
	}
}

// Пример обработчика на запрос отмены оплаты определённый услуги
func cancelPayHandler(w http.ResponseWriter, r *http.Request) {
	str := "Отмена успешно произведена!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка", err.Error())
	} else {
		fmt.Println("Отмена оплаты успешно произведена")
	}
}

// Пример обработчика
func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello world"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP произошла ошибка", err.Error())
	} else {
		fmt.Println("Я корректно обработал HTTP Запрос")
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	str := "Вы попали на корневой endpoint"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP произошла ошибка", err.Error())
	} else {
		fmt.Println("Корректно попали на корневой endpoint")
	}
}

func main() {
	balance := &handlerforpayandsavemoney.Balance{
		Bank:  0,
		Money: 50,
	}

	//Функция, где мы говорим, что можем ожидать определённый сценарий и запускать на него обработчик, просто создаём маршрутизацию, которой могут в последствии воспользоваться
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelPayHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/payHandler", balance.PayServiceHandler())
	http.HandleFunc("/saveHandler", balance.SaveMoneyHandler())

	//Запускаем сам сервер, на ожидание прихода определённых запросов, на 9091 порту, второй аргумент используется под более тонкую настройку handler, ещё не вникал в это
	fmt.Println("Запускаем сервер на приём запрсоов!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Сервер не удалось успешно запустить", err.Error())
	}

	fmt.Println("Программа закончила своё выполнение!")
}
