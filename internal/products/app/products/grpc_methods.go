package products

import (
	"context"

	pb "github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
)

func (pImpl *ProductsApiImplementation) GetProducts(ctx context.Context, in *pb.Empty) (*pb.Products, error) {
	return nil, nil
}

func (pImpl *ProductsApiImplementation) GetProductById(ctx context.Context, in *pb.ProductId) (*pb.Product, error) {
	return nil, nil
}

func (pImpl *ProductsApiImplementation) InsertProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	return nil, nil
}

func (pImpl *ProductsApiImplementation) UpdateProduct(ctx context.Context, in *pb.ProductPutRequest) (*pb.Product, error) {
	return nil, nil
}

func (pImpl *ProductsApiImplementation) DeleteProduct(ctx context.Context, in *pb.ProductId) (*pb.Empty, error) {
	return nil, nil
}
