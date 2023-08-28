package controller

import (
	"encoding/json"
	"net/http"
	"test/app/service"
	"test/helper"
	"test/log/request"
	"test/log/response"

	"github.com/julienschmidt/httprouter"
)

// interface for product controller
type ProductController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

// product for route in main
func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

// struct for category product impl
type ProductControllerImpl struct {
	ProductService service.ProductService
}

func (controller *ProductControllerImpl) CreateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	productCreateRequest := request.ProductCreateRequest{}
	err := decoder.Decode(&productCreateRequest)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.CreateProduct(r.Context(), productCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
