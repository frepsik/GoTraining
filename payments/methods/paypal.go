package methods

import "fmt"

type PayPal struct{}

func (c PayPal) Pay(usd int, description string) int {
	fmt.Println("Оплата PayPal")
	fmt.Printf("Размер оплаты: %d", usd)
	return 1
}

func (c PayPal) Refund(id int) int {
	fmt.Println("Возврат средств, на PayPal пользователя")
	return 1
}
