package interfaces

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	help "github.com/biFebriansyah/gobackend/src/libs"
)

type CartRepo interface {
	FindByUserId(id int) (*models.Cart, error)
	All() (*models.Cart, error)
	Save(data *models.Cart) (*models.Cart, error)
}

type CartService interface {
	Get() *help.Response
	GetByUserId(id string) *help.Response
	Save(data *models.Cart) *help.Response
}
