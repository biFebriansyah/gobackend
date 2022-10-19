package interfaces

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/libs"
)

type AuthService interface {
	Login(body models.User) *libs.Response
}
