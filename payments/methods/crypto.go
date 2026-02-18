package methods

import "fmt"

type Crypto struct{}

func (c Crypto) Pay(usd int, description string) int {
	fmt.Println("Оплата криптовалютной")
	fmt.Printf("Размер оплаты: %d", usd)
	return 1
}

func (c Crypto) Refund(id int) int {
	fmt.Println("Возврат средств, на крипто-карту пользователя")
	return 1
}
