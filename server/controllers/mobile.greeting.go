package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaphobia/rajasinga/server/helpers"
)

func MobileGreetingController(w http.ResponseWriter, r *http.Request) {
	var (
		response helpers.Response
	)

	w.WriteHeader(http.StatusOK)
	response = helpers.NewResponse(http.StatusOK, nil, "welcome mobile mux")
	json.NewEncoder(w).Encode(response)
}
