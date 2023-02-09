package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ProninIgorr/alif-task-payments/app"
	"github.com/ProninIgorr/alif-task-payments/internal/models"
)

// P.S. В ТЗ немного неправильно была формулировака рассчета процента для смартфона
// Потому что диапазаон в 3% для смартофона относится до 9 месяцев, а не до 12
// Следовательно 12 месяцев новый диапазон, для этотого диапазона процент = 6%
//А для 18 месяцев процент равен 9%, а не 6% как указано в ТЗ

func main() {

	//Connect to database
	dsn := "host=localhost user=postgres password=postgres dbname=alif_payments port=54320 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Database connected")

	err = models.InitModels(db)
	if err != nil {
		panic("failed to init models")
	}

	app.Start(db)

}

