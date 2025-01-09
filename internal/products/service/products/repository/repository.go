package repository

import (
	"fmt"

	"github.com/Dmitriy-M1319/crystal-services/internal/products/service/products"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	GetProducts() ([]products.Product, error)
	GetProductByID(id uint64) (products.Product, error)
	InsertProduct(product *products.Product) error
	UpdateProduct(id uint64, product *products.Product) error
	DeleteProduct(id uint64) error
}

type ProductRepositoryImpl struct {
	connection *sqlx.DB
}

func NewProductRepositoryImpl(db *sqlx.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{connection: db}
}

func (repo *ProductRepositoryImpl) GetProducts() ([]products.Product, error) {
	query := "SELECT * FROM products ORDER BY id"
	var result []products.Product
	err := repo.connection.Select(result, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %s", err)
	}

	return result, nil
}

func (repo *ProductRepositoryImpl) GetProductByID(id uint64) (products.Product, error) {
	query := "SELECT * FROM products WHERE id = $1"
	result := products.Product{}
	err := repo.connection.Get(&result, query, id)
	if err != nil {
		return result, fmt.Errorf("failed to get one product: %s", err)
	}
	return result, nil
}

func (repo *ProductRepositoryImpl) InsertProduct(product *products.Product) error {
	query := `INSERT INTO products(product_name, company_name, client_price, purchase_price, count_on_warehouse)
			VALUES($1, $2, $3, $4, $5) RETURNING id`
	newId := 0
	err := repo.connection.Get(&newId, query, product.ProductName, product.CompanyName,
		product.ClientPrice, product.PurchasePrice, product.CountOnWarehouse)

	if err != nil {
		return fmt.Errorf("failed to insert product: %s", err)
	}

	product.ID = uint64(newId)
	return nil
}

func (repo *ProductRepositoryImpl) UpdateProduct(id uint64, product *products.Product) error {
	_, err := repo.GetProductByID(id)
	if err != nil {
		return err
	}

	query := `UPDATE products SET product_name=$1 company_name=$2 client_price=$3 purchase_price=$4 count_on_warehouse=$5
			WHERE id = $6`
	_, err = repo.connection.Exec(query, product.ProductName, product.CompanyName,
		product.ClientPrice, product.PurchasePrice, product.CountOnWarehouse, id)

	if err != nil {
		return fmt.Errorf("failed to update product: %s", err)
	}
	return nil
}

func (repo *ProductRepositoryImpl) DeleteProduct(id uint64) error {
	_, err := repo.GetProductByID(id)
	if err != nil {
		return err
	}

	query := "DELETE FROM products WHERE id = $1"
	_, err = repo.connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %s", err)
	}

	return nil
}
