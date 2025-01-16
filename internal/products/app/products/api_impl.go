package products

import (
	"github.com/Dmitriy-M1319/crystal-services/internal/products/service/products"
	desc "github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
)

type ApiProductsImplementation struct {
	desc.UnimplementedProductsServiceServer
	service products.Service
}

func NewApiProductsImplementation(s products.Service) desc.ProductsServiceServer {
	return &ApiProductsImplementation{service: s}
}
