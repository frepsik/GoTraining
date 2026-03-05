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

// Конструктор - при создании экземпляра передаём сюда все способы оплаты, их может быть более одного
func NewPaymentModule(paymentMethods map[PaymentMethodType]PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethods: paymentMethods,
	}
}

func (p *PaymentModule) Pay(paymentMethodType PaymentMethodType, usd int, description string) Payment {
	//Получаем кастомный тип, который я уже создал, проверяем в мапе на его наличие, то есть тип оплаты (карта, paypal, крипта), проверяем, что такое есть
	methodPay, ok := p.paymentMethods[paymentMethodType]
	if !ok {
		return Payment{}
	}

	//Производим оплату со стороны самого сервиса, что всё успешно проходит, по сути, от нас протсо запустить эту функцию и всё, она уже где то там, кем то была написана, грубо говоря
	id := methodPay.Pay(usd, description)

	//Чисто теоретически на этом этапе нужно отправлять в БД на добавление, но здесь мы как бы возвращаем в main уже после добавления в БД, условно
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
func (p *PaymentModule) Refund(paymentMethodType PaymentMethodType, payment Payment) Payment {

	//Достаём нужный метод оплаты и производим отмену оплаты
	methodPay := p.paymentMethods[paymentMethodType]

	//Здесь мы производим возврат на уровне провадйера, то есть тех методов оплаты, что были переданы, все возможные
	methodPay.Refund(payment.idPay)
	payment.IsCancelled = false
	//Где то тут как бы отправляем дело в БД и от туда уже приходит результат, о том, успешно ли мы всё поменяли или нет и соответственно сам объект возвращаем

	return payment
}

// Просто вывод чека
func (p *PaymentModule) Info(payment Payment) {
	fmt.Printf("paymentId: %d, paymentDescrtiption: %s", payment.Id, payment.Description)
}
