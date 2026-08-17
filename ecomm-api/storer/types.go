package storer

import "time"

type Product struct {
	ID           int64      `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Image        string     `json:"image" db:"image"`
	Category     string     `json:"category" db:"category"`
	Description  string     `json:"description" db:"description"`
	Rating       int64      `json:"rating" db:"rating"`
	NumReviews   int64      `json:"num_reviews" db:"num_reviews"`
	Price        float64    `json:"price" db:"price"`
	CountInStock int64      `json:"count_in_stock" db:"count_in_stock"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt    *time.Time `json:"updatedAt" db:"updated_at"`
}

type Order struct {
	ID            int64      `json:"id" db:"id"`
	PaymentMethod string     `json:"payment_method" db:"payment_method"`
	TaxPrice      float64    `json:"tax_price" db:"tax_price"`
	ShippingPrice float64    `json:"shipping_price" db:"shipping_price"`
	TotalPrice    float64    `json:"total_price" db:"total_price"`
	CreatedAt     time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt     *time.Time `json:"updatedAt" db:"updated_at"`
	Items         []OrderItem
}

type OrderItem struct {
	ID        int64   `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Quantity  int64   `json:"quantity" db:"quantity"`
	Image     string  `json:"image" db:"image"`
	Price     float64 `json:"price" db:"price"`
	ProductID int64   `json:"product_id" db:"product_id"`
	OrderID   int64   `json:"order_id" db:"order_id"`
}
