package models

type Product struct {
	ProductId uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}

type Products []Product
