package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"kaspi.nurgalym.net/internal/validator"
	"time"
)

type Product struct {
	ID         int64     `json:"id"`                   // Unique integer ID for the product
	CreatedAt  time.Time `json:"-"`                    // Timestamp for when the product is added to our database
	Title      string    `json:"title"`                // Product title
	Year       int32     `json:"year,omitempty"`       // Product release year
	Price      Price     `json:"price,omitempty"`      // Product price (in tenge)
	Categories []string  `json:"categories,omitempty"` // Product category (technic, for home, etc.)
	Version    int32     `json:"version"`              // The version number starts at 1 and will be incremented each
}

func (p Product) MarshalJSON() ([]byte, error) {
	var price string

	if p.Price != 0 {
		price = fmt.Sprintf("%d tenge", p.Price)
	}

	type ProductAlias Product

	aux := struct {
		ProductAlias
		Price string `json:"price,omitempty"`
	}{
		ProductAlias: ProductAlias(p),
		Price:        price,
	}
	return json.Marshal(aux)
}

func ValidateProduct(v *validator.Validator, p *Product) {
	v.Check(p.Title != "", "title", "must be provided")
	v.Check(len(p.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(p.Year != 0, "year", "must be provided")
	v.Check(p.Year >= 1888, "year", "must be greater than 1888")
	v.Check(p.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(p.Price != 0, "price", "must be provided")
	v.Check(p.Price > 0, "price ", "must be a positive integer")

	v.Check(p.Categories != nil, "categories", "must be provided")
	v.Check(len(p.Categories) >= 1, "categories", "must contain at least 1 category")
	v.Check(len(p.Categories) <= 5, "categories", "must not contain more than 5 categories")

	v.Check(validator.Unique(p.Categories), "categories", "must not contain duplicate values")
}

type ProductModel struct {
	DB *sql.DB
}

func (p ProductModel) Insert(product *Product) error {
	//return nil
	query := `
		INSERT INTO products (title, year, price, categories)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`

	args := []interface{}{product.Title, product.Year, product.Price, pq.Array(product.Categories)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return p.DB.QueryRowContext(ctx, query, args...).Scan(&product.ID, &product.CreatedAt, &product.Version)
}

func (p ProductModel) Get(id int64) (*Product, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, title, year, price, categories, version
		FROM products
		WHERE id = $1`

	var product Product

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.CreatedAt,
		&product.Title,
		&product.Year,
		&product.Price,
		pq.Array(&product.Categories),
		&product.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &product, nil

	//return nil, nil
}

func (p ProductModel) Update(product *Product) error {
	query := `
		UPDATE products
		SET title = $1, year = $2, price = $3, categories = $4, version = version + 1
		WHERE id = $5
		RETURNING version`

	args := []interface{}{
		product.Title,
		product.Year,
		product.Price,
		pq.Array(product.Categories),
		product.ID,
		product.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&product.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m ProductModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
		DELETE FROM products
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil

}

func (p ProductModel) GetAll(title string, categories []string, filters Filters) ([]*Product, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(),id, created_at, title, year, price, categories, version
		FROM products
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (categories @> $2 OR $2 = '{}')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, pq.Array(categories), filters.limit(), filters.offset()}
	rows, err := p.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	products := []*Product{}

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&totalRecords,
			&product.ID,
			&product.CreatedAt,
			&product.Title,
			&product.Year,
			&product.Price,
			pq.Array(&product.Categories),
			&product.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		products = append(products, &product)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return products, metadata, nil

}
