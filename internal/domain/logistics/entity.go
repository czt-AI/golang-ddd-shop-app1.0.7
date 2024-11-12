package logistics

type LogisticsOrder struct {
	LogisticsID string  `json:"logistics_id"`
	OrderID     string  `json:"order_id"`
	LogisticsCompany string `json:"logistics_company"`
	ShippingTime string `json:"shipping_time"`
	DeliveryTime string `json:"delivery_time"`
	Status      string  `json:"status"`
	Description string `json:"description"`
}