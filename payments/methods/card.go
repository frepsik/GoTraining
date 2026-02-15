package methods

import "fmt"

type Card struct {
}

func (c Card) Pay(usd int, description string) int {
	fmt.Println("Оплата картой")
	fmt.Printf("Размер оплаты: %d", usd)
	return 1
}

func (c Card) Cancell(id int) int {
	fmt.Println("Отмена операции оплаты")
	return 1
}
