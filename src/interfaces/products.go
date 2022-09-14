package interfaces

import "github.com/biFebriansyah/gobackend/src/database/orm/models"

type ProductRepo interface {
	FindAll() (*models.Products, error)
	Save(data *models.Product) (*models.Product, error)
}

type ProductService interface {
	GetAll() (*models.Products, error)
	Add(data *models.Product) (*models.Product, error)
}
