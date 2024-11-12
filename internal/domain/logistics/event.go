package logistics

type LogisticsOrderCreated struct {
	LogisticsOrder *LogisticsOrder `json:"logistics_order"`
}

type LogisticsStatusUpdated struct {
	LogisticsID string `json:"logistics_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type LogisticsOrderCancelled struct {
	LogisticsID string `json:"logistics_id"`
}

type LogisticsOrderDelivered struct {
	LogisticsOrder *LogisticsOrder `json:"logistics_order"`
}