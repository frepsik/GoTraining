package handlerforpayandsavemoney

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Ранее я тут передавал атомики, но потом, когда перешёл к мьютексам в критической секции, пересмотрел логику, и
// если делать через структуру, данную запись с return func(w http.ResponseWriter, r *http.Request) можно не использовать
// можно сразу сделать через хэндлер, но в качестве примера, тут оставил
func (b *Balance) PayServiceHandler() http.HandlerFunc {
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

		b.mu.Lock()
		if b.Money-price >= 0 {
			//Инсценирую вариант, когда могла бы быть гонка данных, но посредствам mutex, мы её исключаем
			time.Sleep(3 * time.Second)
			b.Money -= price
			fmt.Println("-", price)
			fmt.Println("Текущий баланс:", b.Money)

			//Говорят лучше это использовать s := strconv.FormatInt(100, 10), чуток попозже надо будет глянуть, подумать
			w.Write([]byte("Баланс на карте: " + strconv.FormatInt(int64(b.Money), 10)))
		} else {
			fmt.Printf("На балансе не хватает средств")
			w.Write([]byte("Мало мани братух"))
		}
		b.mu.Unlock()
	}

}

func (b *Balance) SaveMoneyHandler() http.HandlerFunc {
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

		b.mu.Lock()
		if sum <= b.Money {
			b.Money -= sum
			b.Bank += sum

			fmt.Println("Текущая сумма в банке:", b.Bank)
			fmt.Println("На кармане остало:", b.Money)
			w.Write([]byte("В банке сейчас: " + strconv.FormatInt(int64(b.Bank), 10)))
		} else {
			fmt.Println("Не хватает бабок закинуть в банк")
			w.Write([]byte("Не хватает бабок закинуть в банк"))
		}
		b.mu.Unlock()
	}

}
