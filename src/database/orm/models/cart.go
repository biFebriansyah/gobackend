package models

import "time"

type Cart struct {
	CartId     uint      `gorm:"primaryKey" json:"cart_id,omitempty"`
	UsersId    uint      `json:"userId"`
	ProductsId uint      `json:"productId"`
	Products   Product   `gorm:"foreignKey:ProductsId; references:ProductId;"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}

type Carts []Cart
