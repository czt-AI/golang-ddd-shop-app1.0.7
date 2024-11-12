package product

import (
	"context"
	"shop-app/internal/dao"
)

type ProductRepository interface {
	GetProductByID(ctx context.Context, id int64) (*dao.Product, error)
	GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*dao.Product, error)
}

type productRepository struct {
	db *dao.DB
}

func NewProductRepository(db *dao.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetProductByID(ctx context.Context, id int64) (*dao.Product, error) {
	return dao.GetProductByID(r.db, id)
}

func (r *productRepository) GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*dao.Product, error) {
	return dao.GetProductList(r.db, categoryID, tag, page, limit)
}