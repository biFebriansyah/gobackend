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

	route.HandleFunc("", middleware.Handle(ctrl.GetAll, middleware.AuthWithRole("users"))).Methods("GET")
	route.HandleFunc("", ctrl.Add).Methods("POST")
}
