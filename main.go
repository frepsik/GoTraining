package main

import (
	"encoding/json"
	"errors"
	"fmt"
	payment "goTraining/Payment"
	"net/http"
	"strconv"
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
	var payment payment.PayInfo
	//Здесь расшифровываем из json в структуру, производится десериализация данных
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("Fail read HTTP request (JSON):", err)
	}

	if err := payService(payment); err != nil {
		fmt.Println("Fail pay service:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Если чуть более грамотно разделять, тут в теории можно было бы сделать и rwMutex, на считывание
	mu.Lock()
	historyPay = append(historyPay, payment)
	payment.Println()
	fmt.Println(historyPay)
	w.Write([]byte("Оплата успешно произведена. Текущий баланс: " + strconv.Itoa(money)))
	mu.Unlock()

}

func main() {
	http.HandleFunc("/payService", payHandler)

	fmt.Println("Стартуем сервер на приём запросов!")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Fail start HTTP server:", err)
	}
}
