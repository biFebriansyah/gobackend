package users

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	"github.com/biFebriansyah/gobackend/src/libs"
)

type users_ctrl struct {
	svc interfaces.UsersService
}

func NewCtrl(reps interfaces.UsersService) *users_ctrl {
	return &users_ctrl{svc: reps}
}

func (re *users_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	claims_users := r.Context().Value("user")

	result := re.svc.GetByUUID(claims_users.(string))
	result.Send(w)
}

func (re *users_ctrl) Add(w http.ResponseWriter, r *http.Request) {

	var datas models.User
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}

	re.svc.Add(&datas).Send(w)

}
