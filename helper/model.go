package helper

import (
	"test/app/model"
	"test/log/response"
)

func ToProductResponse(product model.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:          product.ID,
		NamaBarang:  product.NamaBarang,
		Harga:       product.Harga,
		Jenis:       product.Jenis,
		MetaKeyword: product.MetaKeyword,
	}
}
