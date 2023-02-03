package models

// Product is a struct that contains information about the product
type Product struct {
	Category    string
	Amount      float64
	Installment int
}

// Client is a struct that contains information about the client
type Client struct {
	Name  string
	Phone string
}
