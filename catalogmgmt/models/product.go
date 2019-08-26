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

// CreateProduct get a product
func (p *Product) CreateProduct(db *sql.DB) error {
	stmt := fmt.Sprintf("INSERT INTO product(productId, productName, unitPrice, availableQuantity, productImage, productSeller) values(null, '%s', '%d', '%d', '%s', '%d')", p.ProductName, p.UnitPrice, p.AvailableQuantity, p.ProductImage, p.ProductSeller)
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ProductID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProduct ...
func (p *Product) UpdateProduct(db *sql.DB) error {
	stmt := fmt.Sprintf("UPDATE product SET productName = '%s', unitPrice='%d', availableQuantity='%d', productImage='%s', productSeller='%d' WHERE productId='%d'", p.ProductName, p.UnitPrice, p.AvailableQuantity, p.ProductImage, p.ProductSeller, p.ProductID)
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
