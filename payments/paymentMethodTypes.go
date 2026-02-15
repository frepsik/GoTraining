package payments

type paymentMethodType string

const (
	PayPalMethod paymentMethodType = "paypal"
	CardMethod   paymentMethodType = "card"
	CryptoMethod paymentMethodType = "crypto"
)
