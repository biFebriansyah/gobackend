package carts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
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

	var data models.CartItem
	// claim_user := r.Context().Value("users")
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	result, err := rep.repo.Add(1, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result.Send(w)
}
