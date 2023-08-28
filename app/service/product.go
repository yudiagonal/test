package service

import (
	"context"
	"database/sql"

	"test/app/repository"
	"test/log/request"
	"test/log/response"

	"github.com/go-playground/validator"
)

type ProductService interface {
	CreateProduct(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
}

// struct for implementation Product service
type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}
