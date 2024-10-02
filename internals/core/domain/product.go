// internals\core\domain\product.go

package domain

type Product struct {
	ID          int
	Name        string
	Price       int
	Stock       int
	IsAvailable bool
}

func NewProduct(id int, name string, price int, stock int, isAvailable bool) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Price:       price,
		Stock:       stock,
		IsAvailable: isAvailable,
	}
}
