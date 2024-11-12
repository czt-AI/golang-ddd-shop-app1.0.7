package v1

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/product"
)

type ProductService interface {
	GetProductByID(ctx context.Context, id int64) (*dto.ProductResponse, error)
	GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*dto.ProductResponse, error)
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, req *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(ctx context.Context, id int64) error
}

type productService struct {
	usecase product.ProductUsecase
}

func NewProductService(usecase product.ProductUsecase) *productService {
	return &productService{usecase: usecase}
}

func (s *productService) GetProductByID(ctx context.Context, id int64) (*dto.ProductResponse, error) {
	product, err := s.usecase.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ConvertToProductResponse(product), nil
}

func (s *productService) GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*dto.ProductResponse, error) {
	products, err := s.usecase.GetProductList(ctx, categoryID, tag, page, limit)
	if err != nil {
		return nil, err
	}
	responseProducts := make([]*dto.ProductResponse, 0, len(products))
	for _, p := range products {
		responseProducts = append(responseProducts, dto.ConvertToProductResponse(p))
	}
	return responseProducts, nil
}

func (s *productService) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.usecase.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.ConvertToProductResponse(product), nil
}

func (s *productService) UpdateProduct(ctx context.Context, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.usecase.UpdateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return dto.ConvertToProductResponse(product), nil
}

func (s *productService) DeleteProduct(ctx context.Context, id int64) error {
	return s.usecase.DeleteProduct(ctx, id)
}