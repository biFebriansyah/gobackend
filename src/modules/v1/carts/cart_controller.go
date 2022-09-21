package carts

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	"github.com/biFebriansyah/gobackend/src/libs"
	"github.com/gorilla/mux"
)

type cart_ctrl struct {
	repo interfaces.CartService
}

func NewCtrl(rep interfaces.CartService) *cart_ctrl {
	return &cart_ctrl{rep}
}

func (rep *cart_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	result := rep.repo.Get()
	result.Send(w)
}

func (rep *cart_ctrl) GetByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)["id"]
	result := rep.repo.GetByUserId(vars)
	result.Send(w)
}

func (rep *cart_ctrl) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var data models.Cart

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&data)
	if err != nil {
		libs.Respone(err.Error(), 400, true)
		return
	}

	result := rep.repo.Save(&data)

	result.Send(w)
}
