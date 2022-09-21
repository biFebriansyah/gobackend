package interfaces

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	help "github.com/biFebriansyah/gobackend/src/libs"
)

type CartRepo interface {
	FindByUserId(id int) (*models.Cart, error)
	All() (*models.Cart, error)
	Save(usersId uint, items *models.CartItem) (*models.Cart, error)
}

type CartService interface {
	Get() *help.Response
	GetByUserId(id string) *help.Response
	Add(usersId uint, items *models.CartItem) (*help.Response, error)
}
