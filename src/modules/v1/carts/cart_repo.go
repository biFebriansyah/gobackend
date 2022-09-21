package carts

import (
	"errors"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"gorm.io/gorm"
)

type cart_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *cart_repo {
	return &cart_repo{grm}
}

func (r *cart_repo) FindByUserId(id int) (*models.Cart, error) {
	var cart models.Cart

	data := r.db.Preload("Products").Where("users_id = ?", id).Find(&cart)

	if data.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &cart, nil
}

func (r *cart_repo) All() (*models.Cart, error) {
	var cart models.Cart

	data := r.db.Preload("Products").Find(&cart)

	if data.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &cart, nil
}

func (r *cart_repo) Save(usersId uint, items *models.CartItem) (*models.Cart, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(items).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var cart = &models.Cart{UsersId: usersId, CartId: items.CartId}
	if err := tx.Create(cart).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return cart, nil
}
