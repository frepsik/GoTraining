package main

import (
	"encoding/json"
	"errors"
	"fmt"
	payment "goTraining/Payment"
	"net/http"
	"sync"
)

var mu sync.Mutex
var historyPay = make([]payment.PayInfo, 0)
var money int = 1000

func payService(payment payment.PayInfo) error {
	var errMsg error

	if payment.USD <= money {
		money -= payment.USD
	} else {
		errMsg = errors.New("Недостаточно средств для проведения оплаты")
	}

	return errMsg
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	var currentPayment payment.PayInfo
	//Здесь расшифровываем из json в структуру, производится десериализация данных
	if err := json.NewDecoder(r.Body).Decode(&currentPayment); err != nil {
		fmt.Println("Fail read HTTP request (JSON):", err)
	}

	if err := payService(currentPayment); err != nil {
		fmt.Println("Fail pay service:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Если чуть более грамотно разделять, тут в теории можно было бы сделать и rwMutex, на считывание
	mu.Lock()
	historyPay = append(historyPay, currentPayment)
	httpResponse := payment.HttpResponse{
		PaymentHistory: historyPay,
		Money:          money,
	}
	currentPayment.Println()
	fmt.Println(historyPay)
	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("Fail convert to json httpResponce:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
	mu.Unlock()

}

func main() {
	http.HandleFunc("/payService", payHandler)

	fmt.Println("Стартуем сервер на приём запросов!")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Fail start HTTP server:", err)
	}
}
