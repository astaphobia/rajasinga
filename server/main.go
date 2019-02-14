package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var (
		PORT string
	)
	PORT = os.Getenv("PORT")
	if len(PORT) <= 0 {
		PORT = "3001"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", PORT),
	}
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "yiihaaa")
}
