package payment

import (
	"fmt"
)

type PayInfo struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
}

type HttpResponse struct {
	PaymentHistory []PayInfo `json:"paymentHistory"`
	Money          int       `json:"money"`
}

func (p PayInfo) Println() {
	fmt.Println("Услуга:", p.Description)
	fmt.Println("Стоимость:", p.USD)
	fmt.Println("ФИО:", p.FullName)
	fmt.Println("Адрес", p.Address)
}
