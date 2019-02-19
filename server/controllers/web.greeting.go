package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaphobia/rajasinga/server/helpers"
)

func GreetingController(w http.ResponseWriter, r *http.Request) {
	var (
		response helpers.Response
	)

	w.WriteHeader(http.StatusOK)
	response = helpers.NewResponse(http.StatusOK, nil, "welcome mux/gorilla")
	json.NewEncoder(w).Encode(response)
}
