// build a simple server with Go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RanZH-47/go-my-microservice/handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handler.NewHomeHandler(l)
	gb := handler.NewGoodbyeHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	s := &http.Server{
		Addr:         ":80",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	sig := <-signChan
	l.Println("Received signal", sig)

	d := time.Now().Add(10 * time.Second)
	tc, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	s.Shutdown(tc)
}
