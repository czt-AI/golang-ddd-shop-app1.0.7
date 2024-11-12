package product

import (
	"context"
	"shop-app/domain/product"
)

type ProductUsecase interface {
	GetProductByID(ctx context.Context, id int64) (*product.Product, error)
	GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*product.Product, error)
	CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.Product, error)
	UpdateProduct(ctx context.Context, req *product.UpdateProductRequest) (*product.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}

type productUsecase struct {
	repo product.ProductRepository
}

func NewProductUsecase(repo product.ProductRepository) *productUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) GetProductByID(ctx context.Context, id int64) (*product.Product, error) {
	dbProduct, err := u.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product.ConvertToProduct(dbProduct), nil
}

func (u *productUsecase) GetProductList(ctx context.Context, categoryID int64, tag string, page, limit int) ([]*product.Product, error) {
	dbProducts, err := u.repo.GetProductList(ctx, categoryID, tag, page, limit)
	if err != nil {
		return nil, err
	}
	products := make([]*product.Product, 0, len(dbProducts))
	for _, dbProduct := range dbProducts {
		products = append(products, product.ConvertToProduct(dbProduct))
	}
	return products, nil
}

func (u *productUsecase) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.Product, error) {
	dbProduct := product.ConvertToDBProduct(req)
	if err := u.repo.CreateProduct(ctx, dbProduct); err != nil {
		return nil, err
	}
	return product.ConvertToProduct(dbProduct), nil
}

func (u *productUsecase) UpdateProduct(ctx context.Context, req *product.UpdateProductRequest) (*product.Product, error) {
	dbProduct := product.ConvertToDBProduct(req)
	if err := u.repo.UpdateProduct(ctx, dbProduct); err != nil {
		return nil, err
	}
	return product.ConvertToProduct(dbProduct), nil
}

func (u *productUsecase) DeleteProduct(ctx context.Context, id int64) error {
	return u.repo.DeleteProduct(ctx, id)
}