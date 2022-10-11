package models

import "time"

type Cart struct {
	CartId    string     `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	UsersId   string     `gorm:"unique;" json:"user_id"`
	Items     []CartItem `gorm:"foreignKey:CartId;references:CartId;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  time.Time  `json:"update_at"`
}

type CartItem struct {
	CartItemId string  `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	CartId     string  `json:"cart_id"`
	ProductsId string  `json:"product_id"`
	Products   Product `gorm:"foreignKey:ProductsId;references:ProductId;"`
	Quantity   int     `json:"quantity"`
}

type Carts []Cart
type CartItems []CartItem
