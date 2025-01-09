package products

type Product struct {
	ID               uint64  `db:"id"`
	ProductName      string  `db:"product_name"`
	CompanyName      string  `db:"company_name"`
	ClientPrice      float32 `db:"client_price"`
	PurchasePrice    float32 `db:"purchase_price"`
	CountOnWarehouse uint32  `db:"count_on_warehouse"`
}

func FilterByCompany(products []Product, name string) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if p.CompanyName == name {
			result = append(result, p)
		}
	}

	return result
}

func FilterByPrice(products []Product, from, to float32) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if p.ClientPrice >= from && p.ClientPrice <= to {
			result = append(result, p)
		}
	}

	return result
}
