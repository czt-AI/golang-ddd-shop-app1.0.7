package payment

type Payment struct {
	TransactionID string  `json:"transaction_id"`
	OrderID       string  `json:"order_id"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
}