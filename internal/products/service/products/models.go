package products

type Product struct {
	ID               uint64  `db:"id"`
	ProductName      string  `db:"product_name"`
	CompanyName      string  `db:"company_name"`
	ClientPrice      float32 `db:"client_price"`
	PurchasePrice    float32 `db:"purchase_price"`
	CountOnWarehouse uint32  `db:"count_on_warehouse"`
}

type CompanyNameFilter struct {
	Name string
}

type ClientPriceFilter struct {
	From float32
	To   float32
}

func FilterByCompany(products []Product, filter *CompanyNameFilter) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if p.CompanyName == filter.Name {
			result = append(result, p)
		}
	}

	return result
}

func FilterByPrice(products []Product, filter *ClientPriceFilter) []Product {
	result := make([]Product, 0)
	for _, p := range products {
		if p.ClientPrice >= filter.From && p.ClientPrice <= filter.From {
			result = append(result, p)
		}
	}

	return result
}
