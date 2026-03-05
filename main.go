package main

import (
	"fmt"
	"goTraining/payments"
	"goTraining/payments/methods"
)

func main() {
	card := &methods.Card{}
	paypal := &methods.PayPal{}
	crypto := &methods.Crypto{}

	//Создаем map, с целью того, чтобы сразу хранить в нём всех провайдеров, чтобы был доступ к конкретному при работе модуля оплаты
	providersMap := map[payments.PaymentMethodType]payments.PaymentMethod{
		payments.CardMethod:   card,
		payments.PayPalMethod: paypal,
		payments.CryptoMethod: crypto,
	}

	paymentModule := payments.NewPaymentModule(providersMap)

	cardPay := payments.CardMethod

	payment := paymentModule.Pay(cardPay, 1200, "God of War")

	paymentModule.Info(payment)

	paymentOut := paymentModule.Refund(cardPay, payment)

	fmt.Println(paymentOut.IsCancelled, paymentOut.MethodPay)

}
