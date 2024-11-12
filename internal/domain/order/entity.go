package order

type Order struct {
	OrderID     string  `json:"order_id"`
	UserID      int64   `json:"user_id"`
	ProductList []OrderItem `json:"product_list"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type OrderItem struct {
	ProductID int64  `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     float64 `json:"price"`
}