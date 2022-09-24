package models

import "time"

type Cart struct {
	CartId     uint      `gorm:"primaryKey" json:"cart_id,omitempty"`
	UsersId    uint      `json:"userId"`
	ProductsId uint      `json:"productId"`
	Products   []Product `gorm:"foreignKey:ProductsId; references:ProductId;"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}

type CartItem struct {
	CartItemId uint    `gorm:"primaryKey" json:"id,omitempty"`
	CartId     uint    `json:"user_id"`
	ProductsId uint    `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductsId;references:ProductId;"`
	Quantity   int     `json:"quantity"`
}

type Carts []Cart
