package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/model"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product model.Product) (model.Product, error) {
	query := `INSERT INTO products (id, name, price, stock, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, product.ID, product.Name, product.Price, product.Stock, product.CreatedAt)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *productRepository) GetByID(id uuid.UUID) (model.Product, error) {
	var p model.Product
	query := `SELECT id, name, price, stock, created_at FROM products WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)
	if err != nil {
		return model.Product{}, err
	}
	return p, nil
}

func (r *productRepository) GetAll(nameFilter string) ([]model.Product, error) {
	query := `SELECT id, name, price, stock, created_at FROM products`

	args := []interface{}{}
	if nameFilter != "" {
		query += " WHERE name ILIKE $1"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *productRepository) Update(product model.Product) (model.Product, error) {
	query := `UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4`
	_, err := r.db.Exec(query, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *productRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
