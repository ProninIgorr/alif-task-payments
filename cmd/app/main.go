package main

import (
	"fmt"

	"github.com/ProninIgorr/alif-task-payments/internal/models"
	"github.com/ProninIgorr/alif-task-payments/internal/services"
)

// P.S. В ТЗ немного неправильно была формулировака рассчета процента для смартфона
// Потому что диапазаон в 3% для смартофона относится до 9 месяцев, а не до 12
// Следовательно 12 месяцев новый диапазон, для этотого диапазона процент = 6%
//А для 18 месяцев процент равен 9%, а не 6% как указано в ТЗ

func main() {

	product := &models.Product{
		Category:    "Smartphone",
		Amount:      1000,
		Installment: 12,
	}

	client := services.NewClient("Igor", "+992 907 77 77 77")

	error := services.CheckCategory(product)
	if error != nil {
		fmt.Println(error)
		return
	}

	err := services.CheckInstallment(product)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalPayment := services.TotalAmount(product)
	paymentPerMonth := services.PaymentPerMonth(product)

	services.SMSNotification(*product, *client, totalPayment, paymentPerMonth)

}
