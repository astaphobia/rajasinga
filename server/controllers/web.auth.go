package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaphobia/rajasinga/server/helpers"
	"github.com/astaphobia/rajasinga/server/schemas"
)

func (c *Controllers) WebLoginController(w http.ResponseWriter, r *http.Request) {
	var (
		response helpers.Response
		authReq  schemas.AuthReq
	)

	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&authReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response = helpers.NewResponse(http.StatusInternalServerError, nil, "there's something wrong")
		} else {
			w.WriteHeader(http.StatusOK)
			response = helpers.NewResponse(http.StatusOK, authReq, nil)
		}
		defer r.Body.Close()
	}
	json.NewEncoder(w).Encode(response)
}
