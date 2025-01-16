package products

import (
	"context"

	"github.com/Dmitriy-M1319/crystal-services/internal/products/service/products"
	pb "github.com/Dmitriy-M1319/crystal-services/pkg/crystal-services/products/v1"
)

func (pImpl *ProductsApiImplementation) GetProducts(ctx context.Context, in *pb.ProductFilters) (*pb.Products, error) {
	var nameFilter *products.CompanyNameFilter = nil
	var priceFilter *products.ClientPriceFilter = nil
	if len(in.Name) > 0 {
		nameFilter = &products.CompanyNameFilter{Name: in.Name}
	}
	if in.Price.From != 0.0 && in.Price.To != 0.0 {
		priceFilter = &products.ClientPriceFilter{From: in.Price.From, To: in.Price.To}
	}

	products, err := pImpl.service.GetAllProducts(ctx, nameFilter, priceFilter)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Product, len(products))
	for i, p := range products {
		var pr = &pb.Product{
			Id:               p.ID,
			ProductName:      p.ProductName,
			CompanyName:      p.CompanyName,
			ClientPrice:      p.ClientPrice,
			PurchasePrice:    p.PurchasePrice,
			CountOnWarehouse: p.CountOnWarehouse,
		}

		res[i] = pr
	}
	return &pb.Products{Products: res}, nil
}

func (pImpl *ProductsApiImplementation) GetProductById(ctx context.Context, in *pb.ProductId) (*pb.Product, error) {
	product, err := pImpl.service.GetProduct(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	var result = &pb.Product{
		Id:               product.ID,
		ProductName:      product.ProductName,
		CompanyName:      product.CompanyName,
		ClientPrice:      product.ClientPrice,
		PurchasePrice:    product.PurchasePrice,
		CountOnWarehouse: product.CountOnWarehouse,
	}

	return result, nil
}

func (pImpl *ProductsApiImplementation) InsertProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	inProduct := &products.Product{
		ID:               in.Id,
		ProductName:      in.ProductName,
		CompanyName:      in.CompanyName,
		ClientPrice:      in.ClientPrice,
		PurchasePrice:    in.PurchasePrice,
		CountOnWarehouse: in.CountOnWarehouse,
	}

	err := pImpl.service.AddProduct(ctx, inProduct)
	if err != nil {
		return nil, err
	}

	in.Id = inProduct.ID
	return in, nil
}

func (pImpl *ProductsApiImplementation) UpdateProduct(ctx context.Context, in *pb.ProductPutRequest) (*pb.Product, error) {
	inProduct := &products.Product{
		ID:               in.Id,
		ProductName:      in.Product.ProductName,
		CompanyName:      in.Product.CompanyName,
		ClientPrice:      in.Product.ClientPrice,
		PurchasePrice:    in.Product.PurchasePrice,
		CountOnWarehouse: in.Product.CountOnWarehouse,
	}

	err := pImpl.service.UpdateProduct(ctx, in.Id, inProduct)
	if err != nil {
		return nil, err
	}

	var result = &pb.Product{
		Id:               in.Id,
		ProductName:      inProduct.ProductName,
		CompanyName:      inProduct.CompanyName,
		ClientPrice:      inProduct.ClientPrice,
		PurchasePrice:    inProduct.PurchasePrice,
		CountOnWarehouse: inProduct.CountOnWarehouse,
	}
	return result, nil
}

func (pImpl *ProductsApiImplementation) DeleteProduct(ctx context.Context, in *pb.ProductId) (*pb.Empty, error) {
	err := pImpl.service.DeleteProduct(ctx, in.Id)

	if err != nil {
		return nil, err
	} else {
		return &pb.Empty{}, nil
	}
}
