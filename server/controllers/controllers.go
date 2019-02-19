package controllers

import (
	"github.com/gorilla/mux"
)

type Controllers struct {
	Router *mux.Router
}

func ControllerInit() *Controllers {
	var (
		controller *Controllers
	)

	route := mux.NewRouter()
	route.Headers(
		"Content-Type", "application/json",
	)

	adminAPI := route.PathPrefix("/admin").Subrouter()
	adminAPI.HandleFunc("/", GreetingController).Methods("GET")

	controller = &Controllers{
		Router: route,
	}
	return controller
}
