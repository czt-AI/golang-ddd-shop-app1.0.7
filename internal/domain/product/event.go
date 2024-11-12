package product

type ProductAdded struct {
	Product *Product `json:"product"`
}

type ProductUpdated struct {
	Product *Product `json:"product"`
}

type ProductOutOfStock struct {
	ProductID int64 `json:"product_id"`
}