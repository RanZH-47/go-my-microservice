// build a simple server with Go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ranZH-47/go-my-microservice/handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	gl := handler.GeneralLogger(l)

	hd := http.HandleFunc("/", gl)

	http.ListenAndServe(":80", hd)
}
