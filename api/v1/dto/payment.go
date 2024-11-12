package v1

type RecordPaymentRequest struct {
	OrderID       string  `json:"order_id"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
}

type ConfirmPaymentRequest struct {
	TransactionID string `json:"transaction_id"`
}

type RefundPaymentRequest struct {
	TransactionID string `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}