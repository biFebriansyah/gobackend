package models

import "time"

type Product struct {
	ProductId   string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type Products []Product
