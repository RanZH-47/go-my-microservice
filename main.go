// build a simple server with Go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RanZH-47/go-my-microservice/handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handler.HomeHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":80", sm)
}
