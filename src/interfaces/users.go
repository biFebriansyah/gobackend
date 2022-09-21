package interfaces

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/libs"
)

type UsersRepo interface {
	FindAll() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	UserExsist(username string) bool
	Add(data *models.User) (*models.User, error)
}

type UsersService interface {
	FindAll() *libs.Response
	FindByUsername(username string) *libs.Response
	Add(data *models.User) *libs.Response
}
