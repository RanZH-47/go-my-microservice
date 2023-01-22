package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func HomeHandler(l *log.Logger) *Hello {
	return &Hello{l}
}

// define ServeHTTP - to make Hello a handler
func (h *Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.l.Printf("Hello World!")
	d, err := ioutil.ReadAll(req.Body)
	// defensive programming, do not ignore error handling
	if err != nil {
		http.Error(w, "Oopise", http.StatusBadRequest)
		return
	}
	//print data received
	fmt.Fprintf(w, "Data received: %s\n", d)
}

// handler
// logging
// server configuration
// server mux multiplexer
// language server