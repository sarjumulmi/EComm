package models

import (
	"database/sql"
	"fmt"
)

// Product ...
type Product struct {
	ProductID         int    `json:"productId"`
	ProductName       string `json:"productName"`
	ProductImage      string `json:"productImage"`
	AvailableQuantity int    `json:"availableQuantity"`
	UnitPrice         int    `json:"unitPrice"`
	ProductSeller     int    `json:"productSeller"`
}

// GetProducts ...
func GetProducts(db *sql.DB) ([]Product, error) {
	stmt := fmt.Sprintf("SELECT * from Product")
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ProductID, &p.ProductName, &p.UnitPrice, &p.AvailableQuantity, &p.ProductImage, &p.ProductSeller); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// GetProduct get a product
func (p *Product) GetProduct(db *sql.DB) error {
	stmt := fmt.Sprintf("SELECT * from Product WHERE productId=%d", p.ProductID)
	return db.QueryRow(stmt).Scan(&p.ProductID, &p.ProductName, &p.UnitPrice, &p.AvailableQuantity, &p.ProductImage, &p.ProductSeller)
}
