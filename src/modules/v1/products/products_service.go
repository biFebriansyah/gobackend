package products

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
)

// berinteraksi dengan repo dan controller
// berisi logic bisnis

type prod_service struct {
	repo interfaces.ProductRepo
}

func NewService(reps interfaces.ProductRepo) *prod_service {
	return &prod_service{reps}
}

func (r *prod_service) GetAll() (*models.Products, error) {
	data, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *prod_service) Add(data *models.Product) (*models.Product, error) {
	data, err := r.repo.Save(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
