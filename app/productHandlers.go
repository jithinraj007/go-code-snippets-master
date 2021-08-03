package app

import (
	"encoding/json"
	"net/http"

	"github.com/ashishjuyal/banking/service"
	"github.com/gorilla/mux"
)

type ProductHandlers struct {
	service service.ProductService
}

func (ph *ProductHandlers) getAllProducts(w http.ResponseWriter, r *http.Request) {

	discontinued := r.URL.Query().Get("discontinued")

	products, err := ph.service.GetAllProduct(discontinued)

	if err != nil {
		writeProductResponse(w, err.Code, err.AsMessage())
	} else {
		writeProductResponse(w, http.StatusOK, products)
	}
}

func (ph *ProductHandlers) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	customer, err := ph.service.GetProduct(id)
	if err != nil {
		writeProductResponse(w, err.Code, err.AsMessage())
	} else {
		writeProductResponse(w, http.StatusOK, customer)
	}
}

func writeProductResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
