package handler

import (
	"log"
	"net/http"
)

// define the handler struct
type Goodbye struct {
	l *log.Logger
}

func NewGoodbyeHandler(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Bye-bye!"))
}
