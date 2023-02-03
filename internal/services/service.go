package services

import (
	"fmt"

	"github.com/ProninIgorr/alif-task-payments/internal/models"
)

const (
	SmartphonePercent = 3
	ComputerPercent   = 4
	TVPercent         = 5
)

// CheckCategory checks the category of the product
func CheckCategory(product *models.Product) error {
	if product.Category != "Smartphone" && product.Category != "Computer" && product.Category != "TV" {
		return fmt.Errorf("Нет данных о категории товара")
	}
	return nil
}

// NewClient creates a new client
func NewClient(name string, phone string) *models.Client {
	return &models.Client{
		Name:  name,
		Phone: phone,
	}
}

// CheckInstallment checks the number of installments
func CheckInstallment(product *models.Product) error {
	if product.Installment < 3 || product.Installment > 24 {
		return fmt.Errorf("Количество месяцев платежей должно быть от 3 до 24")
	}
	return nil
}

// TotalAmount calculates the total amount of the payment
func TotalAmount(p *models.Product) float64 {
	var percent float64
	switch p.Category {
	case "Smartphone":
		percent = SmartphonePercent
		if p.Installment > 9 && p.Installment <= 12 {
			percent = percent * 2
		} else if p.Installment > 12 && p.Installment <= 18 {
			percent = percent * 3
		} else if p.Installment > 18 && p.Installment <= 24 {
			percent = percent * 4
		}
	case "Computer":
		percent = ComputerPercent
		if p.Installment > 12 && p.Installment <= 18 {
			percent = percent * 2
		} else if p.Installment > 18 && p.Installment <= 24 {
			percent = percent * 3
		}
	case "TV":
		percent = TVPercent
		if p.Installment > 18 && p.Installment <= 24 {
			percent = percent * 2
		}

	}
	return p.Amount + p.Amount*percent/100
}

// PaymentPerMonth calculates the amount of the payment per month
func PaymentPerMonth(product *models.Product) float64 {
	return TotalAmount(product) / float64(product.Installment)
}

// SMSNotification sends an SMS notification to the client
func SMSNotification(product models.Product, client models.Client, totalPayment float64, PaymentperMonth float64) {
	fmt.Println("Отправка SMS уведомления клиенту:", client.Name)
	fmt.Println("Телефон:", client.Phone)
	fmt.Println("Товар:", product.Category)
	fmt.Println("Цена:", product.Amount, "сомони")
	fmt.Println("Количество месяцев платежей:", product.Installment)
	fmt.Println("Общая сумма платежа:", totalPayment, "сомони")
	fmt.Printf("Сумма платежа в месяц: %.2f сомони", PaymentperMonth)
}
