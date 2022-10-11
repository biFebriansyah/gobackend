package auth

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	"github.com/biFebriansyah/gobackend/src/libs"
)

type auth_service struct {
	repo interfaces.UsersRepo
}

type token_respone struct {
	Tokens string `json:"token"`
}

func NewService(reps interfaces.UsersRepo) *auth_service {
	return &auth_service{reps}
}

func (a auth_service) Login(body models.User) *libs.Response {
	user, err := a.repo.FindByUsername(body.Username)
	if err != nil {
		return libs.Respone("username tidak terdaftar", 401, true)
	}

	if !libs.CheckPassword(user.Password, body.Password) {
		return libs.Respone("Password salah", 401, true)
	}

	token := libs.NewToken(user.UserId, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return libs.Respone(err.Error(), 401, true)
	}

	return libs.Respone(token_respone{Tokens: theToken}, 200, false)
}
