package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaphobia/rajasinga/server/helpers"
	"github.com/astaphobia/rajasinga/server/schemas"
)

func WebLoginController(w http.ResponseWriter, r *http.Request) {
	var (
		response helpers.Response
		authReq  schemas.AuthReq
	)

	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		fmt.Println(err)
		response = helpers.NewResponse(http.StatusInternalServerError, nil, "there's something wrong")
	} else {
		response = helpers.NewResponse(http.StatusOK, authReq, nil)
	}
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
