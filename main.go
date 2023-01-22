// build a simple server with Go
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	gl := GeneralLogger(l)

	http.ListenAndServe(":80", nil)
}
