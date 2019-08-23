package models

import (
	"database/sql"
	"fmt"
)

// Product ...
type product struct {
	ProductID         int    `json:"productId"`
	ProductName       string `json:"productName"`
	ProductImage      string `json:"productImage"`
	AvailableQuantity int    `json:"availableQuantity"`
	UnitPrice         int    `json:"unitPrice"`
	ProductSeller     int    `json:"productSeller"`
}

// GetProducts ...
func GetProducts(db *sql.DB) ([]product, error) {
	stmt := fmt.Sprintf("SELECT * from Product")
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []product{}
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductImage, &p.AvailableQuantity, &p.UnitPrice, &p.ProductSeller); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
