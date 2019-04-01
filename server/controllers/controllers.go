package controllers

import (
	"database/sql"
	"net/http"

	"github.com/astaphobia/rajasinga/server/libs"
	"github.com/gorilla/mux"
)

type Controllers struct {
	Router *mux.Router
	DB     *sql.DB
}

func (c *Controllers) Init() {
	var (
		logger    = libs.Logger{}
		requestID = libs.RequestID{}
	)
	logger.LoggerInit()

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFoundController)
	router.Use(requestID.Middleware, logger.Middleware)
	router.StrictSlash(true)
	router.Headers(
		"Content-Type", "application/json", "X-RequestWith", "Authorization",
	)
	c.Router = router
}

func (c *Controllers) RegisterRoutes() {
	adminAPI := c.Router.PathPrefix("/w").Subrouter()
	adminAPI.HandleFunc("/greeting", c.GreetingController).
		Methods("GET")
	adminAPI.HandleFunc("/login", c.WebLoginController).
		Methods("POST")
}
