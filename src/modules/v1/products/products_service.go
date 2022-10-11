package products

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	"github.com/biFebriansyah/gobackend/src/libs"
)

// berinteraksi dengan repo dan controller
// berisi logic bisnis

type prod_service struct {
	repo interfaces.ProductRepo
}

func NewService(reps interfaces.ProductRepo) *prod_service {
	return &prod_service{reps}
}

func (r *prod_service) GetProdWithId(uid uint) *libs.Response {
	data, err := r.repo.GetById(uid)
	if err != nil {
		return libs.Respone(err.Error(), 500, true)
	}

	return libs.Respone(data, 200, false)
}

func (r *prod_service) GetAll() *libs.Response {
	data, err := r.repo.FindAll()
	if err != nil {
		return libs.Respone(err.Error(), 500, true)
	}

	return libs.Respone(data, 200, false)
}

func (r *prod_service) Add(data *models.Product) *libs.Response {
	fileURL, err := libs.CloudUpload(data.Image)
	if err != nil {
		return libs.Respone(err.Error(), 500, true)
	}

	data.Image = fileURL
	data, err = r.repo.Save(data)
	if err != nil {
		return libs.Respone(err.Error(), 500, true)
	}

	return libs.Respone(data.ProductId, 201, false)
}
