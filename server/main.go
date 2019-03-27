package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/astaphobia/rajasinga/server/controllers"
)

func main() {
	var (
		PORT string
		wait time.Duration
	)

	flag.DurationVar(&wait, "graceful-timeout", time.Second*100, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	PORT = os.Getenv("PORT")
	if len(PORT) <= 0 {
		PORT = "3001"
	}

	controllers := controllers.ControllerInit()
	srv := &http.Server{
		Handler:      controllers.Router,
		Addr:         fmt.Sprintf("0.0.0.0:%s", PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
