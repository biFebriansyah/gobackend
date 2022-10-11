package interfaces

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/libs"
)

type ProductRepo interface {
	FindAll() (*models.Products, error)
	Save(data *models.Product) (*models.Product, error)
	GetById(uid uint) (*models.Products, error)
}

type ProductService interface {
	GetAll() *libs.Response
	Add(data *models.Product) *libs.Response
	GetProdWithId(uid uint) *libs.Response
}
