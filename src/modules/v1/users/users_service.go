package users

import (
	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	"github.com/biFebriansyah/gobackend/src/libs"
)

type user_service struct {
	repo interfaces.UsersRepo
}

func NewService(reps interfaces.UsersRepo) *user_service {
	return &user_service{reps}
}

func (r user_service) FindAll() *libs.Response {

	data, err := r.repo.FindAll()
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(data, 200, false)

}

func (r user_service) FindByUsername(username string) *libs.Response {

	data, err := r.repo.FindByUsername(username)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(data, 200, false)
}

func (r user_service) Add(data *models.User) *libs.Response {
	if check := r.repo.UserExsist(data.Username); check {
		return libs.Respone("username sudah terdaftar", 400, true)
	}

	hassPassword, err := libs.HashPasword(data.Password)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	data.Password = hassPassword
	result, err := r.repo.Add(data)
	if err != nil {
		return libs.Respone(err.Error(), 400, true)
	}

	return libs.Respone(result, 200, false)
}
