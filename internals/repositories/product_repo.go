package repositories

import (
	model "crud-hex/internals/core/domain"
	"crud-hex/internals/core/ports"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type ProductRepository struct {
	db *sql.DB
}

// Create implements ports.IProductRepository.
func (r *ProductRepository) Create(product *model.Product) error {
	result, err := r.db.Exec("INSERT INTO products (name, stock, price, is_available) VALUES (?,?,?,?)", 
		product.Name, product.Stock, product.Price, product.IsAvailable)
	
	if err != nil {
		return err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Update the ID of the product object
	product.ID = int(id)
	return nil
}

// FindAll implements ports.IProductRepository.
func (r *ProductRepository) FindAll() ([]model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, stock, price, is_available FROM products WHERE is_available = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Stock, &p.Price, &p.IsAvailable); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (model.Product, error) {
	var product model.Product
	err := r.db.QueryRow("SELECT id, name, stock, price, is_available FROM products WHERE id = ?", id).Scan(
		&product.ID, &product.Name, &product.Stock, &product.Price, &product.IsAvailable)
	return product, err
}

func (r *ProductRepository) Update(product model.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = ?, stock = ?, price = ? WHERE id = ?", product.Name, product.Stock, product.Price ,product.ID)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.db.Exec("UPDATE products SET is_available = false WHERE id = ?", id)
	return err
}

func NewProductRepository(db *sql.DB) ports.IProductRepository {
	return &ProductRepository{db: db}
}
