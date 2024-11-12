package product

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	Stock       int    `json:"stock"`
	// 其他字段
}