package server

import (
	"ecommerce/controllers"
	"net/http"
)

func Server() {
	server := http.Server{
		Addr:         ":80",
		Handler:      controllers.Multiplexer(),
		ReadTimeout:  10,
		WriteTimeout: 20,
	}
	server.ListenAndServe()
}
