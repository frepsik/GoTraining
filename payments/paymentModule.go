package payments

import (
	"fmt"
	"math/rand"
)

type PaymentMethod interface {
	Pay(usd int, description string) int
	Refund(id int) int
}

type PaymentModule struct {
	paymentMethods map[PaymentMethodType]PaymentMethod
}

// Конструктор - грубо говоря из бд подтягиваем способы оплаты и уже передаём сюда в конструктор в качестве мапы
func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethods: map[PaymentMethodType]PaymentMethod{},
	}
}

func (p *PaymentModule) Pay(paymentMethodType PaymentMethodType, usd int, description string) Payment {
	//Получаем кастомный тип, который я уже создал, проверяем в мапе на его наличие, то есть тип оплаты (карта, paypal, крипта), проверяем, что такое есть
	paymentMethod, ok := p.paymentMethods[paymentMethodType]
	if !ok {
		return Payment{}
	}

	//Производим оплату со стороны самого сервиса, что всё успешно проходит, по сути, от нас протсо запустить эту функцию и всё, она уже где то там, кем то была написана, грубо говоря
	id := paymentMethod.Pay(usd, description)

	//Возвращаем объект в main, для того, чтобы дальше отправить его на запись в бд от туда, или в другую прослойку, и уже от туда начинать записывать в БД
	return Payment{
		rand.Intn(1000),
		usd,
		description,
		false,
		string(paymentMethodType),
		id,
	}
}

// Отмена оплаты, то есть возврат средств
func (p *PaymentModule) Refund(payment Payment) Payment {
	listPayment := map[int]Payment{} //представим что этот мап заполненный уже, просто по мапу искать лучше, чем по слайсу, потому что быстрее

	//Достаём нужный метод оплаты и производим отмену оплаты
	methodPay := p.paymentMethods[PaymentMethodType(payment.MethodPay)]
	methodPay.Refund(payment.idPay)

	paymentGet := listPayment[payment.Id]

	paymentGet.IsCancelled = true

	return paymentGet
}

// Просто вывод чека
func (p *PaymentModule) Info(payment Payment) {
	fmt.Printf("paymentId: %d, paymentDescrtiption: %s", payment.Id, payment.Description)
}
