package models

import "time"

type User struct {
	UserId    string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"password,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Users []User
