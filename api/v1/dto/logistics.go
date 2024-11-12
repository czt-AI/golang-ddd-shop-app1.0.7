package v1

type CreateLogisticsRequest struct {
	LogisticsOrder LogisticsOrder `json:"logistics_order"`
}

type UpdateLogisticsStatusRequest struct {
	LogisticsID    string `json:"logistics_id"`
	Status         string `json:"status"`
	Description    string `json:"description"`
}

type LogisticsOrder struct {
	OrderID        string `json:"order_id"`
	LogisticsCompany string `json:"logistics_company"`
	ShippingTime   string `json:"shipping_time"`
	DeliveryTime   string `json:"delivery_time"`
	Status         string `json:"status"`
	Description    string `json:"description"`
}

type LogisticsOrderResponse struct {
	LogisticsOrder LogisticsOrder `json:"logistics_order"`
}