package main

import (
	"goTraining/panics"
)

// "fmt"
// "goTraining/payments"
// "goTraining/payments/methods"

func main() {
	// 1) ИНТЕРФЕЙСЫ
	//Делаем тут так, потому что, нужно будет передавать в интерфейс, а он там проверяет при приёмке, в Method set, есть ли там методы нужные, а если туда закинуть
	//просто экземпляр, сразу value, будет проблема с тем, что он не увидит наши методы с типом указателя на структуру, которые работают напрямую с экземпляром структуры
	//если перадать в интерфейс, value структуру, в method set, не будут видны методы типа указатель, если передать Pointer, Интерфейс при проверке Method set, увидит
	//сразу все методы и Pointer и Value типа, у меня тут в реализации методы типа pointer (пример: func (c *Card) Refund(id int) int)
	// card := &methods.Card{}
	// paypal := &methods.PayPal{}
	// crypto := &methods.Crypto{}

	// //Создаем map, с целью того, чтобы сразу хранить в нём всех провайдеров, чтобы был доступ к конкретному при работе модуля оплаты
	// providersMap := map[payments.PaymentMethodType]payments.PaymentMethod{
	// 	payments.CardMethod:   card,
	// 	payments.PayPalMethod: paypal,
	// 	payments.CryptoMethod: crypto,
	// }

	// paymentModule := payments.NewPaymentModule(providersMap)

	// cardPay := payments.CardMethod

	// payment := paymentModule.Pay(cardPay, 1200, "God of War")

	// paymentModule.Info(payment)

	// paymentOut := paymentModule.Refund(cardPay, payment)

	// fmt.Println(paymentOut.IsCancelled, paymentOut.MethodPay)

	// 2)ИСКЛЮЧЕНИЯ
	// db := &exceptiontestprogramm.DataBase{}

	// u1 := &entities.User{
	// 	Id:     123,
	// 	Name:   "Лёха",
	// 	Age:    29,
	// 	Email:  "Почта",
	// 	Number: "nomer trubki",
	// }
	// u2 := &entities.User{}

	// fmt.Println(*u1)

	// err := db.AddUser(u1)
	// err2 := db.AddUser(u2)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	//3 ПАНИКИ
	panics.TestFunc()
}
