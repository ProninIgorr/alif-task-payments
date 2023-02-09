package app

import (
	"fmt"

	"github.com/ProninIgorr/alif-task-payments/internal/models"
	"github.com/ProninIgorr/alif-task-payments/internal/services"
	"gorm.io/gorm"
)

// Start is a function that starts the application
func Start(db *gorm.DB) {
	createPurhae(db)
	getPurchase(db)
}

// getPurchase is a function that gets all purchases
func getPurchase(db *gorm.DB) {
	var purchase []models.Purchase
	result := db.Preload("Product").Preload("Client").Find(&purchase)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println(purchase)
}

// createPurhae is a function that creates a purchase
func createPurhae(db *gorm.DB) {

	product := &models.Product{
		Category:    "Smartphone",
		Amount:      10000,
		Installment: 10,
	}

	err := services.CheckCategory(product)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = services.CheckInstallment(product)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalAmount := services.TotalAmount(product)
	paymentPerMonth := totalAmount / float64(product.Installment)

	result := db.Create(product)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	client := &models.Client{
		Name:  "Umed",
		Phone: "+992-907-67-67-67",
	}

	result = db.Create(client)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	purchase := &models.Purchase{

		Product:         []models.Product{*product},
		Client:          []models.Client{*client},
		TotalPayment:    totalAmount,
		PaymentPerMonth: paymentPerMonth,
	}

	result = db.Create(purchase)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	fmt.Println("Purchase created successfully")

	services.SMSNotification(*product, *client, totalAmount, paymentPerMonth)

}
