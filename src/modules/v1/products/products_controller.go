package products

import (
	"encoding/json"
	"net/http"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/biFebriansyah/gobackend/src/interfaces"
	"github.com/biFebriansyah/gobackend/src/libs"
	"github.com/gorilla/schema"
)

// berinterakti dengan service dan router
// untuk mengahandle request yang masuk

type prod_ctrl struct {
	svc interfaces.ProductService
}

func NewCtrl(reps interfaces.ProductService) *prod_ctrl {
	return &prod_ctrl{svc: reps}
}

func (re *prod_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	result := re.svc.GetAll()
	result.Send(w)
}

func (re *prod_ctrl) Add(w http.ResponseWriter, r *http.Request) {

	var datas models.Product
	var decoder = schema.NewDecoder()

	err := r.ParseForm()
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}

	err = decoder.Decode(&datas, r.PostForm)
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}

	uploads := r.Context().Value("file")
	if uploads != nil {
		datas.Image = uploads.(string)
	}

	err = json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Respone(err, 500, true)
		return
	}

	re.svc.Add(&datas).Send(w)

}
