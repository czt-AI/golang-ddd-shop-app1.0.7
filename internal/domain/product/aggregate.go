package product

import (
	"context"
	"shop-app/internal/domain/product/event"
	"shop-app/internal/domain/product/repository"
)

type ProductAggregate struct {
	repo repository.ProductRepository
}

func NewProductAggregate(repo repository.ProductRepository) *ProductAggregate {
	return &ProductAggregate{repo: repo}
}

func (a *ProductAggregate) AddProduct(ctx context.Context, product *Product) error {
	// 创建数据库记录
	dbProduct := ConvertToDBProduct(product)
	if err := a.repo.CreateProduct(ctx, dbProduct); err != nil {
		return err
	}

	// 发布事件
	event := event.ProductAdded{Product: product}
	// 这里可以进一步处理事件发布逻辑，例如通知其他服务

	return nil
}

func (a *ProductAggregate) UpdateProduct(ctx context.Context, product *Product) error {
	// 更新数据库记录
	dbProduct := ConvertToDBProduct(product)
	if err := a.repo.UpdateProduct(ctx, dbProduct); err != nil {
		return err
	}

	// 发布事件
	event := event.ProductUpdated{Product: product}
	// 这里可以进一步处理事件发布逻辑，例如通知其他服务

	return nil
}

func (a *ProductAggregate) CheckStock(ctx context.Context, productID int64) (bool, error) {
	// 检查库存
	stock, err := a.repo.GetProductStock(ctx, productID)
	if err != nil {
		return false, err
	}
	return stock > 0, nil
}

func ConvertToDBProduct(product *Product) *dao.Product {
	return &dao.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		// 其他字段转换
	}
}