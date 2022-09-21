package users

import (
	"github.com/biFebriansyah/gobackend/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("", ctrl.Add).Methods("POST")
}
