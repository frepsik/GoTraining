package handlerforpayandsavemoney

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
)

func PayServiceHandler(money *atomic.Int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			errorMessage := "Fail read http body" + err.Error()
			fmt.Println(errorMessage)

			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				fmt.Println("Fail write message")
				return
			}
			return
		}

		httpRequestBodyString := string(httpRequestBody)

		price, err := strconv.Atoi(httpRequestBodyString)
		if err != nil {
			errorMessage := "Fail convert []byte to int" + err.Error()
			fmt.Println(errorMessage)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				fmt.Println("Fail write message")
				return
			}
			return
		}

		if money.Load()-int64(price) >= 0 {

			money.Add(int64(-price))
			fmt.Println("-", price)
			fmt.Println("Текущий баланс:", money.Load())

			//Говорят лучше это использовать s := strconv.FormatInt(100, 10), чуток попозже надо будет глянуть, подумать
			w.Write([]byte("Баланс на карте: " + strconv.FormatInt(money.Load(), 10)))
		} else {
			fmt.Printf("На балансе не хватает средств")
			w.Write([]byte("Мало мани братух"))
		}
	}

}

func SaveMoneyHandler(bank *atomic.Int64, money *atomic.Int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpRequestBody, err := io.ReadAll(r.Body)
		if err != nil {
			errorMessage := "Fail read http body" + err.Error()
			fmt.Println(errorMessage)

			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				fmt.Println("Fail write message")
				return
			}
			return
		}

		httpRequestBodyString := string(httpRequestBody)

		sum, err := strconv.Atoi(httpRequestBodyString)
		if err != nil {
			errorMessage := "Fail convert []byte to int" + err.Error()
			fmt.Println(errorMessage)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				fmt.Println("Fail write message")
				return
			}
			return
		}

		if sum <= int(money.Load()) {
			money.Add(int64(-sum))

			bank.Add(int64(sum))

			fmt.Println("Текущая сумма в банке:", bank.Load())
			fmt.Println("На кармане остало:", money.Load())
			w.Write([]byte("В банке сейчас: " + strconv.FormatInt(bank.Load(), 10)))
		} else {
			fmt.Println("Не хватает бабок закинуть в банк")
			w.Write([]byte("Не хватает бабок закинуть в банк"))
		}
	}

}
