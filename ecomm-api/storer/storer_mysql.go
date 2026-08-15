package storer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MySQLStorer struct {
	db *sqlx.DB
}

func NewSQLStorer(db *sqlx.DB) *MySQLStorer {
	return &MySQLStorer{
		db: db,
	}
}

func (ms *MySQLStorer) CreateProduct(ctx context.Context, p *Product) (*Product, error) {
	res, err := ms.db.NamedExecContext(ctx, "INSERT INTO products (name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)", p)

	if err != nil {
		return nil, fmt.Errorf("error inserting product %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID %w", err)
	}
	p.ID = id

	return p, nil
}

func (ms *MySQLStorer) GetProduct(ctx context.Context, id int64) (*Product, error) {
	var p Product
	err := ms.db.GetContext(ctx, &p, "SELECT * FROM products WHERE id =?", id)
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID %w", err)
	}

	return &p, nil
}

func (ms *MySQLStorer) ListProducts(ctx context.Context, id int64) ([]*Product, error) {
	var p []*Product
	err := ms.db.SelectContext(ctx, &p, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID %w", err)
	}

	return p, nil
}

func (ms *MySQLStorer) UpdateProducts(ctx context.Context, p *Product) (*Product, error) {
	_, err := ms.db.NamedExecContext(ctx, "UPDATE products SET name=:name, images=:image, category=:category, description=:description, price=:price", p)

	if err != nil {
		return nil, fmt.Errorf("error updating product %w", err)
	}

	return p, nil
}

func (ms *MySQLStorer) DeleteProduct(ctx context.Context, id int64) error {
	_, err := ms.db.ExecContext(ctx, "DELETE FROM products WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("error deleting product %w", err)
	}
	return nil
}
