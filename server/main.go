package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/astaphobia/rajasinga/server/controllers"
)

func main() {
	var (
		PORT string
	)

	PORT = os.Getenv("PORT")
	if len(PORT) <= 0 {
		PORT = "3001"
	}

	controllers := controllers.ControllerInit()
	srv := &http.Server{
		Handler: controllers.Router,
		Addr:    fmt.Sprintf("0.0.0.0:%s", PORT),
	}
	log.Fatal(srv.ListenAndServe(), nil)
}
