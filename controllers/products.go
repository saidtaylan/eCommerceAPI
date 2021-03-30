package controllers

import (
	"ecommerce/views"

	"github.com/gorilla/mux"
)

func ProductsMux(r *mux.Router) {
	r.HandleFunc("/products", views.GetProducts).Methods("GET")
	r.HandleFunc("/products", views.CreateProduct).Methods("POST")
	r.HandleFunc("/sellers", views.CreateSeller).Methods("POST")
	r.HandleFunc("/sellers", views.GetSellers).Methods("GET")
}
