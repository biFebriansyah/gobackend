package carts

import (
	"errors"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cart_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *cart_repo {
	return &cart_repo{grm}
}

func (r *cart_repo) FindByUserId(id string) (*models.Cart, error) {
	var cart models.Cart

	err := r.db.Preload("Items.Products").Preload(clause.Associations).Where("users_id = ?", id).Find(&cart)

	if err.Error != nil {
		return nil, err.Error
	}

	return &cart, nil
}

func (r *cart_repo) All() (*models.Cart, error) {
	var cart models.Cart

	data := r.db.Preload("Items.Products").Preload(clause.Associations).Find(&cart)

	if data.Error != nil {
		return nil, errors.New("Gagal mengambil data")
	}

	return &cart, nil
}

func (r *cart_repo) Save(usersId string, items *models.CartItem) (*models.Cart, error) {
	tx := r.db.Begin()
	var cart models.Cart

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Where("users_id = ?", usersId).Find(&cart); cart.CartId == "" {
		cart.UsersId = usersId
		if err := tx.Create(&cart).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	items.CartId = cart.CartId
	if err := tx.Create(items).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &cart, nil
}
