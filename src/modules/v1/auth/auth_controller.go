package auth

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	"github.com/biFebriansyah/gobackend/src/libs"
)

type auth_ctrl struct {
	rep interfaces.AuthService
}

func NewCtrl(rep interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{rep}
}

func (a auth_ctrl) Signin(w http.ResponseWriter, r *http.Request) {
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 401, true)
		return
	}

	a.rep.Login(data).Send(w)
}
