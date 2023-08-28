package repository

import (
	"context"
	"database/sql"
	"test/app/model"
)

// interface for Product repository implementation
type ProductRepository interface {
	SaveProduct(ctx context.Context, tx *sql.Tx, model model.Product) model.Product
}

// Product repository implementation
type ProductRepositoryImpl struct {
}

// new Product repository
func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// SaveProduct implements ProductRepository.
func (repository *ProductRepositoryImpl) SaveProduct(ctx context.Context, tx *sql.Tx, model model.Product) model.Product {
	panic("unimplemented")
}
