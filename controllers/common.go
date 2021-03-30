package controllers

import (
	"ecommerce/views"

	"github.com/gorilla/mux"
)

func Multiplexer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", views.Home).Methods("GET")
	ProductsMux(r)
	return r
}
