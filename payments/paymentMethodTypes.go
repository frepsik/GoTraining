package payments

type PaymentMethodType string

const (
	PayPalMethod PaymentMethodType = "paypal"
	CardMethod   PaymentMethodType = "card"
	CryptoMethod PaymentMethodType = "crypto"
)
