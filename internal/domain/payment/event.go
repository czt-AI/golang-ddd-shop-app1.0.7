package payment

type PaymentRequested struct {
	OrderID       string  `json:"order_id"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
}

type PaymentSucceeded struct {
	TransactionID string `json:"transaction_id"`
}

type PaymentFailed struct {
	TransactionID string `json:"transaction_id"`
}

type RefundRequested struct {
	TransactionID string `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}

type RefundCompleted struct {
	TransactionID string `json:"transaction_id"`
}