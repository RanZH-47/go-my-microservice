// build a simple server with Go
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	gl := handler.GeneralLogger(l)

	hd := http.HandleFunc("/", gl)

	http.ListenAndServe(":80", hd)
}
