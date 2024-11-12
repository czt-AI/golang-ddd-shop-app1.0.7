package order

type OrderCreated struct {
	Order *Order `json:"order"`
}

type OrderUpdated struct {
	Order *Order `json:"order"`
}

type OrderCancelled struct {
	OrderID string `json:"order_id"`
}

type OrderCompleted struct {
	Order *Order `json:"order"`
}