package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RanZH-47/go-my-microservice/coffee"
)

type Product struct {
	l *log.Logger
}

func NewProductHandler(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("A request is coming in.")
	lp := coffee.GetAllProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal data", http.StatusInternalServerError)
	}
	rw.Write(d)
}
