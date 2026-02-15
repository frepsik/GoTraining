package payments

type Payment struct {
	Id          int
	Usd         int
	Description string
	IsCancelled bool
	MethodPay   string
	idPay       int
}
