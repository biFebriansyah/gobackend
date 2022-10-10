package products

import (
	"github.com/biFebriansyah/gobackend/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// akan memangil semua method
// inisialisasi endpoint

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/product").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", middleware.Handle(ctrl.Add, middleware.AuthWithRole("users"), middleware.FileUpload)).Methods("POST")
}
