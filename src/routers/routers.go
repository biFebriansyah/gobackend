package routers

import (
	"errors"

	"github.com/biFebriansyah/gobackend/src/database/orm"
	"github.com/biFebriansyah/gobackend/src/modules/v1/products"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	products.New(mainRoute, db)

	return mainRoute, nil

}
