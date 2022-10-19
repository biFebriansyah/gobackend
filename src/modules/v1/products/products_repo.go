package products

import (
	"errors"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"gorm.io/gorm"
)

// repo untuk komunikasi ke database

type prod_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *prod_repo {
	return &prod_repo{db}
}

func (r *prod_repo) GetById(uid uint) (*models.Products, error) {

	var data models.Products

	result := r.db.Where("product_id = ?", uid).Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *prod_repo) FindAll() (*models.Products, error) {

	var data models.Products

	result := r.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return &data, nil
}

func (r *prod_repo) Save(data *models.Product) (*models.Product, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	return data, nil
}
