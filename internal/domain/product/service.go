package product

import (
	"context"
	"shop-app/internal/domain/product/repository"
)

type ProductService interface {
	GetProductByID(ctx context.Context, id int64) (*Product, error)
	GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService {
	return &productService{repo: repo}
}

func (s *productService) GetProductByID(ctx context.Context, id int64) (*Product, error) {
	dbProduct, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ConvertToProduct(dbProduct), nil
}

func (s *productService) GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*Product, error) {
	dbProducts, err := s.repo.GetProductList(ctx, categoryID, tag, page, limit)
	if err != nil {
		return nil, err
	}
	products := make([]*Product, 0, len(dbProducts))
	for _, dbProduct := range dbProducts {
		products = append(products, ConvertToProduct(dbProduct))
	}
	return products, nil
}

func ConvertToProduct(dbProduct *dao.Product) *Product {
	return &Product{
		ID:          dbProduct.ID,
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
		Price:       dbProduct.Price,
		Stock:       dbProduct.Stock,
		// 其他字段转换
	}
}