package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaphobia/rajasinga/server/helpers"
)

func NotFoundController(w http.ResponseWriter, r *http.Request) {
	var (
		response helpers.Response
	)

	w.WriteHeader(http.StatusNotFound)
	response = helpers.NewResponse(http.StatusNotFound, nil, "url not found")
	json.NewEncoder(w).Encode(response)
}
