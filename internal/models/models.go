package models

import (
	"gorm.io/gorm"
)

// Product is a struct that contains information about the product
type Product struct {
	ID          uint
	Category    string
	Amount      float64
	Installment int
}

// Client is a struct that contains information about the client
type Client struct {
	ID    uint
	Name  string
	Phone string
}

// Purchase is a struct that contains information about the purchase
type Purchase struct {
	ID              uint
	Product         []Product `gorm:"many2many:purchase_products"`
	Client          []Client  `gorm:"many2many:purchase_clients"`
	TotalPayment    float64
	PaymentPerMonth float64
}

// InitModels is a function that initializes the database
func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Product{}, &Client{}, &Purchase{})
	if err != nil {
		return err
	}
	return nil
}
