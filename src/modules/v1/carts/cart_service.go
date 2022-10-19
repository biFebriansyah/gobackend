package carts

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	help "github.com/biFebriansyah/gobackend/src/libs"
)

type cart_service struct {
	re interfaces.CartRepo
}

func NewService(rep interfaces.CartRepo) *cart_service {
	return &cart_service{rep}
}

func (r *cart_service) GetByUserId(id string) *help.Response {

	data, err := r.re.FindByUserId(id)
	if err != nil {
		return help.Respone(err.Error(), 500, true)
	}

	result := help.Respone(data, 200, false)
	return result
}

func (r *cart_service) Get() *help.Response {

	data, err := r.re.All()
	if err != nil {
		return help.Respone(err.Error(), 500, true)
	}

	result := help.Respone(data, 200, false)
	return result
}

func (r cart_service) Add(usersId string, items *models.CartItem) (*help.Response, error) {

	data, err := r.re.Save(usersId, items)
	if err != nil {
		return nil, err
	}

	result := help.Respone(data, 200, false)
	return result, nil
}
